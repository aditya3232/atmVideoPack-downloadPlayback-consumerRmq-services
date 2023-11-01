package consumer_download_playback

import (
	"encoding/json"
	"fmt"

	log_function "github.com/aditya3232/atmVideoPack-downloadPlayback-consumerRmq-services.git/log"
	"github.com/aditya3232/atmVideoPack-downloadPlayback-consumerRmq-services.git/model/add_download_playback_to_elastic"
	esv7 "github.com/elastic/go-elasticsearch/v7"
	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

type Repository interface {
	ConsumerQueueDownloadPlayback() (RmqConsumerDownloadPlayback, error)
}

type repository struct {
	db            *gorm.DB
	rabbitmq      *amqp.Connection
	elasticsearch *esv7.Client
}

func NewRepository(db *gorm.DB, rabbitmq *amqp.Connection, elasticsearch *esv7.Client) *repository {
	return &repository{db, rabbitmq, elasticsearch}
}

func (r *repository) ConsumerQueueDownloadPlayback() (RmqConsumerDownloadPlayback, error) {
	var rmqConsumerDownloadPlayback RmqConsumerDownloadPlayback

	// create channel
	channel, err := r.rabbitmq.Channel()
	if err != nil {
		return rmqConsumerDownloadPlayback, err
	}
	defer channel.Close()

	// consume queue
	msgs, err := channel.Consume(
		"DownloadPlaybackQueue", // name queue
		"",                      // Consumer name (empty for random name)
		true,                    // Auto-acknowledgment (set to true for auto-ack)
		false,                   // Exclusive
		false,                   // No-local
		false,                   // No-wait
		nil,                     // Arguments
	)

	if err != nil {
		return rmqConsumerDownloadPlayback, err
	}

	// get message
	for d := range msgs {
		newDownloadPlayback := rmqConsumerDownloadPlayback
		err := json.Unmarshal(d.Body, &newDownloadPlayback)
		if err != nil {
			return rmqConsumerDownloadPlayback, err
		}

		// add data download playback to elasticsearch with add_download_playback_to_elastic
		repoElastic := add_download_playback_to_elastic.NewRepository(r.elasticsearch)
		resultElastic, err := repoElastic.CreateElasticDownloadPlayback(
			add_download_playback_to_elastic.ElasticDownloadPlayback{
				Tid:             newDownloadPlayback.Tid,
				DateModified:    newDownloadPlayback.DateModified,
				DurationMinutes: newDownloadPlayback.DurationMinutes,
				FileSizeBytes:   newDownloadPlayback.FileSizeBytes,
				Filename:        newDownloadPlayback.Filename,
				Url:             newDownloadPlayback.Url,
			},
		)
		if err != nil {
			return rmqConsumerDownloadPlayback, err
		}
		// log result elastic
		log_function.Info(fmt.Sprintf("Result elastic: %v\n", resultElastic))

	}

	return rmqConsumerDownloadPlayback, nil

}
