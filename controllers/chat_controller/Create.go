package chat_controller

import (
	"context"
	"crypto/rand"
	"encoding/binary"
	"log"

	"github.com/justbrownbear/microservices_course_chat/pkg/chat_v1"
)

// ***************************************************************************************************
// ***************************************************************************************************
func (s *controller) Create(_ context.Context, req *chat_v1.CreateRequest) (*chat_v1.CreateResponse, error) {

	log.Printf("Create request fired: %v", req.String())

	result := &chat_v1.CreateResponse{
		Id: generateRandomID(),
	}

	return result, nil
}

// ***************************************************************************************************
// ***************************************************************************************************
func generateRandomID() int64 {
	var result int64

	err := binary.Read(rand.Reader, binary.BigEndian, &result)
	if err != nil {
		log.Printf("Failed to generate random id: %v", err)
		return 0
	}

	if result < 0 {
		result = -result
	}

	result = result % 100500

	return result
}
