package main

import (
	"github.com/Sirupsen/logrus"
	micro "github.com/micro/go-micro"
	"github.com/nii236/nii-forex/tickSubscriber/proto"
	server "github.com/nii236/nii-forex/tickSubscriber/server"
	"golang.org/x/net/context"
)

type Tick struct{}

func (t *Tick) TickHandler(ctx context.Context, req *tick.TickRequest, rsp *tick.TickResponse) error {
	return nil
}

var (
	natsURL = "192.168.99.100:32773"
	log     *logrus.Logger
)

func getOptions(o *micro.Options) {
	o.Server = server.NewNatsServer()
}

func main() {
	log = logrus.New()
	service := micro.NewService(getOptions)
	service.Run()
}