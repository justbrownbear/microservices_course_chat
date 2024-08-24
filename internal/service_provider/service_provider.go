package service_provider

import (
	"github.com/jackc/pgx/v5"

	chat_repository "github.com/justbrownbear/microservices_course_chat/internal/repository/chat"
	chat_service "github.com/justbrownbear/microservices_course_chat/internal/service/chat"
)

// ServiceProvider - Интерфейс сервис-провайдера
type ServiceProvider interface {
	GetChatService() chat_service.ChatService
}

type serviceProvider struct {
	dbConnection  *pgx.Conn
	dbTransaction *pgx.Tx

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


func (serviceProviderInstance *serviceProvider) getChatRepository() chat_repository.ChatRepository {
	if serviceProviderInstance.chatRepository == nil {
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
