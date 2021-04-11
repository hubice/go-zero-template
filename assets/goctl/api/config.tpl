package config

import (
	"github.com/tal-tech/go-zero/rest"
    {{.authImport}}
)

type Config struct {
	configz.CommonConf
	rest.RestConf
	{{.auth}}
}
