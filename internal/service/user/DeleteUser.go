package user_service

import "context"

// Удаление пользователя по id
func (userServiceInstance *userService) DeleteUser(ctx context.Context, userID uint64) error {
	err := userServiceInstance.repository.DeleteUser(ctx, int64(userID))
	if err != nil {
		return err
	}

	return nil
}
