package graphql

import graphql "github.com/cli/shurcooL-graphql"

type (
	issueNextQuery struct {
		User struct {
			ContributionsCollection struct {
				IssueContributions issueContributions `graphql:"issueContributions(first: $first, after: $after)"`
			} `graphql:"contributionsCollection(from: $from, to: $to)"`
		} `graphql:"user(login: $username)"`
	}
)

func (p issueNextQuery) execQuery(variable QueryVariable) (paginateQuery, error) {
	err := graphClient.Query("issueNext", &p, variable.ToMap())
	return p, err
}

func (p issueNextQuery) nextQuery() interface{} {
	return issueNextQuery{}
}

func (p issueNextQuery) hasNextPage() graphql.Boolean {
	return p.User.ContributionsCollection.IssueContributions.PageInfo.HasNextPage
}

func (p issueNextQuery) endCursor() graphql.String {
	return p.User.ContributionsCollection.IssueContributions.PageInfo.EndCursor
}

func (p issueNextQuery) nodes() []NodeInfo {
	var nodes []NodeInfo
	for _, n := range p.User.ContributionsCollection.IssueContributions.Nodes {
		nodes = append(nodes, n.Issue)
	}
	return nodes
}
