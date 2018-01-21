package resp

type Releases []Release

type Release struct {
	Url             string  `json:"url"`
	HtmlUrl         string  `json:"html_url"`
	AssetsUrl       string  `json:"assets_url"`
	UploadUrl       string  `json:"upload_url"`
	TarballUrl      string  `json:"tarball_url"`
	ZipballUrl      string  `json:"zipball_url"`
	ID              uint    `json:"id"`
	TagName         string  `json:"tag_name"`
	TargetCommitish string  `json:"target_commitish"`
	Name            string  `json:"Name"`
	Body            string  `json:"body"`
	Draft           bool    `json:"draft"`
	Prerelease      bool    `json:"prerelease"`
	CreatedAt       string  `json:"created_at"`
	PublishedAt     string  `json:"published_at"`
	Author          Author  `json:"author"`
	Assets          []Asset `json:"assets"`
}
