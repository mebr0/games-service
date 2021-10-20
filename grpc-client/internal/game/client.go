package game

import (
	"context"
	pb "github.com/mebr0/grpc-client/pkg/api/v1"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"time"
)

type Generator struct {
	period time.Duration
	conn            *grpc.ClientConn
	client          pb.GameServiceClient
	shutdownChannel chan interface{}
}

func NewGenerator(target string, period time.Duration) (*Generator, error) {
	conn, err := grpc.Dial(target, grpc.WithInsecure())

	if err != nil {
		return nil, err
	}

	return &Generator{
		period: period,
		conn:            conn,
		shutdownChannel: make(chan interface{}),
		client:          pb.NewGameServiceClient(conn),
	}, nil
}

func (c *Generator) Start() error {
	go func() {
		for {
			select {
			case <-c.shutdownChannel:
				return
			default:
			}

			response, err := c.client.Send(context.Background(), Random())

			if err != nil {
				log.Fatalf("Error when calling SayHello: %s", err)
			}

			log.Printf("Response from server: %s", response.Message)

			time.Sleep(c.period)
		}
	}()

	return nil
}

func (c *Generator) Stop(ctx context.Context) error {
	if err := c.conn.Close(); err != nil {
		return err
	}

	close(c.shutdownChannel)

	return nil
}
