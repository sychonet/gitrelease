package model

// Config struct for application config
type Config struct {
	// VCS struct for version control systems
	VCS struct {
		Github struct {
			// Personal access token for https://github.com
			AccessToken string `yaml:"accessToken"`
		} `yaml:"github"`
		Gitlab struct {
			// Personal access token for https://gitlab.com
			AccessToken string `yaml:"accessToken"`
		} `yaml:"gitlab"`
	} `yaml:"vcs"`
}
