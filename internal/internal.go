package internal

import (
	"fmt"
	"runtime"
)

type Config struct {
	Sources         []Source
	Port            string `yaml:"port"`
	RefreshInterval int    `yaml:"refresh_interval"`
	DataFolder      string `yaml:"data_folder"`
}

type Source struct {
	Name           string            `yaml:"name"`
	Url            string            `yaml:"url"`
	ServerFilename string            `yaml:"server_filename"`
	Headers        map[string]string `yaml:"headers"`
	BasicAuth      BasicAuth         `yaml:"basic_auth"`
	Enabled        bool              `yaml:"enabled"`
}

type BasicAuth struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func MemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tHeapAlloc = %v MiB", bToMb(m.HeapAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
