package graphql

import graphql "github.com/cli/shurcooL-graphql"

type (
	issueCommentQuery struct {
		User struct {
			IssueComments issueComments `graphql:"issueComments(first: $first, orderBy: {field: UPDATED_AT, direction: DESC},)"`
		} `graphql:"user(login: $username)"`
	}
)

var issueCommentVariable QueryVariable

func (p issueCommentQuery) execQuery(variable QueryVariable) (paginateQuery, error) {
	issueCommentVariable = variable
	err := graphClient.Query("issueComment", &p, variable.ToMapForComment())
	return p, err
}

func (p issueCommentQuery) nextQuery() interface{} {
	return issueCommentNextQuery{}
}

func (p issueCommentQuery) hasNextPage() graphql.Boolean {
	if !p.User.IssueComments.PageInfo.HasNextPage {
		return false
	}

	createdAt := p.User.IssueComments.Nodes[0].CreatedAt
	return graphql.Boolean(createdAt.After(*issueCommentVariable.From) && createdAt.Before(*issueCommentVariable.To))
}

func (p issueCommentQuery) endCursor() graphql.String {
	return p.User.IssueComments.PageInfo.EndCursor
}

func (p issueCommentQuery) nodes() []NodeInfo {
	var nodes []NodeInfo
	for _, n := range p.User.IssueComments.Nodes {
		nodes = append(nodes, n.Issue)
	}
	return nodes
}
