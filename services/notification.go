package services

import (
	"context"
	"firebase.google.com/go/v4/messaging"
	"github.com/qcodelabsllc/qreeket/notification/config"
	pb "github.com/qcodelabsllc/qreeket/notification/generated"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"log"
)

type QreeketNotificationServer struct {
	messagingClient *messaging.Client
	pb.UnimplementedNotificationServiceServer
}

func NewQreeketNotificationServer(messagingClient *messaging.Client) *QreeketNotificationServer {
	return &QreeketNotificationServer{messagingClient: messagingClient}
}

func (q *QreeketNotificationServer) SendNotification(ctx context.Context, req *pb.SendNotificationRequest) (*wrapperspb.StringValue, error) {
	if len(req.GetTitle()) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "notification title is required")
	}

	if len(req.GetBody()) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "notification body is required")
	}

	if len(req.GetTopic()) > 0 && len(req.GetToken()) > 0 {
		return nil, status.Errorf(codes.InvalidArgument, "notification topic and token cannot be used at the same time")
	}

	// create payload
	payload := &messaging.Message{
		Notification: &messaging.Notification{
			Title: req.GetTitle(),
			Body:  req.GetBody(),
		},
	}

	if len(req.GetToken()) > 0 {
		payload.Token = req.GetToken()
	}

	if len(req.GetTopic()) > 0 {
		payload.Topic = req.GetTopic()
	}

	// test the messaging client
	sendResult, err := config.FirebaseMessaging.Send(ctx, payload)
	if err != nil {
		log.Fatalf("unable to send message: %+v\n", err)
	}

	log.Printf("message sent: %+v\n", sendResult)

	return &wrapperspb.StringValue{Value: sendResult}, nil
}

func (q *QreeketNotificationServer) RegisterDevice(ctx context.Context, req *pb.RegisterDeviceRequest) (*emptypb.Empty, error) {
	if len(req.GetToken()) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "messaging token is required")
	}

	if len(req.GetTopic()) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "topic is required")
	}

	if len(req.GetToken()) > 0 && len(req.GetTopic()) > 0 {
		return nil, status.Errorf(codes.InvalidArgument, "token and topic cannot be used at the same time")
	}

	if len(req.GetTopic()) > 0 {
		// subscribe the device to the topic
		response, err := config.FirebaseMessaging.SubscribeToTopic(ctx, []string{req.GetToken()}, req.GetTopic())
		if err != nil {
			log.Fatalf("unable to subscribe to topic: %+v\n", err)
		}

		log.Printf("subscribed to topic: %+v\n", response.SuccessCount)
	}

	if len(req.GetToken()) > 0 {
		// subscribe the device to the topic
		response, err := config.FirebaseMessaging.Send(ctx, &messaging.Message{
			Token: req.GetToken(),
			Notification: &messaging.Notification{
				Title: "Welcome to Qreeket",
				Body:  "You have successfully registered your device for push notifications 🥳",
			},
		})
		if err != nil {
			log.Fatalf("unable to subscribe to topic: %+v\n", err)
		}

		log.Printf("registered device for notifications: %+v\n", response)
	}

	return &emptypb.Empty{}, nil
}

func (q *QreeketNotificationServer) UnregisterDevice(ctx context.Context, req *pb.RegisterDeviceRequest) (*emptypb.Empty, error) {
	if len(req.GetToken()) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "messaging token is required")
	}

	if len(req.GetTopic()) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "topic is required")
	}

	if len(req.GetToken()) > 0 && len(req.GetTopic()) > 0 {
		return nil, status.Errorf(codes.InvalidArgument, "token and topic cannot be used at the same time")
	}

	if len(req.GetTopic()) > 0 {
		// unsubscribe the device to the topic
		response, err := config.FirebaseMessaging.UnsubscribeFromTopic(ctx, []string{req.GetToken()}, req.GetTopic())
		if err != nil {
			log.Fatalf("unable to unsubscribe to topic: %+v\n", err)
		}

		log.Printf("unsubscribed from topic: %+v\n", response.SuccessCount)
	}

	// ignore the token if topic is provided
	return &emptypb.Empty{}, nil
}
