package graphql

import graphql "github.com/cli/shurcooL-graphql"

type (
	pullRequestReviewQuery struct {
		User struct {
			ContributionsCollection struct {
				PullRequestReviewContributions pullRequestContributions `graphql:"pullRequestReviewContributions(first: $first)"`
			} `graphql:"contributionsCollection(from: $from, to: $to)"`
		} `graphql:"user(login: $username)"`
	}
)

func (p pullRequestReviewQuery) execQuery(variables map[string]interface{}) (paginateQuery, error) {
	query := pullRequestReviewQuery{}
	err := graphClient.Query("pullRequestReview", &query, variables)
	return query, err
}

func (p pullRequestReviewQuery) nextQuery() interface{} {
	return pullRequestReviewNextQuery{}
}

func (p pullRequestReviewQuery) hasNextPage() graphql.Boolean {
	return p.User.ContributionsCollection.PullRequestReviewContributions.PageInfo.HasNextPage
}

func (p pullRequestReviewQuery) endCursor() graphql.String {
	return p.User.ContributionsCollection.PullRequestReviewContributions.PageInfo.EndCursor
}

func (p pullRequestReviewQuery) nodes() []NodeInfo {
	var nodes []NodeInfo
	for _, n := range p.User.ContributionsCollection.PullRequestReviewContributions.Nodes {
		nodes = append(nodes, n.PullRequest)
	}
	return nodes
}
