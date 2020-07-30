package rpc

import (
	"log"
	"net"
	"sync"

	"github.com/abdukahhor/streamer/app"
	"github.com/abdukahhor/streamer/handlers/pb"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

var defaultServer *grpc.Server

//rpc server
type server struct {
	core             *app.Core
	numberOfRequests int
}

//Register RPC Increment Server and grpc Server
func Register(c *app.Core, numReq int) {
	defaultServer = grpc.NewServer()
	pb.RegisterStreamerServer(defaultServer, &server{core: c, numberOfRequests: numReq})
}

//Run serve grpc Server
func Run(addr string) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer l.Close()
	err = defaultServer.Serve(l)
	if err != nil {
		return err
	}
	return nil
}

//Close stop grpc Server
func Close() {
	defaultServer.GracefulStop()
}

func (s *server) GetRandomDataStream(in *empty.Empty, srv pb.Streamer_GetRandomDataStreamServer) error {
	var wg sync.WaitGroup
	for i := 0; i < s.numberOfRequests; i++ {
		wg.Add(1)
		go func(sv pb.Streamer_GetRandomDataStreamServer, c *app.Core) {
			defer wg.Done()
			url, data, err := c.GetURL(sv.Context())
			if err != nil {
				log.Println("c.GetURL", err)
			}
			err = sv.Send(&pb.Response{Url: url, Data: data})
			if err != nil {
				log.Println("sv.Send", err)
			}
		}(srv, s.core)
	}
	wg.Wait()
	return nil
}
