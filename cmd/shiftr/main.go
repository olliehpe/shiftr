package main

import (
	"github.com/olliehpe/shiftr/internal"
	cnf "github.com/olliehpe/shiftr/internal/config"
	"github.com/olliehpe/shiftr/internal/fetch"
	"github.com/olliehpe/shiftr/internal/serve"
	"github.com/olliehpe/shiftr/internal/store"
	"log"
	"time"
)

var (
	version   = "Development build"
	goversion = "Unknown"
)

func main() {
	config := cnf.LoadConfig()

	// create data folder if not exist
	if err := store.CheckCreateDataFolder(config.DataFolder); err != nil {
		log.Fatalf("Error creating data folder %s: %v", config.DataFolder, err)
	}

	go serve.StartStaticFileServer(config)

	for {
		// get json payloads
		log.Println("Fetching data from sources...")
		payloads := map[string][]byte{}

		for _, source := range config.Sources {
			payloads[source.ServerFilename] = fetch.GetData(&source)
		}

		// save to files
		log.Println("Saving responses to data files...")
		for fileName, data := range payloads {
			if err := store.SaveFile(fileName, data, config.DataFolder); err != nil {
				log.Printf("Error saving file: %s", err)
			}
		}

		// create an index.html file so the cached resources are documented
		if err := store.WriteIndex(config.Sources, config.DataFolder, config.RefreshInterval); err != nil {
			log.Fatalf("Error writing index.html: %v", err)
		}

		// pause
		internal.MemUsage()
		log.Println("Sleeping...")
		time.Sleep(time.Duration(config.RefreshInterval) * time.Minute)
	}
}
