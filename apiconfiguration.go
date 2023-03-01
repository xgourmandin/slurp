package slurp

type ApiConfiguration struct {
	Url                   string                  `yaml:"url"`
	Method                string                  `yaml:"method"`
	AuthConfig            AuthenticationConfig    `yaml:"auth"`
	PaginationConfig      PaginationConfiguration `yaml:"pagination"`
	DataConfig            DataConfiguration       `yaml:"data"`
	AdditionalHeaders     map[string]string       `yaml:"additional_headers"`
	AdditionalQueryParams map[string]string       `yaml:"additional_queryparams"`
	OutputConfig          OutputConfig            `yaml:"output"`
}

type DataConfiguration struct {
	DataType string `yaml:"type"`
	DataRoot string `yaml:"root"`
}

type PaginationConfiguration struct {
	PaginationType string `yaml:"type"`
	PageParam      string `yaml:"page_param"`
	LimitParam     string `yaml:"limit_param"`
	PageSize       int    `yaml:"page_size"`
	NextLinkPath   string `yaml:"next_link_path"`
}

type AuthenticationConfig struct {
	AuthType   string `yaml:"type"`
	InHeader   bool   `yaml:"in_header"`
	TokenEnv   string `yaml:"token_env"`
	TokenParam string `yaml:"token_param"`
}

type OutputConfig struct {
	OutputType string `yaml:"type"`
	FileName   string `yaml:"filename"`
	BucketName string `yaml:"bucket"`
	Project    string `yaml:"project"`
	Dataset    string `yaml:"dataset"`
	Table      string `yaml:"table"`
	Autodetect bool   `yaml:"autodetect"`
}
