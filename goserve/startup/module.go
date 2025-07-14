package startup

import (
	"context"

	"github.com/unusualcodeorg/goserve/api/auth"
	"github.com/unusualcodeorg/goserve/api/blog"
	"github.com/unusualcodeorg/goserve/api/user"
	"github.com/unusualcodeorg/goserve/arch/mongo"
	"github.com/unusualcodeorg/goserve/arch/network"
	"github.com/unusualcodeorg/goserve/arch/redis"
	"github.com/unusualcodeorg/goserve/config"
)

type Module network.Module[module]

type module struct {
	Context     context.Context
	Env         *config.Env
	DB          mongo.Database
	Store       redis.Store
	UserService user.Service
	AuthService auth.Service
	BlogService blog.Service
}

func (m *module) GetInstanse() *module {
	return m
}

func (m *module) Controllers() []network.Controller {
	return []network.Controller{
		auth.NewController(m.AuthenticationProvider(), m.AuthorizationProvider(), m.AuthService),
		user.NewController(m.AuthenticationProvider(), m.AuthorizationProvider(), m.UserService),
		blog.NewController(m.AuthenticationProvider(), m.AuthorizationProvider(), m.BlogService),
		author.NewController(m.AuthenticationProvider(), m.AuthorizationProvider(), author.NewService(m.DB, m.BlogService)),
		editor.NewController(m.AuthenticationProvider(), m.AuthorizationProvider(), editor.NewService(m.DB, m.UserService)),
		blogs.NewController(m.AuthenticationProvider(), m.AuthorizationProvider(), blogs.NewService(m.DB, m.Store)),
		contact.NewController(m.AuthenticationProvider(), m.AuthorizationProvider(), contact.NewService(m.DB)),
}
