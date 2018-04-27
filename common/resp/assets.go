package resp

// Uploader ..
type Uploader struct {
	Login             string `json:"login"`
	ID                uint   `json:"id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

// Asset ..
type Asset struct {
	URL                string   `json:"url"`
	BrowserDownloadURL string   `json:"browser_download_url"`
	ID                 uint     `json:"id"`
	Name               string   `json:"name"`
	Label              string   `json:"label"`
	State              string   `json:"state"`
	ContentType        string   `json:"content_type"`
	Size               uint     `json:"size"`
	CreatedAt          string   `json:"created_at"`
	UpdatedAt          string   `json:"updated_at"`
	Uploader           Uploader `json:"uploader"`
}
