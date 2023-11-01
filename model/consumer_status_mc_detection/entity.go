package consumer_status_mc_detection

// json di struct ini disesuaikan dengan key payload rmq
type RmqConsumerStatusMcDetection struct {
	TidID         *int   `json:"tid_id"`
	DateTime      string `json:"date_time"`
	StatusSignal  string `json:"status_signal"`
	StatusStorage string `json:"status_storage"`
	StatusRam     string `json:"status_ram"`
	StatusCpu     string `json:"status_cpu"`
}
