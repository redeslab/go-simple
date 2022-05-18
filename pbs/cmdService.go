package pbs

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type cmdService struct{}

var param struct {
	password     string
	minerIP      string
	user         string
	priKey       string
	path         string
	confOp       int8
	one          bool
	all          bool
	id           string
	contractAddr string
}

var CMDServicePort = "12776"

func StartCmdService(port string) {
	CMDServicePort = port
	address := net.JoinHostPort("127.0.0.1", CMDServicePort)
	l, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	cmdServer := grpc.NewServer()

	RegisterCmdServiceServer(cmdServer, &cmdService{})

	reflection.Register(cmdServer)
	if err := cmdServer.Serve(l); err != nil {
		panic(err)
	}
}

func DialToCmdService() CmdServiceClient {
	var address = "127.0.0.1:" + CMDServicePort
	conn, err := grpc.Dial(address)
	if err != nil {
		panic(err)
	}

	client := NewCmdServiceClient(conn)

	return client
}

func (s *cmdService) SetLogLevel(ctx context.Context, req *LogLevel) (result *CommonResponse, err error) {
	return nil, nil
}
