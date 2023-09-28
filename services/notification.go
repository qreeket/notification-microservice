package services

import (
	"context"
	"firebase.google.com/go/v4/messaging"
	"github.com/qcodelabsllc/qreeket/notification/config"
	pb "github.com/qcodelabsllc/qreeket/notification/generated"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type QreeketNotificationServer struct {
	messagingClient *messaging.Client
	pb.UnimplementedNotificationServiceServer
}

func NewQreeketNotificationServer(messagingClient *messaging.Client) *QreeketNotificationServer {
	return &QreeketNotificationServer{messagingClient: messagingClient}
}

func (q *QreeketNotificationServer) SendNotification(ctx context.Context, req *pb.SendNotificationRequest) (*pb.StringValue, error) {
	if len(req.GetTitle()) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "notification title is required")
	}
	
	if len(req.GetBody()) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "notification body is required")
	}
	
	if len(req.GetTopic()) > 0 && len(req.GetToken()) > 0 {
		return nil, status.Errorf(codes.InvalidArgument, "notification topic and token cannot be used at the same time")
	}
	
	// message data
	msgData := map[string]string{
		"click_action": "FLUTTER_NOTIFICATION_CLICK",
		"sound":        "default",
		"channel_id":   getChannelIdFromNotificationType(req.GetChannelType()),
	}
	
	for k, v := range req.GetData() {
		msgData[k] = v
	}
	
	// create payload
	payload := &messaging.Message{
		Notification: &messaging.Notification{
			Title: req.GetTitle(),
			Body:  req.GetBody(),
		},
		Data: msgData,
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
		return nil, status.Errorf(codes.Internal, "unable to send notification: %+v", err)
	}
	
	log.Printf("message sent: %+v\n", sendResult)
	
	return &pb.StringValue{Value: sendResult}, nil
}

func (q *QreeketNotificationServer) RegisterDevice(ctx context.Context, req *pb.RegisterDeviceRequest) (*pb.Empty, error) {
	if len(req.GetToken()) > 0 && len(req.GetTopic()) > 0 {
		return nil, status.Errorf(codes.InvalidArgument, "token and topic cannot be used at the same time")
	}
	
	empty := &pb.Empty{}
	if len(req.GetTopic()) > 0 {
		// subscribe the device to the topic
		response, err := config.FirebaseMessaging.SubscribeToTopic(ctx, []string{req.GetToken()}, req.GetTopic())
		if err != nil {
			return nil, status.Errorf(codes.Internal, "unable to subscribe to topic: %+v", err)
		}
		
		log.Printf("subscribed to topic: %+v\n", response.SuccessCount)
		return empty, nil
	}
	
	if len(req.GetToken()) > 0 {
		// subscribe the device to the topic
		response, err := config.FirebaseMessaging.Send(ctx, &messaging.Message{
			Token: req.GetToken(),
			Notification: &messaging.Notification{
				Title: "Welcome to Qreeket",
				Body:  "You have successfully registered your device for push notifications 🥳",
			},
			Data: map[string]string{
				"click_action": "FLUTTER_NOTIFICATION_CLICK",
				"sound":        "default",
				"channel_id":   getChannelIdFromNotificationType(pb.NotificationChannelType_BROADCAST),
			},
		})
		if err != nil {
			return nil, status.Errorf(codes.Internal, "unable to send notification: %+v", err)
		}
		
		log.Printf("registered device for notifications: %+v\n", response)
		return empty, nil
	}
	
	return empty, status.Errorf(codes.InvalidArgument, "token or topic is required")
}

func (q *QreeketNotificationServer) UnregisterDevice(ctx context.Context, req *pb.RegisterDeviceRequest) (*pb.Empty, error) {
	if len(req.GetToken()) > 0 && len(req.GetTopic()) > 0 {
		return nil, status.Errorf(codes.InvalidArgument, "token and topic cannot be used at the same time")
	}
	
	empty := &pb.Empty{}
	if len(req.GetTopic()) > 0 {
		// unsubscribe the device to the topic
		response, err := config.FirebaseMessaging.UnsubscribeFromTopic(ctx, []string{req.GetToken()}, req.GetTopic())
		if err != nil {
			return nil, status.Errorf(codes.Internal, "unable to unsubscribe from topic: %+v", err)
		}
		
		log.Printf("unsubscribed from topic: %+v\n", response.SuccessCount)
		return empty, nil
	}
	
	// ignore the token if topic is provided
	return empty, status.Errorf(codes.InvalidArgument, "token or topic is required")
}

func getChannelIdFromNotificationType(channelType pb.NotificationChannelType) string {
	var channelId string
	switch channelType {
	case pb.NotificationChannelType_E2E_PERSONAL_CHAT:
		channelId = "e2e_personal_chat"
	case pb.NotificationChannelType_E2E_GROUP_CHAT:
		channelId = "e2e_group_chat"
	case pb.NotificationChannelType_CHANNEL_INVITATION:
		channelId = "channel_invitation"
	case pb.NotificationChannelType_TOPIC:
		channelId = "special_topic"
	case pb.NotificationChannelType_SUBSCRIPTION:
		channelId = "subscription"
	case pb.NotificationChannelType_ACCOUNT:
		channelId = "account"
	case pb.NotificationChannelType_UPDATE,
		pb.NotificationChannelType_BROADCAST,
		pb.NotificationChannelType_UNKNOWN:
		channelId = "general"
	}
	
	return channelId
}
