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
	IsLeaves     bool   `yaml:"is_leave_original_url"`
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

func (d Data) GetEndpoints() []string {
	var res []string
	for _, p := range d.Proxies {
		res = append(res, p.Endpoints)
	}
	return res
}

func (d Data) GetAddr() []int {
	var res []int
	for _, p := range d.Proxies {
		res = append(res, p.RedirectAddr)
	}
	return res
}

func (d Data) GetIsLeaves() []bool {
	var res []bool
	for _, p := range d.Proxies {
		res = append(res, p.IsLeaves)
	}
	return res
}
