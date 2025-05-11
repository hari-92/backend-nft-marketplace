package discovery_properties

import (
	"fmt"
	"github.com/golibs-starter/golib/config"
)

// NewServicesDiscoveryProperties create a new NewServicesDiscoveryProperties instance
func NewServicesDiscoveryProperties(loader config.Loader) (*ServicesProperties, error) {
	props := ServicesProperties{}
	err := loader.Bind(&props)
	fmt.Println("NewServicesDiscoveryProperties: ", props)
	return &props, err
}

// ServicesProperties represent statistics configuration properties
type ServicesProperties struct {
	Services map[string]Service `json:"services"`
}

type Service struct {
	Host string `json:"host"`
	Port uint64 `json:"port"`
}

// Prefix return yaml prefix of configuration
func (f ServicesProperties) Prefix() string {
	return "app"
}
