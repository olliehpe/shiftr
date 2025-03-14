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
			compact, err := store.Compact(&data)
			if err != nil {
				log.Printf("Error compacting data: %v", err)
			}
			if err := store.SaveFile(fileName, compact, config.DataFolder); err != nil {
				log.Printf("Error saving file: %s", err)
			}
		}

		// pause
		internal.MemUsage()
		log.Println("Sleeping...")
		time.Sleep(time.Duration(config.RefreshInterval) * time.Minute)
	}
}
