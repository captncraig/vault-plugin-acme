package main

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/hashicorp/vault/helper/pluginutil"
	"github.com/hashicorp/vault/logical"
	"github.com/hashicorp/vault/logical/plugin"
)

func main() {

	apiClientMeta := &pluginutil.APIClientMeta{}
	flags := apiClientMeta.FlagSet()
	flags.Parse(os.Args[1:])
	ioutil.WriteFile("out.txt", []byte(strings.Join(os.Args, "*")), 0644)
	tlsConfig := apiClientMeta.GetTLSConfig()
	tlsProviderFunc := pluginutil.VaultPluginTLSProvider(tlsConfig)

	if err := plugin.Serve(&plugin.ServeOpts{
		BackendFactoryFunc: Factory,
		TLSProviderFunc:    tlsProviderFunc,
	}); err != nil {
		log.Fatal(err)
	}
}

func Factory(ctx context.Context, c *logical.BackendConfig) (logical.Backend, error) {
	// TODO
	return nil, nil
}
