package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"

	context "golang.org/x/net/context"

	"github.com/kelseyhightower/envconfig"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/pressly/warpdrive/proto"
)

type grpcConfig struct {
	certPool    *x509.CertPool
	certificate tls.Certificate
}

func (g *grpcConfig) createServer() (*grpc.Server, error) {
	tlsConfig := &tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{g.certificate},
		ClientCAs:    g.certPool,
	}

	serverOption := grpc.Creds(credentials.NewTLS(tlsConfig))
	server := grpc.NewServer(serverOption)

	return server, nil
}

func newGrpcConfig(ca, crt, key string) (*grpcConfig, error) {
	certificate, err := tls.LoadX509KeyPair(
		crt,
		key,
	)

	certPool := x509.NewCertPool()
	bs, err := ioutil.ReadFile(ca)
	if err != nil {
		return nil, fmt.Errorf("failed to read client ca cert: %s", err)
	}

	ok := certPool.AppendCertsFromPEM(bs)
	if !ok {
		return nil, fmt.Errorf("failed to append client certs")
	}

	return &grpcConfig{certPool, certificate}, nil
}

type commandServer struct {
}

func (c *commandServer) CreateApp(context.Context, *pb.App) (*pb.App, error) {
	return nil, nil
}
func (c *commandServer) GetAllApps(*pb.Empty, pb.Command_GetAllAppsServer) error {
	return nil
}
func (c *commandServer) RemoveApp(context.Context, *pb.App) (*pb.Empty, error) {
	return nil, nil
}
func (c *commandServer) CreateRelease(context.Context, *pb.Release) (*pb.Release, error) {
	return nil, nil
}
func (c *commandServer) GetRelease(context.Context, *pb.Release) (*pb.Release, error) {
	return nil, nil
}
func (c *commandServer) UpdateRelease(context.Context, *pb.Release) (*pb.Release, error) {
	return nil, nil
}

type queryServer struct {
}

func (q *queryServer) GetUpgrade(context.Context, *pb.Upgrade) (*pb.Release, error) {
	return nil, nil
}

func main() {
	commandEnv := &struct {
		CA  string `require:"true"`
		Crt string `require:"true"`
		Key string `require:"true"`
	}{}

	err := envconfig.Process("command", commandEnv)
	if err != nil {
		log.Fatal(err.Error())
	}

	queryEnv := &struct {
		CA  string `require:"true"`
		Crt string `require:"true"`
		Key string `require:"true"`
	}{}

	err = envconfig.Process("query", queryEnv)
	if err != nil {
		log.Fatal(err.Error())
	}

	grpcCommandConfig, err := newGrpcConfig(commandEnv.CA, commandEnv.Crt, commandEnv.Key)
	if err != nil {
		log.Fatal(err.Error())
	}

	grpcQueryConfig, err := newGrpcConfig(queryEnv.CA, queryEnv.Crt, queryEnv.Key)
	if err != nil {
		log.Fatal(err.Error())
	}

	grpcCommandServer, err := grpcCommandConfig.createServer()
	if err != nil {
		log.Fatal(err.Error())
	}

	grpcQueryServer, err := grpcQueryConfig.createServer()
	if err != nil {
		log.Fatal(err.Error())
	}

	go func() {
		pb.RegisterCommandServer(grpcCommandServer, &commandServer{})
	}()

	go func() {
		pb.RegisterQueryServer(grpcQueryServer, &queryServer{})
	}()
}
