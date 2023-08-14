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

func writeLines(infos []gh_ql.NodeInfo) {
	for _, c := range infos {
		fprintln(c.ToMarkdownList())
	}
}

func main() {
	issueChan := make(chan []gh_ql.NodeInfo, 1)
	issueCommentChan := make(chan []gh_ql.NodeInfo, 1)
	prChan := make(chan []gh_ql.NodeInfo, 1)
	prReviewChan := make(chan []gh_ql.NodeInfo, 1)
	wg := sync.WaitGroup{}

	// defer で出力処理しているので、逆順で処理を記述する

	if cli.Item.PullRequestReview || optionAny {
		wg.Add(1)
		go func() {
			prReviewChan <- gh_ql.ExcludeAuthor(gh_ql.GetPullRequestReview(variable), variable.Username)
			wg.Done()
		}()
		defer func() {
			fprintln("## Review")
			writeLines(<-prReviewChan)
		}()
		defer close(prReviewChan)
	}

	if cli.Item.PullRequest || optionAny {
		wg.Add(1)
		go func() {
			prChan <- gh_ql.GetPullRequest(variable)
			wg.Done()
		}()
		defer func() {
			fprintln("## PR")
			writeLines(<-prChan)
		}()
		defer close(prChan)
	}

	if cli.Item.IssueComment || optionAny {
		wg.Add(1)
		go func() {
			issueCommentChan <- gh_ql.ExcludeAuthor(gh_ql.GetIssueComment(variable), variable.Username)
			wg.Done()
		}()
		defer func() {
			fprintln("## Issue Comment")
			writeLines(<-issueCommentChan)
		}()
		defer close(issueCommentChan)
	}

	if cli.Item.Issue || optionAny {
		wg.Add(1)
		go func() {
			issueChan <- gh_ql.GetIssue(variable)
			wg.Done()
		}()
		defer func() {
			fprintln("## Issue")
			writeLines(<-issueChan)
		}()
		defer close(issueChan)
	}

	wg.Wait()
}
