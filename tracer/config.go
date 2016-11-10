package tracer

import "github.com/imdario/mergo"

// Config contains the configuration of the tracer e.g. to know where to
// contact the Zipkin endpoint.
type Config struct {
	// ZipkinEndpoint contains the url to the Zipkin endpoint (string)
	ZipkinHost string
	// ServiceName is the name/label that is used to identify the trace in Zipkin (string)
	ServiceName string
}

// DefaultConfig a default configuration for the tracer
func DefaultConfig() Config {
	return Config{"localhost:9411", "iris"}
}

// Merge merges the default with the given config and returns the result
func (c Config) Merge(cfg []Config) (config Config) {

	if len(cfg) > 0 {
		config = cfg[0]
		mergo.Merge(&config, c)
	} else {
		_default := c
		config = _default
	}

	return
}
