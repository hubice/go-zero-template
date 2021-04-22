package config

import (
	"github.com/tal-tech/go-zero/rest"
)

type Config struct {
	configz.CommonConf
	rest.RestConf
}
