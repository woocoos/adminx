package graph

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/woocoos/adminx/ent"
	"github.com/woocoos/adminx/graph/generated"
	"github.com/woocoos/adminx/service/resource"
	"github.com/woocoos/adminx/service/ucenter"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type options struct {
	client   *ent.Client
	ucenter  *ucenter.Service
	resource *resource.Service
}

type Option func(*options)

func RegisterEntClient(c *ent.Client) Option {
	return func(o *options) {
		o.client = c
	}
}

func RegistryService(svcs ...any) Option {
	return func(o *options) {
		for _, svc := range svcs {
			switch svc.(type) {
			case *ucenter.Service:
				o.ucenter = svc.(*ucenter.Service)
			case *resource.Service:
				o.resource = svc.(*resource.Service)
			}
		}
	}
}

type Resolver struct {
	client   *ent.Client
	ucenter  *ucenter.Service
	resource *resource.Service
}

// NewSchema creates a graphql executable schema.
func NewSchema(opts ...Option) graphql.ExecutableSchema {
	opt := &options{}
	for _, o := range opts {
		o(opt)
	}
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{
			client:   opt.client,
			ucenter:  opt.ucenter,
			resource: opt.resource,
		},
	})
}
