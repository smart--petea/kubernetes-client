package config

import (
    "fmt"
)

type UserT struct {
    Name string `yaml:"name"`
    Meta struct {
        ClientCertificate string `yaml:"client-certificate"`
        ClientKey string `yaml:"client-key"`
    } `yaml:"user"`
}

type UserTAr []UserT

func (users UserTAr) FindByName(name string) (*UserT, error) {
    for _, user := range users {
        if user.Name == name {
            return &user, nil
        }
    }

    return nil, fmt.Errorf("User with name=%s is not found", name)
}
