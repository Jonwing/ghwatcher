package ghwatcher

type PushPayload struct {
	BaseRef string         `json:"base_ref"`
	Ref     string         `json:"ref"`
	Before  string         `json:"befor"`
	After   string         `json:"after"`
	Created bool           `json:"created"`
	Deleted bool           `json:"deleted`
	Forced  bool           `json:"forced"`
	Compare string         `json:"compare"`
	Commits []GithubCommit `json:"commits"`
}

type GithubCommit struct {
	ID        string     `json:"id"`
	TreeID    string     `json:"tree_id"`
	Distinct  bool       `json:"distinct"`
	Message   string     `json:"message"`
	Timestamp string     `json:"timestamp"`
	Url       string     `json:"url"`
	Author    GithubUser `json:"author"`
	Committer GithubUser `json:"committer"`
	Added     []string   `json:"added"`
	Removed   []string   `json:"removed"`
	Modified  []string   `json:"modified"`
}

type GithubUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
