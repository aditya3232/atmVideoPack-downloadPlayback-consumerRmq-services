package consumer_download_playback

type Service interface {
	ConsumerQueueDownloadPlayback() (RmqConsumerDownloadPlayback, error)
}

type service struct {
	downloadPlaybackRepository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

// consume and save to db
func (s *service) ConsumerQueueDownloadPlayback() (RmqConsumerDownloadPlayback, error) {

	// consume queue
	newRmqConsumerDownloadPlayback, err := s.downloadPlaybackRepository.ConsumerQueueDownloadPlayback()
	if err != nil {
		return newRmqConsumerDownloadPlayback, err
	}

	return newRmqConsumerDownloadPlayback, nil

}
