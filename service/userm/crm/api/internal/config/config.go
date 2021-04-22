package config

import (
	"dark-work/common/configz"
	"github.com/tal-tech/go-zero/rest"
)

type Config struct {
	configz.CommonConf
	rest.RestConf
}
