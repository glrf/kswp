package kswp

import (
	"fmt"
	"io"
	"os"
)

type KubeConfig struct {
	Name string
	Path string
}
type Kswp struct {
	KubeConf string
	Configs  []KubeConfig
}

func (k Kswp) Swap(config string) error {
  kc, err := k.getConfig(config)
  if err != nil {
    return err
  }

  source, err := os.Open(kc.Path)
	if err != nil {
		return fmt.Errorf("Cannot read kubeconfig %s. \n %w", kc.Path, err)
	}
	defer source.Close()
	destination, err := os.Create(k.KubeConf)
	if err != nil {
		return fmt.Errorf("Cannot read kubeconfig %s. \n %w", kc.Path, err)
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return err
}


func (k Kswp) getConfig(config string)(KubeConfig, error) {
  var kc KubeConfig
	for _, c := range k.Configs {
		if config == c.Name {
			kc = c
		}
	}
	if kc.Name != config {
		return kc, fmt.Errorf("Cannot find kubeconfig with name %s.", config)
	}
  return kc, nil
}
