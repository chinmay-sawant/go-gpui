package music

// ovResponse is the Openverse search envelope.
type ovResponse struct {
	Results []ovResult `json:"results"`
}

// ovResult is one Openverse audio record.
type ovResult struct {
	Title          string `json:"title"`
	Creator        string `json:"creator"`
	License        string `json:"license"`
	LicenseVersion string `json:"license_version"`
	URL            string `json:"url"`
	Filetype       string `json:"filetype"`
	Filesize       int    `json:"filesize"`
}

// playable reports an MP3 record with a direct URL small enough to decode.
func (r ovResult) playable() bool {
	if r.URL == "" {
		return false
	}

	if r.Filetype != "mp3" && r.Filetype != "mp32" {
		return false
	}

	return r.Filesize <= 0 || r.Filesize <= maxClipBytes
}
