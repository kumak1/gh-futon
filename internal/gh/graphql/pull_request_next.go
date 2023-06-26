package graphql

import graphql "github.com/cli/shurcooL-graphql"

type (
	PullRequestNextQuery struct {
		User struct {
			ContributionsCollection struct {
				PullRequestContributions pullRequestContributions `graphql:"pullRequestContributions(first: $first, after: $after)"`
			} `graphql:"contributionsCollection(from: $from, to: $to)"`
		} `graphql:"user(login: $username)"`
	}
)

func (p PullRequestNextQuery) execQuery(variables map[string]interface{}) paginateQuery {
	query := PullRequestNextQuery{}
	if err := graphClient.Query("PullRequestNextQuery", &query, variables); err != nil {
		panic(err)
	}
	return query
}

func (p PullRequestNextQuery) nextQuery() interface{} {
	return PullRequestNextQuery{}
}

func (p PullRequestNextQuery) hasNextPage() graphql.Boolean {
	return p.User.ContributionsCollection.PullRequestContributions.PageInfo.HasNextPage
}

func (p PullRequestNextQuery) endCursor() graphql.String {
	return p.User.ContributionsCollection.PullRequestContributions.PageInfo.EndCursor
}

func (p PullRequestNextQuery) nodes() []NodeInfo {
	var nodes []NodeInfo
	for _, n := range p.User.ContributionsCollection.PullRequestContributions.Nodes {
		nodes = append(nodes, n.PullRequest)
	}
	return nodes
}
