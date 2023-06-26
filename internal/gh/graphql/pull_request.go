package graphql

import graphql "github.com/cli/shurcooL-graphql"

type (
	PullRequestQuery struct {
		User struct {
			ContributionsCollection struct {
				PullRequestContributions pullRequestContributions `graphql:"pullRequestContributions(first: $first)"`
			} `graphql:"contributionsCollection(from: $from, to: $to)"`
		} `graphql:"user(login: $username)"`
	}
)

func (p PullRequestQuery) execQuery(variables map[string]interface{}) paginateQuery {
	query := PullRequestQuery{}
	if err := graphClient.Query("PullRequestQuery", &query, variables); err != nil {
		panic(err)
	}
	return query
}

func (p PullRequestQuery) nextQuery() interface{} {
	return PullRequestNextQuery{}
}

func (p PullRequestQuery) hasNextPage() graphql.Boolean {
	return p.User.ContributionsCollection.PullRequestContributions.PageInfo.HasNextPage
}

func (p PullRequestQuery) endCursor() graphql.String {
	return p.User.ContributionsCollection.PullRequestContributions.PageInfo.EndCursor
}

func (p PullRequestQuery) nodes() []NodeInfo {
	var nodes []NodeInfo
	for _, n := range p.User.ContributionsCollection.PullRequestContributions.Nodes {
		nodes = append(nodes, n.PullRequest)
	}
	return nodes
}
