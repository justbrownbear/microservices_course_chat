package grpc_api

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	chat_service "github.com/justbrownbear/microservices_course_chat/internal/service/chat"
	"github.com/justbrownbear/microservices_course_chat/internal/service_provider"
)

func (grpcApiInstance *grpcAPI) CreateChat(ctx context.Context, in *chat_service.CreateChatRequest) (uint64, error) {
	var chatID uint64;

	// Запускаем код с транзакцией, и указываем это явно
	err := withTransaction( ctx, grpcApiInstance.dbPool,
		func ( ctx context.Context, serviceProvider service_provider.ServiceProvider ) error {
			// В этом месте нам пришел сервис-провайдер, который уже имеет connection внутри себя
			// Нам осталось только получить нужные сервисы, и...
			chatService := serviceProvider.GetChatService()

			// ...Передать их функции, которая на входе принимает только используемые сервисы и in
			var err error
			chatID, err = createChat( ctx, chatService, in )
			if err != nil {
				return err
			}

			return nil
		} )
	if err != nil {
		return 0, err
	}

	return chatID, nil
}


// А это уже простая и красивая функция, реализующая бизнес-логику
func createChat(ctx context.Context, chatService chat_service.ChatService, in *chat_service.CreateChatRequest) (uint64, error) {
	chatID, err := chatService.CreateChat(ctx, in)
	if err != nil {
		return 0, err
	}

	return chatID, nil
}



// Handler - функция, которая выполняется в транзакции, это будет лежать в оригинале в другом package
type Handler func(ctx context.Context, serviceProvider service_provider.ServiceProvider) error

// Разумеется, это будет лежать в оригинале в другом package
func withTransaction( ctx context.Context, dbPool *pgxpool.Pool, fn Handler ) error {
	// Инициализируем соединение
	transaction, err := dbPool.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return err
	}

	// Настраиваем функцию отсрочки для отката или коммита транзакции.
	defer func() {
		// восстанавливаемся после паники
		if r := recover(); r != nil {
			err = errors.Errorf("panic recovered: %v", r)
		}

		// откатываем транзакцию, если произошла ошибка
		if err != nil {
			if errRollback := tx.Rollback(ctx); errRollback != nil {
				err = errors.Wrapf(err, "errRollback: %v", errRollback)
			}

			return
		}

		// если ошибок не было, коммитим транзакцию
		if nil == err {
			err = tx.Commit(ctx)
			if err != nil {
				err = errors.Wrap(err, "tx commit failed")
			}
		}
	}()

	// Инициализируем сервис-провайдер транзакцией
	serviceProvider := getServiceProviderWithTransaction(&transaction)
	// А можем и с коннекшеном
	// serviceProvider := getServiceProviderWithConnection(&connection)

	// Выполняем бизнес-логику
	err = fn( ctx, serviceProvider )
	if err != nil {
		err = errors.Wrap(err, "failed executing code inside transaction")
	}

	return err
}
