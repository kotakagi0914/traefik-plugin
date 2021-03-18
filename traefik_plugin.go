package plugin

import (
	"context"
	"fmt"
	"net/http"
)

// Config the plugin configuration.
type Config struct {
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{}
}

type Demo struct {
	next    http.Handler
	headers map[string]string
	name    string
}

func New(ctx context.Context, next http.Handler, config Config, name string) (*Demo, error) {
	return &Demo{
		next: next,
		name: name,
	}, nil
}

func (a *Demo) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("plugin served")

	rw.WriteHeader(http.StatusOK)
	rw.Header().Add("x-test", "a.name")
	return

	// a.next.ServeHTTP(rw, req)
}
