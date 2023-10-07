package nomad

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/nomad/api"
)

const (
	shortNodeIDLength = 8

	taskRunningState = "running"
	taskDeadState    = "dead"
	taskPendingState = "pending"

	defaultTaskName = "start"
)

type AllocResult struct {
	Alloc *api.AllocationListStub
	Err   error
}

type JobSubscriber struct {
	jobID     string
	events    chan AllocResult
	taskState string
}

func (n *NomadClient) newSubscriber(jobID string, taskState string) *JobSubscriber {
	sub := &JobSubscriber{
		jobID:     jobID,
		events:    make(chan AllocResult),
		taskState: taskState,
	}
	n.subscribers.Insert(jobID, sub)

	return sub
}

func (n *NomadClient) ListenToJobs(ctx context.Context) {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			jobs, _, err := n.client.Jobs().List(nil)
			if err != nil {
				log.Printf("Error getting jobs: %s\n", err)
				return
			}

			for _, job := range jobs {
				err := n.processJobEvent(job)
				if err != nil {
					log.Printf("Error processing job event: %s\n", err)
				}
			}
		case <-ctx.Done():
			log.Println("Context canceled, stopping ListenToJobs")
			return
		}
	}
}

func (n *NomadClient) processJobEvent(job *api.JobListStub) error {
	sub, ok := n.subscribers.Get(job.ID)
	if !ok {
		return nil
	}

	switch job.Status {
	case taskDeadState:
		if sub.taskState == taskDeadState {
			alloc, allocErr, err := n.getAllocTaskErr(job, defaultTaskName)
			if err != nil {
				errMsg := fmt.Errorf("error getting allocation for job %s: %s", job.ID, err)
				sub.events <- AllocResult{
					Err: errMsg,
				}
			}

			if allocErr != nil {
				sub.events <- AllocResult{
					Err: allocErr,
				}
			}

			if alloc != nil {
				sub.events <- AllocResult{
					Alloc: alloc,
				}
			}
		}

	case taskRunningState:
		if sub.taskState == taskRunningState {
			sub.events <- AllocResult{}
		}

	case taskPendingState:

	default:
		return fmt.Errorf("unknown job status: %s", job.Status)
	}

	return nil
}

func (n *NomadClient) getAllocTaskErr(job *api.JobListStub, taskName string) (*api.AllocationListStub, error, error) {
	allocations, _, err := n.client.Jobs().Allocations(job.ID, true, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("error getting allocation for job %s: %w", job.ID, err)
	}

	for _, alloc := range allocations {
		if alloc == nil {
			continue
		}

		if alloc.TaskStates[taskName] == nil {
			continue
		}

		if alloc.TaskStates[taskName].State == taskDeadState {
			if alloc.TaskStates[taskName].Failed {
				return nil, fmt.Errorf("allocation is %s for '%s' job", alloc.TaskStates[taskName].State, job.ID), nil
			} else {
				return alloc, nil, nil
			}
		}
	}

	return nil, nil, fmt.Errorf("no allocation with the task name %s found", taskName)
}

func (n *NomadClient) WaitForJob(ctx context.Context, jobID, taskState string, result chan AllocResult) {
	sub := n.newSubscriber(jobID, taskState)
	defer n.subscribers.Remove(jobID)
	defer close(sub.events)

	select {
	case err := <-sub.events:
		result <- err

	case <-ctx.Done():
		result <- AllocResult{
			Err: fmt.Errorf("waiting for job '%s' cancelled", jobID),
		}
	}
}
