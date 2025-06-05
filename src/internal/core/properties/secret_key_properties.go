package core_properties

import (
	"github.com/golibs-starter/golib/config"
)

type SecretKeyProperties struct {
	SecretKey string `json:"secret_key"`
}

func (s SecretKeyProperties) Prefix() string {
	return "app"
}

func NewSecretKeyProperties(loader config.Loader) (*SecretKeyProperties, error) {
	props := SecretKeyProperties{}
	err := loader.Bind(&props)

	return &props, err
}
