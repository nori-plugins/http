package main

import (
	"context"
	"fmt"

	"github.com/nori-plugins/http/internal/router"

	"github.com/nori-io/common/v5/pkg/domain/config"
	em "github.com/nori-io/common/v5/pkg/domain/enum/meta"
	"github.com/nori-io/common/v5/pkg/domain/event"
	"github.com/nori-io/common/v5/pkg/domain/logger"
	"github.com/nori-io/common/v5/pkg/domain/meta"
	p "github.com/nori-io/common/v5/pkg/domain/plugin"
	"github.com/nori-io/common/v5/pkg/domain/registry"
	m "github.com/nori-io/common/v5/pkg/meta"
	"github.com/nori-io/interfaces/nori/http/v2"
	"github.com/nori-plugins/http/internal/server"
)

func New() p.Plugin {
	return &plugin{}
}

type plugin struct {
	server *server.Server
	config conf
	log    logger.FieldLogger
}

type conf struct {
	host config.String
	port config.String
}

func (p *plugin) Init(ctx context.Context, config config.Config, log logger.FieldLogger) error {
	p.config = conf{
		host: config.String("host", "http server host"),
		port: config.String("port", "http server port"),
	}
	p.log = log

	return nil
}

func (p *plugin) Instance() interface{} {
	return p.getInstance()
}

func (p *plugin) getInstance() http.Router {
	return p.server.Router()
}

func (p *plugin) Meta() meta.Meta {
	return m.Meta{
		ID: m.ID{
			ID:      meta.PluginID("nori/http/HTTP"),
			Version: "0.1.0",
		},
		Author: m.Author{
			Name: "Nori Authors",
			URL:  "https://nori.io",
		},
		Dependencies: nil,
		Description: m.Description{
			Title:       "Nori HTTP Interface",
			Description: "Official implementation of nori/http/HTTP interface",
		},
		Interface: http.RouterInterface,
		License: []meta.License{m.License{
			Title: "",
			Type:  em.Apache2_0,
			URL:   "http://www.apache.org/licenses/",
		}},
		Links: nil,
		Repository: m.Repository{
			Type: em.Git,
			URL:  "github.com/nori-plugins/http",
		},
		Tags: []string{"nori", "http"},
	}
}

func (p *plugin) Start(ctx context.Context, registry registry.Registry) error {
	addr := fmt.Sprintf("%s:%s", p.config.host(), p.config.port())
	p.log.Info(fmt.Sprintf("http addr %s", addr))
	router := router.NewRouter()
	p.server = server.NewServer(addr, router)
	return nil
}

func (p *plugin) Stop(ctx context.Context, registry registry.Registry) error {
	return p.server.Shutdown(ctx)
}

func (p *plugin) Subscribe(emitter event.EventEmitter) {
	ch1, _ := emitter.On("/nori/plugins/started")

	go func() {
		for {
			select {
			case <-ch1:
				err := p.server.Start()
				if err != nil {
					p.log.Error(err.Error())
				}
			}
		}
	}()
}
