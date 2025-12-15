// Package review defines data models used by the review-bot.
package review

// Response represents the response payload produced by the review-bot.
type Response struct {
	Summary  Summary    `json:"summary"`
	Comments []Comments `json:"comments"`
}

// Summary is Overall Evaluation of target pull-request.
type Summary struct {
	OverallEvaluation string `json:"overallEvaluation"`
	Score             int    `json:"score"`
	SummaryJa         string `json:"summaryJa"`
}

// Comments is details of review.
type Comments struct {
	Severity      string `json:"severity"`
	Category      string `json:"category"`
	Title         string `json:"title"`
	Detail        string `json:"detail"`
	TargetSnippet string `json:"targetSnippet"`
	Suggestion    string `json:"suggestion"`
}
