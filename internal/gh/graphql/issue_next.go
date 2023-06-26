package graphql

import graphql "github.com/cli/shurcooL-graphql"

type (
	IssueNextQuery struct {
		User struct {
			ContributionsCollection struct {
				IssueContributions issueContributions `graphql:"issueContributions(first: $first, after: $after)"`
			} `graphql:"contributionsCollection(from: $from, to: $to)"`
		} `graphql:"user(login: $username)"`
	}
)

func (p IssueNextQuery) execQuery(variables map[string]interface{}) paginateQuery {
	query := IssueNextQuery{}
	if err := graphClient.Query("IssueNextQuery", &query, variables); err != nil {
		panic(err)
	}
	return query
}

func (p IssueNextQuery) nextQuery() interface{} {
	return IssueNextQuery{}
}

func (p IssueNextQuery) hasNextPage() graphql.Boolean {
	return p.User.ContributionsCollection.IssueContributions.PageInfo.HasNextPage
}

func (p IssueNextQuery) endCursor() graphql.String {
	return p.User.ContributionsCollection.IssueContributions.PageInfo.EndCursor
}

func (p IssueNextQuery) nodes() []NodeInfo {
	var nodes []NodeInfo
	for _, n := range p.User.ContributionsCollection.IssueContributions.Nodes {
		nodes = append(nodes, n.Issue)
	}
	return nodes
}
