package configuration

import "github.com/margostino/owid-api/common"

type Configuration struct {
	MetadataPath      string `yaml:"metadata_path"`
	SchemaFile        string `yaml:"schema_file"`
	UrlMappingFile    string `yaml:"urls_mapping_file"`
	GithubAccessToken string `yaml:"github_access_token"`
}

func GetConfig() *Configuration {
	var configuration Configuration
	common.UnmarshalYaml("./configuration.yml", &configuration)
	return &configuration
}
