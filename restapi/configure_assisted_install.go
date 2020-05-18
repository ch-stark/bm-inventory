// Code generated by go-swagger; DO NOT EDIT.

package restapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/filanov/bm-inventory/restapi/operations"
	"github.com/filanov/bm-inventory/restapi/operations/installer"
)

type contextKey string

const AuthKey contextKey = "Auth"

//go:generate mockery -name InstallerAPI -inpkg

// InstallerAPI
type InstallerAPI interface {
	DeregisterCluster(ctx context.Context, params installer.DeregisterClusterParams) middleware.Responder
	DeregisterHost(ctx context.Context, params installer.DeregisterHostParams) middleware.Responder
	DisableHost(ctx context.Context, params installer.DisableHostParams) middleware.Responder
	DownloadClusterFiles(ctx context.Context, params installer.DownloadClusterFilesParams) middleware.Responder
	DownloadClusterISO(ctx context.Context, params installer.DownloadClusterISOParams) middleware.Responder
	EnableHost(ctx context.Context, params installer.EnableHostParams) middleware.Responder
	GenerateClusterISO(ctx context.Context, params installer.GenerateClusterISOParams) middleware.Responder
	GetCluster(ctx context.Context, params installer.GetClusterParams) middleware.Responder
	GetHost(ctx context.Context, params installer.GetHostParams) middleware.Responder
	GetNextSteps(ctx context.Context, params installer.GetNextStepsParams) middleware.Responder
	InstallCluster(ctx context.Context, params installer.InstallClusterParams) middleware.Responder
	ListClusters(ctx context.Context, params installer.ListClustersParams) middleware.Responder
	ListHosts(ctx context.Context, params installer.ListHostsParams) middleware.Responder
	PostStepReply(ctx context.Context, params installer.PostStepReplyParams) middleware.Responder
	RegisterCluster(ctx context.Context, params installer.RegisterClusterParams) middleware.Responder
	RegisterHost(ctx context.Context, params installer.RegisterHostParams) middleware.Responder
	SetDebugStep(ctx context.Context, params installer.SetDebugStepParams) middleware.Responder
	UpdateCluster(ctx context.Context, params installer.UpdateClusterParams) middleware.Responder
	UpdateHostInstallProgress(ctx context.Context, params installer.UpdateHostInstallProgressParams) middleware.Responder
}

// Config is configuration for Handler
type Config struct {
	InstallerAPI
	Logger func(string, ...interface{})
	// InnerMiddleware is for the handler executors. These do not apply to the swagger.json document.
	// The middleware executes after routing but before authentication, binding and validation
	InnerMiddleware func(http.Handler) http.Handler

	// Authorizer is used to authorize a request after the Auth function was called using the "Auth*" functions
	// and the principal was stored in the context in the "AuthKey" context value.
	Authorizer func(*http.Request) error
}

// Handler returns an http.Handler given the handler configuration
// It mounts all the business logic implementers in the right routing.
func Handler(c Config) (http.Handler, error) {
	h, _, err := HandlerAPI(c)
	return h, err
}

// HandlerAPI returns an http.Handler given the handler configuration
// and the corresponding *AssistedInstall instance.
// It mounts all the business logic implementers in the right routing.
func HandlerAPI(c Config) (http.Handler, *operations.AssistedInstallAPI, error) {
	spec, err := loads.Analyzed(swaggerCopy(SwaggerJSON), "")
	if err != nil {
		return nil, nil, fmt.Errorf("analyze swagger: %v", err)
	}
	api := operations.NewAssistedInstallAPI(spec)
	api.ServeError = errors.ServeError
	api.Logger = c.Logger

	api.JSONConsumer = runtime.JSONConsumer()
	api.BinProducer = runtime.ByteStreamProducer()
	api.JSONProducer = runtime.JSONProducer()
	api.InstallerDeregisterClusterHandler = installer.DeregisterClusterHandlerFunc(func(params installer.DeregisterClusterParams) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		return c.InstallerAPI.DeregisterCluster(ctx, params)
	})
	api.InstallerDeregisterHostHandler = installer.DeregisterHostHandlerFunc(func(params installer.DeregisterHostParams) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		return c.InstallerAPI.DeregisterHost(ctx, params)
	})
	api.InstallerDisableHostHandler = installer.DisableHostHandlerFunc(func(params installer.DisableHostParams) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		return c.InstallerAPI.DisableHost(ctx, params)
	})
	api.InstallerDownloadClusterFilesHandler = installer.DownloadClusterFilesHandlerFunc(func(params installer.DownloadClusterFilesParams) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		return c.InstallerAPI.DownloadClusterFiles(ctx, params)
	})
	api.InstallerDownloadClusterISOHandler = installer.DownloadClusterISOHandlerFunc(func(params installer.DownloadClusterISOParams) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		return c.InstallerAPI.DownloadClusterISO(ctx, params)
	})
	api.InstallerEnableHostHandler = installer.EnableHostHandlerFunc(func(params installer.EnableHostParams) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		return c.InstallerAPI.EnableHost(ctx, params)
	})
	api.InstallerGenerateClusterISOHandler = installer.GenerateClusterISOHandlerFunc(func(params installer.GenerateClusterISOParams) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		return c.InstallerAPI.GenerateClusterISO(ctx, params)
	})
	api.InstallerGetClusterHandler = installer.GetClusterHandlerFunc(func(params installer.GetClusterParams) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		return c.InstallerAPI.GetCluster(ctx, params)
	})
	api.InstallerGetHostHandler = installer.GetHostHandlerFunc(func(params installer.GetHostParams) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		return c.InstallerAPI.GetHost(ctx, params)
	})
	api.InstallerGetNextStepsHandler = installer.GetNextStepsHandlerFunc(func(params installer.GetNextStepsParams) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		return c.InstallerAPI.GetNextSteps(ctx, params)
	})
	api.InstallerInstallClusterHandler = installer.InstallClusterHandlerFunc(func(params installer.InstallClusterParams) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		return c.InstallerAPI.InstallCluster(ctx, params)
	})
	api.InstallerListClustersHandler = installer.ListClustersHandlerFunc(func(params installer.ListClustersParams) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		return c.InstallerAPI.ListClusters(ctx, params)
	})
	api.InstallerListHostsHandler = installer.ListHostsHandlerFunc(func(params installer.ListHostsParams) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		return c.InstallerAPI.ListHosts(ctx, params)
	})
	api.InstallerPostStepReplyHandler = installer.PostStepReplyHandlerFunc(func(params installer.PostStepReplyParams) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		return c.InstallerAPI.PostStepReply(ctx, params)
	})
	api.InstallerRegisterClusterHandler = installer.RegisterClusterHandlerFunc(func(params installer.RegisterClusterParams) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		return c.InstallerAPI.RegisterCluster(ctx, params)
	})
	api.InstallerRegisterHostHandler = installer.RegisterHostHandlerFunc(func(params installer.RegisterHostParams) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		return c.InstallerAPI.RegisterHost(ctx, params)
	})
	api.InstallerSetDebugStepHandler = installer.SetDebugStepHandlerFunc(func(params installer.SetDebugStepParams) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		return c.InstallerAPI.SetDebugStep(ctx, params)
	})
	api.InstallerUpdateClusterHandler = installer.UpdateClusterHandlerFunc(func(params installer.UpdateClusterParams) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		return c.InstallerAPI.UpdateCluster(ctx, params)
	})
	api.InstallerUpdateHostInstallProgressHandler = installer.UpdateHostInstallProgressHandlerFunc(func(params installer.UpdateHostInstallProgressParams) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		return c.InstallerAPI.UpdateHostInstallProgress(ctx, params)
	})
	api.ServerShutdown = func() {}
	return api.Serve(c.InnerMiddleware), api, nil
}

// swaggerCopy copies the swagger json to prevent data races in runtime
func swaggerCopy(orig json.RawMessage) json.RawMessage {
	c := make(json.RawMessage, len(orig))
	copy(c, orig)
	return c
}

// authorizer is a helper function to implement the runtime.Authorizer interface.
type authorizer func(*http.Request) error

func (a authorizer) Authorize(req *http.Request, principal interface{}) error {
	if a == nil {
		return nil
	}
	ctx := storeAuth(req.Context(), principal)
	return a(req.WithContext(ctx))
}

func storeAuth(ctx context.Context, principal interface{}) context.Context {
	return context.WithValue(ctx, AuthKey, principal)
}