package graphql

import graphql "github.com/cli/shurcooL-graphql"

type (
	pullRequestQuery struct {
		User struct {
			ContributionsCollection struct {
				PullRequestContributions pullRequestContributions `graphql:"pullRequestContributions(first: $first)"`
			} `graphql:"contributionsCollection(from: $from, to: $to)"`
		} `graphql:"user(login: $username)"`
	}
)

func (p pullRequestQuery) execQuery(variable QueryVariable) (paginateQuery, error) {
	err := graphClient.Query("pullRequest", &p, variable.ToMap())
	return p, err
}

func (p pullRequestQuery) nextQuery() interface{} {
	return pullRequestNextQuery{}
}

func (p pullRequestQuery) hasNextPage() graphql.Boolean {
	return p.User.ContributionsCollection.PullRequestContributions.PageInfo.HasNextPage
}

func (p pullRequestQuery) endCursor() graphql.String {
	return p.User.ContributionsCollection.PullRequestContributions.PageInfo.EndCursor
}

func (p pullRequestQuery) nodes() []NodeInfo {
	var nodes []NodeInfo
	for _, n := range p.User.ContributionsCollection.PullRequestContributions.Nodes {
		nodes = append(nodes, n.PullRequest)
	}
	return nodes
}
