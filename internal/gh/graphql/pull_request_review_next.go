package graphql

import graphql "github.com/cli/shurcooL-graphql"

type (
	pullRequestReviewNextQuery struct {
		User struct {
			ContributionsCollection struct {
				PullRequestReviewContributions pullRequestContributions `graphql:"pullRequestReviewContributions(first: $first, after: $after)"`
			} `graphql:"contributionsCollection(from: $from, to: $to)"`
		} `graphql:"user(login: $username)"`
	}
)

func (p pullRequestReviewNextQuery) execQuery(variables map[string]interface{}) (paginateQuery, error) {
	err := graphClient.Query("pullRequestReviewNext", &p, variables)
	return p, err
}

func (p pullRequestReviewNextQuery) nextQuery() interface{} {
	return pullRequestReviewNextQuery{}
}

func (p pullRequestReviewNextQuery) hasNextPage() graphql.Boolean {
	return p.User.ContributionsCollection.PullRequestReviewContributions.PageInfo.HasNextPage
}

func (p pullRequestReviewNextQuery) endCursor() graphql.String {
	return p.User.ContributionsCollection.PullRequestReviewContributions.PageInfo.EndCursor
}

func (p pullRequestReviewNextQuery) nodes() []NodeInfo {
	var nodes []NodeInfo
	for _, n := range p.User.ContributionsCollection.PullRequestReviewContributions.Nodes {
		nodes = append(nodes, n.PullRequest)
	}
	return nodes
}
