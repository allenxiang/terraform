package httpclient

import (
	"crypto/tls"
	"net/http"

	cleanhttp "github.com/hashicorp/go-cleanhttp"
)

// New returns the DefaultPooledClient from the cleanhttp
// package that will also send a Terraform User-Agent string.
func New() *http.Client {
	transport := cleanhttp.DefaultPooledTransport()
	transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	cli := &http.Client{
		Transport: transport,
	}
	cli.Transport = &userAgentRoundTripper{
		userAgent: UserAgentString(),
		inner:     cli.Transport,
	}
	return cli
}
