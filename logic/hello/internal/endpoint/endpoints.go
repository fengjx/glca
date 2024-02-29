package endpoint

import (
	"context"

	"github.com/fengjx/luchen"
	"github.com/go-kit/kit/endpoint"

	"github.com/fengjx/glca/logic/hello/internal/service"
	"github.com/fengjx/glca/pb"
)

func Init(_ context.Context, httpServer *luchen.HTTPServer) {
	httpServer.Handler(newGreeterHandler())
}

func MakeSayHelloEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		logger := luchen.Logger(ctx)
		logger.Info("greeter say hello")
		helloReq := request.(*pb.HelloReq)
		msg, err := service.GreetService.SayHi(ctx, helloReq.Name)
		if err != nil {
			return nil, err
		}
		return &pb.HelloResp{Message: msg}, nil
	}
}
