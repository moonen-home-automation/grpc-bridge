package server

import (
	"context"
	"encoding/json"
	"github.com/hashicorp/go-hclog"
	go_hasocket "github.com/moonen-home-automation/go-hasocket"
	"github.com/moonen-home-automation/go-hasocket/pkg/services"
	"github.com/moonen-home-automation/grpc-bridge/internal/proto"
)

type ServicesServer struct {
	log hclog.Logger
	proto.UnimplementedServicesServer
}

func NewServicesServer(l hclog.Logger) *ServicesServer {
	return &ServicesServer{l, proto.UnimplementedServicesServer{}}
}

func (s *ServicesServer) CallService(ctx context.Context, sc *proto.ServiceCall) (*proto.ServiceResponse, error) {
	hass, err := go_hasocket.GetApp()
	if err != nil {
		s.log.Warn("Error getting hass app", err)
		return nil, err
	}

	data := make(map[string]interface{})
	_ = json.Unmarshal([]byte(sc.JsonData), &data)

	sr := services.NewServiceRequest()
	sr.Domain = sc.Domain
	sr.Service = sc.Service
	sr.Target = services.ServiceTarget{
		AreaId:   sc.GetAreaId(),
		DeviceId: sc.GetDeviceId(),
		EntityId: sc.GetEntityId(),
		LabelId:  sc.GetLabelId(),
	}
	sr.ServiceData = data
	sr.ReturnResponse = sc.Returns

	resp, err := hass.CallService(sr)
	if err != nil {
		s.log.Warn("Error calling service", err)
		return nil, err
	}

	jsonResp := []byte("")

	if sc.Returns {
		jsonResp, err = json.Marshal(resp.Result.Response)
		if err != nil {
			s.log.Warn("Error unmarshalling service response", err)
			return nil, err
		}
	}

	return &proto.ServiceResponse{
		JsonData: string(jsonResp),
	}, nil
}
