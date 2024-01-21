package model

type Variable struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}

type Metadata struct {
	Name        string      `yaml:"name"`
	Description string      `yaml:"description"`
	DataFile    string      `yaml:"data_file"`
	DataBaseUrl string      `yaml:"data_base_url_file"`
	Arguments   []*Variable `yaml:"arguments"`
	Variables   []*Variable `yaml:"variables"`
}
