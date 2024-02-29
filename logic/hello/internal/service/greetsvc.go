package service

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"github.com/fengjx/luchen"
)

var GreetService *greetService

type greetService struct {
}

func init() {
	GreetService = newGreetService()
}

func newGreetService() *greetService {
	return &greetService{}
}

func (svc *greetService) SayHi(ctx context.Context, name string) (string, error) {
	logger := luchen.Logger(ctx)
	logger.Info("say hi", zap.Any("name", name))
	return fmt.Sprintf("Hi: %s", name), nil
}
