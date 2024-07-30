package service_provider

import (
	"github.com/jackc/pgx/v5"

	chat_repository "github.com/justbrownbear/microservices_course_chat/internal/repository/chat"
	user_repository "github.com/justbrownbear/microservices_course_chat/internal/repository/user"
	chat_service "github.com/justbrownbear/microservices_course_chat/internal/service/chat"
	user_service "github.com/justbrownbear/microservices_course_chat/internal/service/user"
)

// ServiceProvider - Интерфейс сервис-провайдера
type ServiceProvider interface {
	GetUserService() user_service.UserService
	GetChatService() chat_service.ChatService
}

type serviceProvider struct {
	dbConnection  *pgx.Conn
	dbTransaction *pgx.Tx

	userRepository *user_repository.Queries
	userService    user_service.UserService

	chatRepository *chat_repository.Queries
	chatService    chat_service.ChatService
}

// NewWithConnection создает новый экземпляр сервис-провайдера с соединением
func NewWithConnection(dbConnection *pgx.Conn) ServiceProvider {
	return &serviceProvider{
		dbConnection: dbConnection,
	}
}

// NewWithTransaction создает новый экземпляр сервис-провайдера с транзакцией
func NewWithTransaction(dbTransaction *pgx.Tx) ServiceProvider {
	return &serviceProvider{
		dbTransaction: dbTransaction,
	}
}

func (serviceProviderInstance *serviceProvider) getUserRepository() user_repository.UserRepository {
	if serviceProviderInstance.userRepository == nil {
		serviceProviderInstance.userRepository = user_repository.New(serviceProviderInstance.dbConnection)

		if serviceProviderInstance.dbTransaction != nil {
			serviceProviderInstance.userRepository = serviceProviderInstance.userRepository.WithTx(*serviceProviderInstance.dbTransaction)
		}
	}

	return serviceProviderInstance.userRepository
}

func (serviceProviderInstance *serviceProvider) GetUserService() user_service.UserService {
	if serviceProviderInstance.userService == nil {
		serviceProviderInstance.userService = user_service.New(serviceProviderInstance.getUserRepository())
	}

	return serviceProviderInstance.userService
}

func (serviceProviderInstance *serviceProvider) getChatRepository() chat_repository.ChatRepository {
	if serviceProviderInstance.userRepository == nil {
		serviceProviderInstance.chatRepository = chat_repository.New(serviceProviderInstance.dbConnection)

		if serviceProviderInstance.dbTransaction != nil {
			serviceProviderInstance.chatRepository = serviceProviderInstance.chatRepository.WithTx(*serviceProviderInstance.dbTransaction)
		}
	}

	return serviceProviderInstance.chatRepository
}

func (serviceProviderInstance *serviceProvider) GetChatService() chat_service.ChatService {
	if serviceProviderInstance.chatService == nil {
		serviceProviderInstance.chatService = chat_service.New(serviceProviderInstance.getChatRepository())
	}

	return serviceProviderInstance.chatService
}
