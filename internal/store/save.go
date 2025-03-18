package store

import (
	"fmt"
	"github.com/olliehpe/shiftr/internal"
	"os"
	"path/filepath"
	"time"
)

func CheckCreateDataFolder(dataFolder string) error {
	if _, err := os.Stat(dataFolder); os.IsNotExist(err) {
		if err := os.MkdirAll(dataFolder, 0755); err != nil {
			return err
		}
	}
	return nil
}

func SaveFile(name string, data []byte, folder string) error {
	target := filepath.Join(folder, name)
	if err := os.WriteFile(target, data, 0755); err != nil {
		return err
	}
	return nil
}

func WriteIndex(sources []internal.Source, folder string, refresh int) error {
	file := `<html>
	<head>
		<meta charset="utf-8" />
		<meta name="viewport" content="width=device-width, initial-scale=1" />
		<title>Shiftr - Caching Server</title>
		<style>
			body {
				margin:500px;
				margin-top:200px;
			}
			* {
				font-family: sans-serif;
			}	
		</style>
	</head>
	<body>
		<h2>Cached Endpoints</h2>
		<p>The following data sources are cached locally and available:</p>
			%s
		<p>Last refresh: %s</p>
		<p>Refresh Interval: %d mins</p>
	</body>
</html>
`
	htmlSources := ""
	htmlSource := `<div>
	<h4>%s</h4>
	<p>Remote URL: %s</p>
	<p>Local Endpoint: <a href="%s">/%s</a></p>
	<hr>
</div>
`

	for _, source := range sources {
		htmlSourceEdit := fmt.Sprintf(htmlSource, source.Name, source.Url, source.ServerFilename, source.ServerFilename)
		htmlSources += htmlSourceEdit
	}

	nowTime := time.Now().String()
	file = fmt.Sprintf(file, htmlSources, nowTime, refresh)
	target := filepath.Join(folder, "index.html")
	if err := os.WriteFile(target, []byte(file), 0755); err != nil {
		return err
	}

	return nil
}
