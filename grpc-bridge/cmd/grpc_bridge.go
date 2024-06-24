package main

import (
	"github.com/hashicorp/go-hclog"
	go_hasocket "github.com/moonen-home-automation/go-hasocket"
	"github.com/moonen-home-automation/grpc-bridge/internal/proto"
	"github.com/moonen-home-automation/grpc-bridge/internal/server"
	"google.golang.org/grpc"
	"net"
	"os"
)

func main() {
	log := hclog.Default()

	_, err := go_hasocket.NewApp("ws://supervisor/core/websocket", os.Getenv("SUPERVISOR_TOKEN"))
	if err != nil {
		log.Error("Error getting hass client", "error", err)
		os.Exit(1)
	}
	log.Info("Initialized new hass app")

	gs := grpc.NewServer()

	es := server.NewEventServer(log)
	proto.RegisterEventsServer(gs, es)

	ss := server.NewServicesServer(log)
	proto.RegisterServicesServer(gs, ss)

	l, err := net.Listen("tcp", ":9091")
	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}
	log.Info("Started listening", "port", "9091")

	err = gs.Serve(l)
	if err != nil {
		log.Error("Error serving", err)
		os.Exit(1)
	}
}
