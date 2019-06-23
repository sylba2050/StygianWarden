package analysis

import (
	"errors"
	"strings"
)

func IsRedirectTarget(origin, config string) bool {
	o := strings.Split(origin, "/")
	c := strings.Split(config, "/")

	if len(o) < len(c) {
		return false
	}

	for i := 0; i < len(o); i++ {
		if o[i] != c[i] {
			return false
		}
	}

	return true
}

func GetConfigIdx(origin string, configs []string) (int, error) {
	for i := 0; i < len(configs); i++ {
		if IsRedirectTarget(origin, configs[i]) {
			return i, nil
		}
	}

	return -1, errors.New("Not target")
}

func GetRedirectPath(origin, config string, isRemove bool) (string, error) {
	if isRemove {
		o := strings.Split(origin, "/")
		c := strings.Split(config, "/")

		for i := 0; i < len(c); i++ {
			o = o[1:]
		}

		return strings.Join(o, "/"), nil
	}

	return origin, nil
}
