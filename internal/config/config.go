package config

type ConfigT struct {
    Clusters ClusterTAr `yaml:"clusters"`
    Contexts ContextTAr `yaml:"contexts"`
    Users UserTAr       `yaml:"users"`
}
