package gh

type LoginResponse struct {
	Login string
}

type EventResponse struct {
	Id   string
	Type string
	Repo struct {
		Name string
		Url  string
	}
	Payload struct {
		Action      string
		PullRequest struct {
			Title string
			URL   string `json:"html_url"`
		} `json:"pull_request"`
		Issue struct {
			Title string
			URL   string `json:"html_url"`
		}
		Review struct {
			URL string `json:"html_url"`
		}
		Comment struct {
			URL string `json:"html_url"`
		}
	}
}
