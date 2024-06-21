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

	_, err := go_hasocket.NewApp("ws://192.168.4.100:8123/api/websocket", os.Getenv("HA_TOKEN"))
	if err != nil {
		log.Error("Hass client error", "error", err)
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
