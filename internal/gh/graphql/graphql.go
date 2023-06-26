package graphql

import (
	"github.com/cli/go-gh/v2/pkg/api"
	graphql "github.com/cli/shurcooL-graphql"
	"time"
)

var graphClient *api.GraphQLClient

func init() {
	var err error
	graphClient, err = api.DefaultGraphQLClient()
	if err != nil {
		panic(err)
	}
}

type (
	DateTime struct{ time.Time }

	NodeInfo struct {
		Title  graphql.String
		Url    graphql.String
		Author struct {
			Login graphql.String
		}
		Labels struct {
			Nodes []struct {
				Name graphql.String
			}
		} `graphql:"labels(first: 20)"`
	}

	pageInfo struct {
		HasNextPage graphql.Boolean
		EndCursor   graphql.String
	}

	issueContributions struct {
		PageInfo pageInfo
		Nodes    []struct {
			Issue NodeInfo
		}
	}

	pullRequestContributions struct {
		PageInfo pageInfo
		Nodes    []struct {
			PullRequest NodeInfo
		}
	}

	paginateQuery interface {
		nextQuery() interface{}
		hasNextPage() graphql.Boolean
		endCursor() graphql.String
		nodes() []NodeInfo
	}
)

func GetIssue(username string, from time.Time, to time.Time) []NodeInfo {
	return getContributions(IssueQuery{}, getVariables(username, from, to))
}

func GetPullRequest(username string, from time.Time, to time.Time) []NodeInfo {
	return getContributions(PullRequestQuery{}, getVariables(username, from, to))
}

func GetPullRequestReview(username string, from time.Time, to time.Time) []NodeInfo {
	return getContributions(PullRequestReviewQuery{}, getVariables(username, from, to))
}

func getVariables(username string, from time.Time, to time.Time) map[string]interface{} {
	return map[string]interface{}{
		"username": graphql.String(username),
		"first":    graphql.Int(10),
		"from":     DateTime{from},
		"to":       DateTime{to},
	}
}

func getContributions(q interface{}, variables map[string]interface{}) []NodeInfo {
	var contributes []NodeInfo
	var paginate paginateQuery

	switch q.(type) {
	case IssueQuery:
		paginate = IssueQuery{}.execQuery(variables)
	case IssueNextQuery:
		paginate = IssueNextQuery{}.execQuery(variables)
	case PullRequestQuery:
		paginate = PullRequestQuery{}.execQuery(variables)
	case PullRequestNextQuery:
		paginate = PullRequestNextQuery{}.execQuery(variables)
	case PullRequestReviewQuery:
		paginate = PullRequestReviewQuery{}.execQuery(variables)
	case PullRequestReviewNextQuery:
		paginate = PullRequestReviewNextQuery{}.execQuery(variables)
	}

	if paginate.hasNextPage() {
		variables["after"] = paginate.endCursor()
		contributes = append(getContributions(paginate.nextQuery(), variables), contributes...)
	}

	contributes = append(paginate.nodes(), contributes...)
	return contributes
}