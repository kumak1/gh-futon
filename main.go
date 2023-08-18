package main

import (
	"fmt"
	"github.com/kumak1/gh-futon/internal/cli"
	gh_ql "github.com/kumak1/gh-futon/internal/gh/graphql"
	"io"
	"os"
	"sync"
)

var (
	variable  gh_ql.QueryVariable
	optionAny bool
	writer    io.Writer
)

type (
	info struct {
		name  string
		nodes []gh_ql.NodeInfo
	}
)

func init() {
	variable = gh_ql.QueryVariable{
		Username: cli.User,
		From:     &cli.From,
		To:       &cli.To,
	}
	optionAny = cli.Item.Any()
	writer = os.Stdout
}

func fprintln(a ...interface{}) {
	fmt.Fprintln(writer, a...)
}

func main() {
	for _, a := range fetch() {

		fprintln("## " + a.name)
		for _, node := range a.nodes {
			fprintln(node.ToMarkdownList())
		}
	}
}

func fetch() []info {
	var results []info

	wg := sync.WaitGroup{}

	if cli.Item.Issue || optionAny {
		wg.Add(1)
		go func() {
			results = append(
				results,
				info{name: "Issue", nodes: gh_ql.GetIssue(variable)},
			)
			wg.Done()
		}()
	}

	if cli.Item.IssueComment || optionAny {
		wg.Add(1)
		go func() {
			results = append(results, info{name: "Issue Comment", nodes: gh_ql.ExcludeAuthor(gh_ql.GetIssueComment(variable), variable.Username)})
			wg.Done()
		}()
	}

	if cli.Item.PullRequest || optionAny {
		wg.Add(1)
		go func() {
			results = append(results, info{name: "PR", nodes: gh_ql.GetPullRequest(variable)})
			wg.Done()
		}()
	}

	if cli.Item.PullRequestReview || optionAny {
		wg.Add(1)
		go func() {
			results = append(results, info{name: "Review", nodes: gh_ql.ExcludeAuthor(gh_ql.GetPullRequestReview(variable), variable.Username)})
			wg.Done()
		}()
	}

	wg.Wait()

	return results
}
