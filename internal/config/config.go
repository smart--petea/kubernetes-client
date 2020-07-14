package config

import (
    "io/ioutil"
    "os"
    "gopkg.in/yaml.v2"
)

type ConfigT struct {
    Clusters ClusterTAr `yaml:"clusters"`
    Contexts ContextTAr `yaml:"contexts"`
    Users UserTAr       `yaml:"users"`
}

var config *ConfigT

func GetConfig() (*ConfigT, error) {
    if config == nil {
        configBytes, err := ioutil.ReadFile(os.Getenv("CONFIG_PATH"))
        if err != nil {
            return nil, err
        }

        config = new(ConfigT)
        err = yaml.Unmarshal(configBytes, &config)
        if err != nil {
            return nil, err
        }
    }

    return config, nil
}
