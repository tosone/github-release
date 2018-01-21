package resp

type Uploader struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	AvatarUrl         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	Url               string `json:"url"`
	HtmlUrl           string `json:"html_url"`
	FollowersUrl      string `json:"followers_url"`
	FollowingUrl      string `json:"following_url"`
	GistsUrl          string `json:"gists_url"`
	StarredUrl        string `json:"starred_url"`
	SubscriptionsUrl  string `json:"subscriptions_url"`
	OrganizationsUrl  string `json:"organizations_url"`
	ReposUrl          string `json:"repos_url"`
	EventsUrl         string `json:"events_url"`
	ReceivedEventsUrl string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

type Asset struct {
	Url                string   `json:"url"`
	BrowserDownloadUrl string   `json:"browser_download_url"`
	ID                 string   `json:"id"`
	Name               string   `json:"name"`
	Label              string   `json:"label"`
	State              string   `json:"state"`
	ContentType        string   `json:"content_type"`
	Size               uint     `json:"size"`
	CreatedAt          string   `json:"created_at"`
	UpdatedAt          string   `json:"updated_at"`
	Uploader           Uploader `json:"uploader"`
}
