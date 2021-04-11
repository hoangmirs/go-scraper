package custom_matchers

import (
	"fmt"

	"github.com/gocraft/work"
	"github.com/onsi/gomega/types"
)

// EnqueueJob checks if the `jobName` is enqueued `count` time
func EnqueueJob(jobName string, count int64) types.GomegaMatcher {
	return &EnqueueJobMatcher{
		jobName: jobName,
		count:   count,
	}
}

type EnqueueJobMatcher struct {
	jobName string
	count   int64
}

func (matcher *EnqueueJobMatcher) Match(actual interface{}) (bool, error) {
	workerClient := actual.(*work.Client)

	queues, err := workerClient.Queues()
	if err != nil {
		return false, err
	}

	for _, queue := range queues {
		if queue.JobName == matcher.jobName {
			return queue.Count == matcher.count, nil
		}
	}

	return false, nil
}

func (matcher *EnqueueJobMatcher) FailureMessage(actual interface{}) string {
	return fmt.Sprintf("Expected app to enqueue %v job %d time(s)", matcher.jobName, matcher.count)
}

func (matcher *EnqueueJobMatcher) NegatedFailureMessage(actual interface{}) string {
	return fmt.Sprintf("Expected app NOT to enqueue `%s` job %d time(s)", matcher.jobName, matcher.count)
}
