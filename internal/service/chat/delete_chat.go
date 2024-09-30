package chat_service

import "context"

func (chatServiceInstance *chatService) DeleteChat(ctx context.Context, chatID uint64) error {
	err := chatServiceInstance.repository.DeleteChat(ctx, int64(chatID))
	if err != nil {
		return err
	}

	return err
}
