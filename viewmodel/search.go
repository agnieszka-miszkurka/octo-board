package viewmodel

// Content is a type that will be dispatched to home page template.
type Content struct {
	Labels                        []string
	Label, Organization, Language string
	Uncommented, Unassigned       bool
	NextPage                      int
	PrevPage                      int
	Issues                        []GithubIssue
}

// GithubIssue is a type that holds needed values of Github issue.
type GithubIssue struct {
	Title          string
	Repo           string
	HTMLURL        string
	Number         int
	Body           string
	CommentsNumber int
}
