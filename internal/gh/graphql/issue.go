package graphql

import graphql "github.com/cli/shurcooL-graphql"

type (
	IssueQuery struct {
		User struct {
			ContributionsCollection struct {
				IssueContributions issueContributions `graphql:"issueContributions(first: $first)"`
			} `graphql:"contributionsCollection(from: $from, to: $to)"`
		} `graphql:"user(login: $username)"`
	}
)

func (p IssueQuery) execQuery(variables map[string]interface{}) paginateQuery {
	query := IssueQuery{}
	if err := graphClient.Query("IssueQuery", &query, variables); err != nil {
		panic(err)
	}
	return query
}

func (p IssueQuery) nextQuery() interface{} {
	return IssueNextQuery{}
}

func (p IssueQuery) hasNextPage() graphql.Boolean {
	return p.User.ContributionsCollection.IssueContributions.PageInfo.HasNextPage
}

func (p IssueQuery) endCursor() graphql.String {
	return p.User.ContributionsCollection.IssueContributions.PageInfo.EndCursor
}

func (p IssueQuery) nodes() []NodeInfo {
	var nodes []NodeInfo
	for _, n := range p.User.ContributionsCollection.IssueContributions.Nodes {
		nodes = append(nodes, n.Issue)
	}
	return nodes
}
