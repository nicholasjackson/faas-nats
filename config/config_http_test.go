package config

import (
	"testing"

	"github.com/matryer/is"
	"github.com/nicholasjackson/faas-nats/providers/http"
)

func TestParsesConfigHTTPProviderHCL(t *testing.T) {
	is := is.New(t)

	c, err := ParseHCLFile("../test_fixtures/providers/http_output.hcl")

	is.NoErr(err)               // error should have been nil
	is.Equal(1, len(c.Outputs)) // should have returned one input provider

	p, ok := c.Outputs["open_faas"].(*http.HTTPProvider)

	is.True(ok)
	is.Equal("http", p.Protocol)        // should have set the protocol
	is.Equal("192.168.1.123", p.Server) // should have set the server
	is.Equal(80, p.Port)                // should have set the port
	is.Equal("/", p.Path)               // should have set the path

	is.Equal("key", p.TLS.TLSClientKey)   // should have set the ca key
	is.Equal("cert", p.TLS.TLSClientCert) // should have set the ca cert
}
