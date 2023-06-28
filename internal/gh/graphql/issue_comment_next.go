package graphql

import graphql "github.com/cli/shurcooL-graphql"

type (
	issueCommentNextQuery struct {
		User struct {
			IssueComments issueComments `graphql:"issueComments(first: $first, orderBy: {field: UPDATED_AT, direction: DESC}, after: $after)"`
		} `graphql:"user(login: $username)"`
	}
)

func (p issueCommentNextQuery) execQuery(variable QueryVariable) (paginateQuery, error) {
	issueCommentVariable = variable
	err := graphClient.Query("issueCommentNext", &p, variable.ToMapForComment())
	return p, err
}

func (p issueCommentNextQuery) nextQuery() interface{} {
	return issueCommentNextQuery{}
}

func (p issueCommentNextQuery) hasNextPage() graphql.Boolean {
	if !p.User.IssueComments.PageInfo.HasNextPage {
		return false
	}

	createdAt := p.User.IssueComments.Nodes[0].CreatedAt
	return graphql.Boolean(createdAt.After(*issueCommentVariable.From) && createdAt.Before(*issueCommentVariable.To))
}

func (p issueCommentNextQuery) endCursor() graphql.String {
	return p.User.IssueComments.PageInfo.EndCursor
}

func (p issueCommentNextQuery) nodes() []NodeInfo {
	var nodes []NodeInfo
	for _, n := range p.User.IssueComments.Nodes {
		nodes = append(nodes, n.Issue)
	}
	return nodes
}
