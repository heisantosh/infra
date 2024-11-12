package local_storage

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/e2b-dev/infra/packages/orchestrator/internal/sandbox/local_storage/nbd"
	template "github.com/e2b-dev/infra/packages/shared/pkg/storage"

	"github.com/pojntfx/go-nbd/pkg/backend"
)

const ChunkSize = 2 * 1024 * 1024 // 2MiB

type RootfsOverlay struct {
	// TODO: Remove - Only for debugging
	cachePath string

	storage    *template.BlockStorage
	mnt        *nbd.ManagedPathMount
	localCache *os.File

	ctx       context.Context
	cancelCtx context.CancelFunc

	ready chan string
}

func (t *Template) NewRootfsOverlay(cachePath string) (*RootfsOverlay, error) {
	ctx, cancel := context.WithCancel(context.Background())

	rootfs, err := t.Rootfs()
	if err != nil {
		cancel()
		return nil, fmt.Errorf("error getting rootfs: %w", err)
	}

	f, err := os.Create(cachePath)
	if err != nil {
		cancel()
		return nil, fmt.Errorf("error creating overlay file: %w", err)
	}

	size, err := rootfs.Size()
	if err != nil {
		cancel()
		return nil, fmt.Errorf("error getting rootfs size: %w", err)
	}

	err = f.Truncate(size)
	if err != nil {
		cancel()
		return nil, fmt.Errorf("error truncating overlay file: %w", err)
	}

	mnt := nbd.NewManagedPathMount(
		ctx,
		rootfs,
		backend.NewFileBackend(f),
		ChunkSize,
	)

	ready := make(chan string, 1)

	return &RootfsOverlay{
		cachePath:  cachePath,
		ready:      ready,
		mnt:        mnt,
		localCache: f,
		storage:    rootfs,
		ctx:        ctx,
		cancelCtx:  cancel,
	}, nil
}

func (o *RootfsOverlay) Run() error {
	defer close(o.ready)
	defer o.cancelCtx()

	var wg sync.WaitGroup

	wg.Add(1)

	file, _, err := o.mnt.Open(o.ctx)
	if err != nil {
		return fmt.Errorf("error opening overlay file: %w", err)
	}

	go func() {
		defer wg.Done()

		<-o.ctx.Done()

		err := o.mnt.Close()
		if err != nil {
			log.Printf("error closing overlay mount: %v\n", err)
		}

		err = o.localCache.Close()
		if err != nil {
			log.Printf("[%s] error closing overlay file: %v\n", o.cachePath, err)
		}

		err = os.Remove(o.localCache.Name())
		if err != nil {
			log.Printf("[%s] error removing overlay file: %v\n", o.cachePath, err)
		}

		counter := 0
		for {
			counter++
			err := nbd.Pool.ReleaseDevice(file)
			if err != nil {
				if counter%100 == 0 {
					log.Printf("[%s - %dth try] error releasing overlay device: %v\n", o.cachePath, counter, err)
				}

				continue
			}

			break
		}
	}()

	o.ready <- file

	wg.Wait()

	return o.mnt.Wait()
}

func (o *RootfsOverlay) Close() {
	o.cancelCtx()
}

// Path can only be called once.
func (o *RootfsOverlay) Path(ctx context.Context) (string, error) {
	select {
	case <-o.ctx.Done():
		return "", fmt.Errorf("overlay context canceled when getting overlay path: %w", o.ctx.Err())
	case <-ctx.Done():
		return "", fmt.Errorf("context canceled when getting overlay path: %w", ctx.Err())
	case path, ok := <-o.ready:
		if !ok {
			return "", fmt.Errorf("overlay path channel closed")
		}

		return path, nil
	}
}
