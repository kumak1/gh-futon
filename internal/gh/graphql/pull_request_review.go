package graphql

import graphql "github.com/cli/shurcooL-graphql"

type (
	PullRequestReviewQuery struct {
		User struct {
			ContributionsCollection struct {
				PullRequestReviewContributions pullRequestContributions `graphql:"pullRequestReviewContributions(first: $first)"`
			} `graphql:"contributionsCollection(from: $from, to: $to)"`
		} `graphql:"user(login: $username)"`
	}
)

func (p PullRequestReviewQuery) execQuery(variables map[string]interface{}) paginateQuery {
	query := PullRequestReviewQuery{}
	if err := graphClient.Query("PullRequestReviewQuery", &query, variables); err != nil {
		panic(err)
	}
	return query
}

func (p PullRequestReviewQuery) nextQuery() interface{} {
	return PullRequestReviewNextQuery{}
}

func (p PullRequestReviewQuery) hasNextPage() graphql.Boolean {
	return p.User.ContributionsCollection.PullRequestReviewContributions.PageInfo.HasNextPage
}

func (p PullRequestReviewQuery) endCursor() graphql.String {
	return p.User.ContributionsCollection.PullRequestReviewContributions.PageInfo.EndCursor
}

func (p PullRequestReviewQuery) nodes() []NodeInfo {
	var nodes []NodeInfo
	for _, n := range p.User.ContributionsCollection.PullRequestReviewContributions.Nodes {
		nodes = append(nodes, n.PullRequest)
	}
	return nodes
}
