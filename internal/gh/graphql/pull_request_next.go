package graphql

import graphql "github.com/cli/shurcooL-graphql"

type (
	pullRequestNextQuery struct {
		User struct {
			ContributionsCollection struct {
				PullRequestContributions pullRequestContributions `graphql:"pullRequestContributions(first: $first, after: $after)"`
			} `graphql:"contributionsCollection(from: $from, to: $to)"`
		} `graphql:"user(login: $username)"`
	}
)

func (p pullRequestNextQuery) execQuery(variables map[string]interface{}) (paginateQuery, error) {
	err := graphClient.Query("PullRequestNext", &p, variables)
	return p, err
}

func (p pullRequestNextQuery) nextQuery() interface{} {
	return pullRequestNextQuery{}
}

func (p pullRequestNextQuery) hasNextPage() graphql.Boolean {
	return p.User.ContributionsCollection.PullRequestContributions.PageInfo.HasNextPage
}

func (p pullRequestNextQuery) endCursor() graphql.String {
	return p.User.ContributionsCollection.PullRequestContributions.PageInfo.EndCursor
}

func (p pullRequestNextQuery) nodes() []NodeInfo {
	var nodes []NodeInfo
	for _, n := range p.User.ContributionsCollection.PullRequestContributions.Nodes {
		nodes = append(nodes, n.PullRequest)
	}
	return nodes
}
