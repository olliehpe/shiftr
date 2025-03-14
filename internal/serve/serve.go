package serve

import (
	"fmt"
	"github.com/olliehpe/shiftr/internal"
	"log"
	"net/http"
)

func StartStaticFileServer(config *internal.Config) {
	fs := http.FileServer(http.Dir(fmt.Sprintf("./%s", config.DataFolder)))
	http.Handle("/", fs)

	port := fmt.Sprintf(":%s", config.Port)
	log.Printf("File Server started and listening on %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
