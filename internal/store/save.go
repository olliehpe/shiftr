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
				margin-top:80px;
				background-color: #bed3fb;
			}
			* {
				font-family: courier;
			}
			.source {
				padding-top:5px;
				padding-bottom:5px;
			}
			.logo {
				text-align: center;
				padding-bottom:30px;
			}
			h2 {
				margin:-20px;
				padding:0px;
				padding-bottom:50px;
			}
			.footer{
				margin-top:100px;
				text-align: center;
				font-size:14px;
			}
		</style>
	</head>
	<body>
		<div class="logo">
			<pre>
  _________.__    .__  _____  __                
 /   _____/|  |__ |__|/ ____\/  |_  ___________ 
 \_____  \ |  |  \|  \   __\\   __\/ __ \_  __ \
 /        \|   Y  \  ||  |   |  | \  ___/|  | \/
/_______  /|___|  /__||__|   |__|  \___  >__|   
        \/      \/                     \/
			</pre>
			<h2>Cached Endpoints</h2>
		</div>
		
		<hr>
			%s
		<div class="footer">
			<p>Last refresh: %s</p>
			<p>Refresh Interval: %d mins</p>
		</div>
	</body>
</html>
`
	htmlSources := ""
	htmlSource := `<div class="source">
	<h3>%s</h3>
	<p><b>Remote URL:</b> %s<br/>
	<b>Local Endpoint:</b> <a href="%s">/%s</a></p>
</div>
<hr>
`

	for _, source := range sources {
		url := source.Url
		if len(source.Url) > 150 {
			url = source.Url[:150] + "..."
		}
		htmlSourceEdit := fmt.Sprintf(htmlSource, source.Name, url, source.ServerFilename, source.ServerFilename)
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
