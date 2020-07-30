package rpc

import (
	"log"
	"net"

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
	for i := 0; i < s.numberOfRequests; i++ {
		go func(sv pb.Streamer_GetRandomDataStreamServer, c *app.Core) {
			url, data, err := c.GetURL(sv.Context())
			if err != nil {
				log.Println(err)
			}
			err = sv.Send(&pb.Response{Url: url, Data: data})
			if err != nil {
				log.Println(err)
			}
		}(srv, s.core)
	}
	return nil
}
