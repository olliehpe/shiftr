package internal

type Config struct {
	Sources         []Source
	Port            string `yaml:"port"`
	RefreshInterval int    `yaml:"refresh_interval"`
	DataFolder      string `yaml:"data_folder"`
}

type Source struct {
	Name string `yaml:"name"`
	Url  string `yaml:"url"`

	ServerFilename string            `yaml:"server_filename"`
	Headers        map[string]string `yaml:"headers"`
	BasicAuth      BasicAuth         `yaml:"basic_auth"`
}

type BasicAuth struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
