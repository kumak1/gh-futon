package graphql

import graphql "github.com/cli/shurcooL-graphql"

type (
	issueQuery struct {
		User struct {
			ContributionsCollection struct {
				IssueContributions issueContributions `graphql:"issueContributions(first: $first)"`
			} `graphql:"contributionsCollection(from: $from, to: $to)"`
		} `graphql:"user(login: $username)"`
	}
)

func (p issueQuery) execQuery(variables map[string]interface{}) (paginateQuery, error) {
	err := graphClient.Query("issue", &p, variables)
	return p, err
}

func (p issueQuery) nextQuery() interface{} {
	return issueNextQuery{}
}

func (p issueQuery) hasNextPage() graphql.Boolean {
	return p.User.ContributionsCollection.IssueContributions.PageInfo.HasNextPage
}

func (p issueQuery) endCursor() graphql.String {
	return p.User.ContributionsCollection.IssueContributions.PageInfo.EndCursor
}

func (p issueQuery) nodes() []NodeInfo {
	var nodes []NodeInfo
	for _, n := range p.User.ContributionsCollection.IssueContributions.Nodes {
		nodes = append(nodes, n.Issue)
	}
	return nodes
}
