package config

type ContextT struct {
    Meta struct {
        Cluster string `yaml:"cluster"`
        User string `yaml:"user"`
    } `yaml:"context"`
    Name string `yaml:"name"`
}

type ContextTAr []ContextT
