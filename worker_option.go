package gue

import (
	"time"

	"github.com/vgarvardt/gue/v3/adapter"
)

// WorkerOption defines a type that allows to set worker properties during the build-time.
type WorkerOption func(*Worker)

// WorkerPoolOption defines a type that allows to set worker pool properties during the build-time.
type WorkerPoolOption func(pool *WorkerPool)

// WithWorkerPollInterval overrides default poll interval with the given value.
// Poll interval is the "sleep" duration if there were no jobs found in the DB.
func WithWorkerPollInterval(d time.Duration) WorkerOption {
	return func(w *Worker) {
		w.interval = d
	}
}

// WithWorkerQueue overrides default worker queue name with the given value.
func WithWorkerQueue(queue string) WorkerOption {
	return func(w *Worker) {
		w.queue = queue
	}
}

// WithWorkerID sets worker ID for easier identification in logs
func WithWorkerID(id string) WorkerOption {
	return func(w *Worker) {
		w.id = id
	}
}

// WithWorkerLogger sets Logger implementation to worker
func WithWorkerLogger(logger adapter.Logger) WorkerOption {
	return func(w *Worker) {
		w.logger = logger
	}
}

// WithWorkerPreserveCompletedJobs determines whether jobs are deleted from the gue_jobs table after being performed.
func WithWorkerPreserveCompletedJobs(preserve bool) WorkerOption {
	return func(w *Worker) {
		w.preserveCompletedJobs = preserve
	}
}

func WithWorkerMigrateCompletedJobs(migrate bool) WorkerOption {
	return func(w *Worker) {
		w.migrateCompletedJob = migrate
	}
}

// WithWorkerHooksJobLocked sets hooks that are called right after the job was polled from the DB.
// Depending on the polling results hook will have either error or job set, but not both.
// If the error field is set - no other lifecycle hooks will be called for the job.
func WithWorkerHooksJobLocked(hooks ...HookFunc) WorkerOption {
	return func(w *Worker) {
		w.hooksJobLocked = hooks
	}
}

// WithWorkerHooksUnknownJobType sets hooks that are called when worker finds a job with unknown type.
// Error field for this event type is always set since this is an error situation.
// If this hook is called - no other lifecycle hooks will be called for the job.
func WithWorkerHooksUnknownJobType(hooks ...HookFunc) WorkerOption {
	return func(w *Worker) {
		w.hooksUnknownJobType = hooks
	}
}

// WithWorkerHooksJobDone sets hooks that are called when worker finished working the job.
// Error field is set for the cases when the job was worked with an error.
func WithWorkerHooksJobDone(hooks ...HookFunc) WorkerOption {
	return func(w *Worker) {
		w.hooksJobDone = hooks
	}
}

// WithWorkerPollStrategy overrides default poll strategy with given value
func WithWorkerPollStrategy(s PollStrategy) WorkerOption {
	return func(w *Worker) {
		w.pollStrategy = s
	}
}

// WithPoolPollInterval overrides default poll interval with the given value.
// Poll interval is the "sleep" duration if there were no jobs found in the DB.
func WithPoolPollInterval(d time.Duration) WorkerPoolOption {
	return func(w *WorkerPool) {
		w.interval = d
	}
}

// WithPoolQueue overrides default worker queue name with the given value.
func WithPoolQueue(queue string) WorkerPoolOption {
	return func(w *WorkerPool) {
		w.queue = queue
	}
}

// WithPoolID sets worker pool ID for easier identification in logs
func WithPoolID(id string) WorkerPoolOption {
	return func(w *WorkerPool) {
		w.id = id
	}
}

// WithPoolLogger sets Logger implementation to worker pool
func WithPoolLogger(logger adapter.Logger) WorkerPoolOption {
	return func(w *WorkerPool) {
		w.logger = logger
	}
}

// WithPoolPreserveCompletedJobs determines whether jobs are deleted from the gue_jobs table after being performed.
func WithPoolPreserveCompletedJobs(preserve bool) WorkerPoolOption {
	return func(pool *WorkerPool) {
		pool.preserveCompletedJobs = preserve
	}
}

func WithPoolMigrateCompletedJobs(migrate bool) WorkerPoolOption {
	return func(pool *WorkerPool) {
		pool.migrateCompletedJob = migrate
	}
}

// WithPoolPollStrategy overrides default poll strategy with given value
func WithPoolPollStrategy(s PollStrategy) WorkerPoolOption {
	return func(w *WorkerPool) {
		w.pollStrategy = s
	}
}

// WithPoolHooksJobLocked calls WithWorkerHooksJobLocked for every worker in the pool.
func WithPoolHooksJobLocked(hooks ...HookFunc) WorkerPoolOption {
	return func(w *WorkerPool) {
		w.hooksJobLocked = hooks
	}
}

// WithPoolHooksUnknownJobType calls WithWorkerHooksUnknownJobType for every worker in the pool.
func WithPoolHooksUnknownJobType(hooks ...HookFunc) WorkerPoolOption {
	return func(w *WorkerPool) {
		w.hooksUnknownJobType = hooks
	}
}

// WithPoolHooksJobDone calls WithWorkerHooksJobDone for every worker in the pool.
func WithPoolHooksJobDone(hooks ...HookFunc) WorkerPoolOption {
	return func(w *WorkerPool) {
		w.hooksJobDone = hooks
	}
}
