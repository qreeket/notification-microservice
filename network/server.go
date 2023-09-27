package network

import (
	"github.com/qcodelabsllc/qreeket/notification/config"
	pb "github.com/qcodelabsllc/qreeket/notification/generated"
	"github.com/qcodelabsllc/qreeket/notification/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"strconv"
)

func InitServer() {
	// create a new grpc server
	s := grpc.NewServer()

	// register the grpc server
	pb.RegisterNotificationServiceServer(s, services.NewQreeketNotificationServer(config.FirebaseMessaging))

	// register the grpc server for reflection
	reflection.Register(s)

	// get the port number from .env file
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	// listen on the port
	if lis, err := net.Listen("tcp", ":"+strconv.Itoa(port)); err == nil {
		log.Printf("groups server started on %v\n", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("unable to start grpc server: %+v\n", err)
		}
	}
}
