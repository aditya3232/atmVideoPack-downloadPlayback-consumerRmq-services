package add_download_playback_to_elastic

// ini entity data yg akan dikirim ke elastic

type ElasticDownloadPlayback struct {
	Tid             string `json:"tid"`
	DateModified    string `json:"date_modified"`
	DurationMinutes string `json:"duration_minutes"`
	FileSizeBytes   string `json:"file_size_bytes"`
	Filename        string `json:"filename"`
	Url             string `json:"url"`
}
