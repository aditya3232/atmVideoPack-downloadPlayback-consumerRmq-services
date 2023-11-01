package consumer_download_playback

// json di struct ini disesuaikan dengan key payload rmq
type RmqConsumerDownloadPlayback struct {
	Tid             string `json:"tid"`
	DateModified    string `json:"date_modified"`
	DurationMinutes string `json:"duration_minutes"`
	FileSizeBytes   string `json:"file_size_bytes"`
	Filename        string `json:"filename"`
	Url             string `json:"url"`
}
