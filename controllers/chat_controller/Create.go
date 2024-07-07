package chat_controller

import (
	"context"
	"log"
	"math/rand"

	"github.com/justbrownbear/microservices_course_chat/pkg/chat_v1"
)



func (s *controller) Create(ctx context.Context, req *chat_v1.CreateRequest) (*chat_v1.CreateResponse, error) {

	log.Printf("Create request fired: %v", req.String())

	result := &chat_v1.CreateResponse{
		Id: rand.Int63n( 100500 ),
	}

	return result, nil
}
