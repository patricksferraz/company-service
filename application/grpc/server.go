package grpc

import (
	"fmt"
	"log"
	"net"

	_service "github.com/patricksferraz/company-service/domain/service"
	"github.com/patricksferraz/company-service/infrastructure/db"
	"github.com/patricksferraz/company-service/infrastructure/external"
	"github.com/patricksferraz/company-service/infrastructure/repository"
	"github.com/patricksferraz/company-service/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartGrpcServer(database *db.Postgres, authConn *grpc.ClientConn, kafkaProducer *external.KafkaProducer, port int) {

	authClient := external.NewAuthClient(authConn)
	interceptor := NewAuthInterceptor(authClient)
	repository := repository.NewRepository(database, kafkaProducer)
	service := _service.NewService(repository)
	grpcService := NewGrpcService(service)

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	)
	reflection.Register(grpcServer)
	pb.RegisterCompanyServiceServer(grpcServer, grpcService)

	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start grpc server", err)
	}

	log.Printf("gRPC server has been started on port %d", port)

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start grpc server", err)
	}
}
