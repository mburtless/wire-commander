// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/mburtless/wire-commander/gen/restapi/operations"
)

//go:generate swagger generate server --target ../../gen --name WireCommander --spec ../../swagger/swagger.yml --exclude-main

func configureFlags(api *operations.WireCommanderAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.WireCommanderAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.CreatePeerHandler == nil {
		api.CreatePeerHandler = operations.CreatePeerHandlerFunc(func(params operations.CreatePeerParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.CreatePeer has not yet been implemented")
		})
	}
	if api.DeletePeerHandler == nil {
		api.DeletePeerHandler = operations.DeletePeerHandlerFunc(func(params operations.DeletePeerParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.DeletePeer has not yet been implemented")
		})
	}
	if api.GetConfigHandler == nil {
		api.GetConfigHandler = operations.GetConfigHandlerFunc(func(params operations.GetConfigParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetConfig has not yet been implemented")
		})
	}
	if api.GetPeerHandler == nil {
		api.GetPeerHandler = operations.GetPeerHandlerFunc(func(params operations.GetPeerParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetPeer has not yet been implemented")
		})
	}
	if api.GetPeersHandler == nil {
		api.GetPeersHandler = operations.GetPeersHandlerFunc(func(params operations.GetPeersParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetPeers has not yet been implemented")
		})
	}
	if api.UpdatePeerHandler == nil {
		api.UpdatePeerHandler = operations.UpdatePeerHandlerFunc(func(params operations.UpdatePeerParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.UpdatePeer has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
