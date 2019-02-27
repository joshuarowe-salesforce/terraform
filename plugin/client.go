package plugin

import (
	"os"
	"os/exec"
	"time"

	hclog "github.com/hashicorp/go-hclog"
	plugin "github.com/hashicorp/go-plugin"
	"github.com/hashicorp/terraform/plugin/discovery"
)

// EnabledDebugging can be set to true to enable attaching to a debugger without timing out.
var EnableDebugging = false

// ClientConfig returns a configuration object that can be used to instantiate
// a client for the plugin described by the given metadata.
func ClientConfig(m discovery.PluginMeta) *plugin.ClientConfig {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:   "plugin",
		Level:  hclog.Trace,
		Output: os.Stderr,
	})

	clientConfig := &plugin.ClientConfig{
		Cmd:              exec.Command(m.Path),
		HandshakeConfig:  Handshake,
		Managed:          true,
		Plugins:          PluginMap,
		Logger:           logger,
		ConnectionConfig: plugin.DefaultConnectionConfig(),
	}

	if EnableDebugging {
		clientConfig.ConnectionConfig.EnableKeepAlive = false
		clientConfig.ConnectionConfig.ConnectionWriteTimeout = 365 * 84600 * time.Second
	}

	return clientConfig
}

// Client returns a plugin client for the plugin described by the given metadata.
func Client(m discovery.PluginMeta) *plugin.Client {
	return plugin.NewClient(ClientConfig(m))
}
