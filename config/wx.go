package config

type Wx struct {
	AppID      string   `mapstructure:"app-id" json:"appID" yaml:"app-id"`                // AppID
	AppSecret  string   `mapstructure:"app-secret" json:"appSecret" yaml:"app-secret"`    // AppSecret
	TemplateID []string `mapstructure:"template-id" json:"templateID" yaml:"template-id"` // AppSecret
}
