package main

import (
	"context"

	"github.com/nori-io/common/v3/pkg/domain/event"

	"github.com/nori-io/common/v3/pkg/domain/config"
	em "github.com/nori-io/common/v3/pkg/domain/enum/meta"
	"github.com/nori-io/common/v3/pkg/domain/logger"
	"github.com/nori-io/common/v3/pkg/domain/meta"
	p "github.com/nori-io/common/v3/pkg/domain/plugin"
	"github.com/nori-io/common/v3/pkg/domain/registry"
	m "github.com/nori-io/common/v3/pkg/meta"
	"github.com/nori-io/http/internal/server"
	"github.com/nori-io/interfaces/nori/http"
)

var (
	Plugin p.Plugin = plugin{}
)

type plugin struct {
	server *server.Server
	config conf
}

type conf struct {
	port config.String
}

func (p plugin) Meta() meta.Meta {
	return m.Meta{
		ID:     m.ID{},
		Author: m.Author{
			//
		},
		Dependencies: nil,
		Description:  m.Description{
			//
		},
		Interface: http.HttpInterface,
		License:   nil,
		Links:     nil,
		Repository: m.Repository{
			Type: em.Git,
			URL:  "github.com/nori-io/http",
		},
		Tags: nil,
	}
}

func (p plugin) Instance() interface{} {
	return p.server
}

func (p plugin) Init(ctx context.Context, config config.Config, log logger.FieldLogger) error {
	p.config = conf{
		port: config.String("port", "http server port"),
	}

	p.server = server.New(p.config.port())

	return nil
}

func (p plugin) Start(ctx context.Context, registry registry.Registry) error {
	return nil
}

func (p plugin) Stop(ctx context.Context, registry registry.Registry) error {
	return nil
}

func (p plugin) Subscribe(emitter event.EventEmitter) {
	ch1, _ := emitter.On("/nori/plugins/started")
	ch2, _ := emitter.On("/nori/plugins/stopped")

	go func() {
		for {
			select {
			case <-ch1:
				p.server.Start()
			case <-ch2:
				p.server.Shutdown(context.Background())
			}
		}
	}()

}
