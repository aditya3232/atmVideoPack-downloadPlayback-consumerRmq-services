package add_download_playback_to_elastic

import (
	"strings"

	esv7 "github.com/elastic/go-elasticsearch/v7"
)

type Repository interface {
	CreateElasticDownloadPlayback(elasticDownloadPlayback ElasticDownloadPlayback) (ElasticDownloadPlayback, error)
}

type repository struct {
	elasticsearch *esv7.Client
}

func NewRepository(elasticsearch *esv7.Client) *repository {
	return &repository{elasticsearch}
}

func (r *repository) CreateElasticDownloadPlayback(elasticDownloadPlayback ElasticDownloadPlayback) (ElasticDownloadPlayback, error) {

	// Menggunakan library "github.com/elastic/go-elasticsearch" untuk melakukan operasi penyimpanan
	// Gantilah `indexName` dengan nama index Elasticsearch yang sesuai
	indexName := "download_playback_index"

	// Anda dapat membuat body dokumen yang akan disimpan di Elasticsearch
	// Misalnya, jika Anda ingin menyimpan data deteksi manusia yang diberikan sebagai JSON:
	body := []byte(`{
		"tid": "` + elasticDownloadPlayback.Tid + `",
		"date_modified": "` + elasticDownloadPlayback.DateModified + `",
		"duration_minutes": "` + elasticDownloadPlayback.DurationMinutes + `",
		"file_size_bytes": "` + elasticDownloadPlayback.FileSizeBytes + `",
		"filename": "` + elasticDownloadPlayback.Filename + `",
		"url": "` + elasticDownloadPlayback.Url + `"
		
	}`)

	// Mengirimkan data ke Elasticsearch untuk disimpan
	_, err := r.elasticsearch.Index(indexName, strings.NewReader(string(body)))
	if err != nil {
		return elasticDownloadPlayback, err
	}

	// Jika operasi berhasil, Anda dapat mengembalikan data yang sama yang Anda terima sebagai argumen.
	return elasticDownloadPlayback, nil

}
