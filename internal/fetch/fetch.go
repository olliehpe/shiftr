package fetch

import (
	"github.com/olliehpe/shiftr/internal"
	"io"
	"log"
	"net/http"
)

func GetData(source *internal.Source) []byte {

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, source.Url, nil)
	if err != nil {
		log.Printf("Error creating GET request for %s\n", source.Url)
	}

	// add headers
	for k, v := range source.Headers {
		req.Header.Add(k, v)
	}

	// if basic auth provided, add to request
	if source.BasicAuth.Username != "" && source.BasicAuth.Password != "" {
		req.SetBasicAuth(source.BasicAuth.Username, source.BasicAuth.Password)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Printf("Error fetching %s: %s\n", source.Url, err)
	}

	if res.StatusCode != 200 {
		log.Printf("Error fetching %s: %s\n", source.Url, res.Status)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error reading %s: %s\n", source.Url, err)
	}

	return data
}
