package web

import (
	"strings"
	"time"

	"github.com/jwowillo/landgrab/app/api"
	"github.com/jwowillo/trim/application"
)

// New ...
func New(sd, h, sf string) *application.Application {
	clientConf, apiConf, staticConf := configs(sd, h, sf)
	app := application.NewWebWithConfig(
		clientConf,
		apiConf,
		staticConf,
	)
	app.SetAPI(api.New())
	return app.Application
}

// configs to use for the Application based on the host the trim.Application
// will be served on.
func configs(sd, h, sf string) (
	application.ClientConfig,
	application.APIConfig,
	application.StaticConfig,
) {
	clientConf := application.ClientDefault
	clientConf.Subdomain = sd
	clientConf.TemplateFolder = sf
	staticConf := application.StaticDefault
	staticConf.BaseFolder = sf
	if !strings.HasPrefix(h, "localhost") {
		clientConf.CacheDuration = time.Hour
		staticConf.CacheDuration = time.Hour
	} else {
		clientConf.CacheDuration = 0
		staticConf.CacheDuration = 0
	}
	staticConf.Include = []string{".js"}
	return clientConf, application.APIDefault, staticConf
}
