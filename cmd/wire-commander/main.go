package main

import (
	"flag"
	"log"

	"github.com/go-openapi/loads"

	"github.com/mburtless/wire-commander/gen/restapi"
	"github.com/mburtless/wire-commander/gen/restapi/operations"
	"github.com/mburtless/wire-commander/api/handlers"
)

var portFlag = flag.Int("port", 3000, "Port to run this service on")

func main() {
	// load embedded swagger file
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	// create new service API
	api := operations.NewWireCommanderAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer func() {
		_ = server.Shutdown()
	}()

	// parse flags
	flag.Parse()
	// set the port this service will be run on
	server.Port = *portFlag

    // handlers
    api.GetPeerHandler = handlers.NewGetPeer()

	// serve API
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
