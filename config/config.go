package config

import (
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
)

type Data struct {
	Addr    int     `yaml:"addr"`
	Proxies []Proxy `yaml:"proxies"`
}

type Proxy struct {
	RedirectAddr int    `yaml:"redirect_addr"`
	Endpoints    string `yaml:"endpoints"`
}

func Load(path string) Data {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var d Data
	err = yaml.Unmarshal(buf, &d)
	if err != nil {
		panic(err)
	}

	return d
}
