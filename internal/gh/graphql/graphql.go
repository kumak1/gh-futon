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
		Repository struct {
			Name          graphql.String
			NameWithOwner graphql.String
		}
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
	return getContributions(issueQuery{}, getVariables(username, from, to))
}

func GetPullRequest(username string, from time.Time, to time.Time) []NodeInfo {
	return getContributions(pullRequestQuery{}, getVariables(username, from, to))
}

func GetPullRequestReview(username string, from time.Time, to time.Time) []NodeInfo {
	return getContributions(pullRequestReviewQuery{}, getVariables(username, from, to))
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
	case issueQuery:
		paginate = issueQuery{}.execQuery(variables)
	case issueNextQuery:
		paginate = issueNextQuery{}.execQuery(variables)
	case pullRequestQuery:
		paginate = pullRequestQuery{}.execQuery(variables)
	case pullRequestNextQuery:
		paginate = pullRequestNextQuery{}.execQuery(variables)
	case pullRequestReviewQuery:
		paginate = pullRequestReviewQuery{}.execQuery(variables)
	case pullRequestReviewNextQuery:
		paginate = pullRequestReviewNextQuery{}.execQuery(variables)
	}

	if paginate.hasNextPage() {
		variables["after"] = paginate.endCursor()
		contributes = append(getContributions(paginate.nextQuery(), variables), contributes...)
	}

	contributes = append(paginate.nodes(), contributes...)
	return contributes
}
