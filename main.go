package main

import (
	"os"

	"github.com/google/go-github/v42/github"
	"github.com/sethvargo/go-githubactions"
)

func main() {
	eventName := os.Getenv("GITHUB_EVENT_NAME")
	eventPath := os.Getenv("GITHUB_EVENT_PATH")

	eventPayload, err := os.ReadFile(eventPath)
	if err != nil {
		githubactions.Fatalf("failed to load event payload from %s: %s", eventPath, err)
	}

	os.Exit(Exec(eventName, eventPayload))
}

func Exec(eventName string, eventPayload []byte) int {
	event, err := github.ParseWebHook(eventName, eventPayload)
	if err != nil {
		githubactions.Fatalf("failed to parse event: %v", err)
	}

	switch e := event.(type) {
	case *github.PullRequestEvent:
		if !isTaskListCompleted([]byte(e.PullRequest.GetBody())) {
			return 1
		}
	default:
		githubactions.Infof("unknown event: %s", eventName)
	}

	return 0
}

func isTaskListCompleted(body []byte) bool {
	githubactions.Debugf("body: %s", body)

	tasks := ExtractTaskList(body)

	countDone := 0

	for _, task := range tasks {
		githubactions.Infof("- %s", task.RawText)

		if task.IsChecked {
			countDone++
		}
	}

	if countDone < len(tasks) {
		githubactions.Warningf("some task are not completed: %d/%d", countDone, len(tasks))

		return false
	}

	githubactions.Infof("all done: %d/%d", countDone, len(tasks))

	return true
}
