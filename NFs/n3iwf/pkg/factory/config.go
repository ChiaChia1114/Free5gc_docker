/*
 * N3IWF Configuration Factory
 */

package factory

import (
	"github.com/free5gc/logger_util"
	"github.com/free5gc/n3iwf/pkg/context"
)

const (
	N3IWF_EXPECTED_CONFIG_VERSION = "1.0.0"
)

type Config struct {
	Info          *Info               `yaml:"info"`
	Configuration *Configuration      `yaml:"configuration"`
	Logger        *logger_util.Logger `yaml:"logger"`
}

type Info struct {
	Version     string `yaml:"version,omitempty"`
	Description string `yaml:"description,omitempty"`
}

type Configuration struct {
	N3IWFInfo        context.N3IWFNFInfo        `yaml:"N3IWFInformation"`
	AMFSCTPAddresses []context.AMFSCTPAddresses `yaml:"AMFSCTPAddresses"`

	IKEBindAddr          string `yaml:"IKEBindAddress"`
	IPSecGatewayAddr     string `yaml:"IPSecInterfaceAddress"`
	GTPBindAddr          string `yaml:"GTPBindAddress"`
	TCPPort              uint16 `yaml:"NASTCPPort"`
	FQDN                 string `yaml:"FQDN"`                 // e.g. n3iwf.free5gc.org
	PrivateKey           string `yaml:"PrivateKey"`           // file path
	CertificateAuthority string `yaml:"CertificateAuthority"` // file path
	Certificate          string `yaml:"Certificate"`          // file path
	UEIPAddressRange     string `yaml:"UEIPAddressRange"`     // e.g. 10.0.1.0/24
	InterfaceMark        uint32 `yaml:"IPSecInterfaceMark"`   // must != 0, if not specified, random one
}

func (c *Config) GetVersion() string {
	if c.Info != nil && c.Info.Version != "" {
		return c.Info.Version
	}
	return ""
}
