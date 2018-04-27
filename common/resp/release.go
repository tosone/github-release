package resp

// Releases ..
type Releases []Release

// Release ..
type Release struct {
	URL             string  `json:"url"`
	HTMLURL         string  `json:"html_url"`
	AssetsURL       string  `json:"assets_url"`
	UploadURL       string  `json:"upload_url"`
	TarballURL      string  `json:"tarball_url"`
	ZipballURL      string  `json:"zipball_url"`
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
