package core_properties

import "github.com/golibs-starter/golib/config"

type ChainProperties struct {
	Chains map[uint64]Chain `json:"chains"`
}

func (c ChainProperties) Prefix() string {
	return "app"
}

type Chain struct {
	Name    string
	ChainID uint64
	Rpc     []string
}

func NewChainProperties(loader config.Loader) (*ChainProperties, error) {
	props := ChainProperties{}
	err := loader.Bind(&props)
	return &props, err
}
