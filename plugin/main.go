package main

import (
	"context"

	"github.com/nori-io/common/v3/pkg/domain/config"
	em "github.com/nori-io/common/v3/pkg/domain/enum/meta"
	"github.com/nori-io/common/v3/pkg/domain/logger"
	"github.com/nori-io/common/v3/pkg/domain/meta"
	p "github.com/nori-io/common/v3/pkg/domain/plugin"
	"github.com/nori-io/common/v3/pkg/domain/registry"
	m "github.com/nori-io/common/v3/pkg/meta"
	"github.com/nori-io/http/internal/server"
	"github.com/nori-io/http/pkg"
)

var (
	Plugin p.Plugin = plugin{}
)

type plugin struct {
	server *server.Server
	config conf
}

type conf struct {
	port config.UInt
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
		Interface: pkg.HttpInterface,
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
		port: config.UInt("port", "http server port"),
	}

	p.server = server.New()
	// todo: take port from config
	// todo: integrate EventEmitter interface

	return nil
}

func (p plugin) Start(ctx context.Context, registry registry.Registry) error {
	return p.server.Start(p.config.port())
}

func (p plugin) Stop(ctx context.Context, registry registry.Registry) error {
	return p.server.Shutdown(ctx)
}
