package graphql

import (
	"fmt"
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
		Title     graphql.String
		Url       graphql.String
		CreatedAt DateTime
		ClosedAt  DateTime
		Author    struct {
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

	issueComments struct {
		PageInfo pageInfo
		Nodes    []struct {
			CreatedAt DateTime
			Issue     NodeInfo
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

	QueryVariable struct {
		Username string
		From     *time.Time
		To       *time.Time
		After    graphql.String
	}
)

func (r NodeInfo) ToMarkdownList() string {
	return fmt.Sprintf("- [%s](%s)", r.Title, r.Url)
}

func (q QueryVariable) ToMap() map[string]interface{} {
	result := map[string]interface{}{
		"username": graphql.String(q.Username),
		"first":    graphql.Int(100),
	}
	if q.From != nil {
		result["from"] = DateTime{*q.From}
	}
	if q.From != nil {
		result["to"] = DateTime{*q.To}
	}
	if q.After != "" {
		result["after"] = q.After
	}
	return result
}

func (q QueryVariable) ToMapForComment() map[string]interface{} {
	result := q.ToMap()
	delete(result, "from")
	delete(result, "to")
	return result
}

func GetIssue(variable QueryVariable) []NodeInfo {
	return getContributions(issueQuery{}, variable)
}

func GetIssueComment(variable QueryVariable) []NodeInfo {
	return getContributions(issueCommentQuery{}, variable)
}

func GetPullRequest(variable QueryVariable) []NodeInfo {
	return getContributions(pullRequestQuery{}, variable)
}

func GetPullRequestReview(variable QueryVariable) []NodeInfo {
	return getContributions(pullRequestReviewQuery{}, variable)
}

func execQuery(q interface{}, variable QueryVariable) (paginateQuery, error) {
	switch q.(type) {
	case issueQuery:
		return q.(issueQuery).execQuery(variable)
	case issueNextQuery:
		return q.(issueNextQuery).execQuery(variable)
	case issueCommentQuery:
		return q.(issueCommentQuery).execQuery(variable)
	case issueCommentNextQuery:
		return q.(issueCommentNextQuery).execQuery(variable)
	case pullRequestQuery:
		return q.(pullRequestQuery).execQuery(variable)
	case pullRequestNextQuery:
		return q.(pullRequestNextQuery).execQuery(variable)
	case pullRequestReviewQuery:
		return q.(pullRequestReviewQuery).execQuery(variable)
	case pullRequestReviewNextQuery:
		return q.(pullRequestReviewNextQuery).execQuery(variable)
	}

	return nil, fmt.Errorf("invalid query")
}

func getContributions(q interface{}, variable QueryVariable) []NodeInfo {
	var contributes []NodeInfo

	paginate, err := execQuery(q, variable)
	if err != nil {
		panic(err)
	}

	if paginate.hasNextPage() {
		variable.After = paginate.endCursor()
		contributes = append(getContributions(paginate.nextQuery(), variable), contributes...)
	}

	contributes = append(paginate.nodes(), contributes...)
	return contributes
}
