package sandbox

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/e2b-dev/infra/packages/shared/pkg/consts"
)

func (s *Sandbox) logHeathAndUsage(exited chan struct{}) {
	ctx := context.Background()
	for {
		select {
		case <-time.After(10 * time.Second):
			childCtx, cancel := context.WithTimeout(ctx, time.Second)
			s.Healthcheck(childCtx, false)
			cancel()

			stats, err := s.stats.getStats()
			if err != nil {
				s.Logger.Warnf("failed to get stats: %s", err)
			} else {
				s.Logger.CPUUsage(stats.CPUCount)
				s.Logger.MemoryUsage(stats.MemoryMB)
			}
		case <-exited:
			return
		}
	}
}

func (s *Sandbox) Healthcheck(ctx context.Context, alwaysReport bool) {
	var err error
	// Report healthcheck status
	defer s.Logger.Healthcheck(err == nil, alwaysReport)

	address := fmt.Sprintf("http://%s:%d/health", s.slot.HostSnapshotIP(), consts.DefaultEnvdServerPort)

	request, err := http.NewRequestWithContext(ctx, "GET", address, nil)
	if err != nil {
		return
	}

	response, err := httpClient.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusNoContent {
		err = fmt.Errorf("unexpected status code: %d", response.StatusCode)
		return
	}

	_, err = io.Copy(io.Discard, response.Body)
	if err != nil {
		return
	}
}
