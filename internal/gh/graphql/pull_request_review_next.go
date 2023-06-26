package graphql

import graphql "github.com/cli/shurcooL-graphql"

type (
	PullRequestReviewNextQuery struct {
		User struct {
			ContributionsCollection struct {
				PullRequestReviewContributions pullRequestContributions `graphql:"pullRequestReviewContributions(first: $first, after: $after)"`
			} `graphql:"contributionsCollection(from: $from, to: $to)"`
		} `graphql:"user(login: $username)"`
	}
)

func (p PullRequestReviewNextQuery) execQuery(variables map[string]interface{}) paginateQuery {
	query := PullRequestReviewNextQuery{}
	if err := graphClient.Query("PullRequestReviewNextQuery", &query, variables); err != nil {
		panic(err)
	}
	return query
}

func (p PullRequestReviewNextQuery) nextQuery() interface{} {
	return PullRequestReviewNextQuery{}
}

func (p PullRequestReviewNextQuery) hasNextPage() graphql.Boolean {
	return p.User.ContributionsCollection.PullRequestReviewContributions.PageInfo.HasNextPage
}

func (p PullRequestReviewNextQuery) endCursor() graphql.String {
	return p.User.ContributionsCollection.PullRequestReviewContributions.PageInfo.EndCursor
}

func (p PullRequestReviewNextQuery) nodes() []NodeInfo {
	var nodes []NodeInfo
	for _, n := range p.User.ContributionsCollection.PullRequestReviewContributions.Nodes {
		nodes = append(nodes, n.PullRequest)
	}
	return nodes
}
