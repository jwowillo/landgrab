package web

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jwowillo/landgrab/api"
	"github.com/jwowillo/static"
	"github.com/jwowillo/trim/application"
	"github.com/jwowillo/trim/controller"
)

// New ...
func New(sd, h, sf string) *application.Application {
	static.SetLogger(log.New(os.Stderr, "", 0))
	build := sd
	if sd == "" {
		build = "build"
	}
	dl := static.NewGithubDownloader("jwowillo", "landgrab")
	if err := static.DartProject(dl, "landgrab", build); err != nil {
		return nil
	}
	clientConf, apiConf, staticConf := configs(sd, h, sf)
	app := application.NewWebWithConfig(
		clientConf,
		apiConf,
		staticConf,
	)
	app.AddController(controller.NewHome(
		"/api",
		clientConf.HomePath,
		clientConf.CacheDuration,
		clientConf.CacheMaxSize,
	))
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
	clientConf.HomePath = filepath.Join(sf, "index.html")
	staticConf := application.StaticDefault
	staticConf.BaseFolder = sf
	if !strings.HasPrefix(h, "localhost") {
		clientConf.CacheDuration = time.Hour
		staticConf.CacheDuration = time.Hour
	} else {
		clientConf.CacheDuration = 0
		staticConf.CacheDuration = 0
	}
	staticConf.IncludeExtensions = []string{".js"}
	return clientConf, application.APIDefault, staticConf
}
