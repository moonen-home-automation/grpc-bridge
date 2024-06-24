package server

import (
	"github.com/hashicorp/go-hclog"
	go_hasocket "github.com/moonen-home-automation/go-hasocket"
	"github.com/moonen-home-automation/go-hasocket/pkg/events"
	"github.com/moonen-home-automation/grpc-bridge/internal/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

type EventServer struct {
	log hclog.Logger
	proto.UnimplementedEventsServer
}

func NewEventServer(l hclog.Logger) *EventServer {
	return &EventServer{l, proto.UnimplementedEventsServer{}}
}

func (e *EventServer) Subscribe(rr *proto.EventSubscriberRequest, src proto.Events_SubscribeServer) error {
	hass, err := go_hasocket.GetApp()
	if err != nil {
		e.log.Warn("Error getting hass app", err)
		return err
	}

	eventChannel := make(chan events.EventData, 10)
	eventListener, err := hass.RegisterListener(rr.GetType())
	if err != nil {
		e.log.Warn("Error registering event listener", err)
		return err
	}
	e.log.Info("Registered event listener", "event_type", rr.GetType())

	go eventListener.Listen(eventChannel)
	defer func() {
		err := eventListener.Close()
		if err != nil {
			e.log.Warn("Error closing event listener", err)
		}
		e.log.Info("Closed event listener", "event_type", rr.GetType())
	}()

	for {
		select {
		case <-src.Context().Done():
			e.log.Info("Client disconnected or request cancelled")
			return src.Context().Err()
		case msg, ok := <-eventChannel:
			if !ok {
				e.log.Warn("Received event not OK, continuing", msg)
				continue
			}

			err := src.Send(&proto.Event{
				Type: msg.Type,
				Data: &anypb.Any{Value: msg.RawEventJSON},
			})
			if err != nil {
				e.log.Warn("Error sending event", err)
				return err
			}
		}
	}
}
