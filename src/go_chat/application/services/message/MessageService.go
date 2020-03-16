package message

import "go_chat/application/services/push"

type MessageService struct {
	push.PushEventGenerator
}
