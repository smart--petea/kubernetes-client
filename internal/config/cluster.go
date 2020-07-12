package config

import (
    "fmt"
)

type ClusterT struct {
    Meta struct { 
        CertificateAuthority string `yaml:"certificate-authority"`
        Server string `yaml:"server"`
    } `yaml:"cluster"`
    Name string `yaml:"name"`
}

type ClusterTAr []ClusterT

func (clusters ClusterTAr) FindByName(name string) (clusterConfig *ClusterT, err error) {
    for _, cluster := range clusters {
        if cluster.Name == name {
            return &cluster, err
        }
    }

    return nil, fmt.Errorf("Cluster with name %s not found", name)
}
