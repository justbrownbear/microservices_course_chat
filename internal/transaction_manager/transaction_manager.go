package transaction_manager

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"

	"github.com/justbrownbear/microservices_course_chat/internal/service_provider"
)

// Handler - функция, которая выполняется в транзакции или коннекшне,
// это будет лежать в оригинале в другом package
type Handler func(ctx context.Context, serviceProvider service_provider.ServiceProvider) error

// TxManager defines an interface for managing database transactions.
// It provides methods to execute handlers within the context of a connection,
// either with or without a transaction.
type TxManager interface {
	// Выполнение handler в контексте подключения без транзакции
	// WithConnection(ctx context.Context, handler Handler) error
	// Выполнение handler в контексте подключения с транзакцией
	WithTransaction(ctx context.Context, handler Handler) error
}

// В перспективе тут может быть не только pg, но и другие подключения.
// Например, redis
type resources struct {
	dbPool *pgxpool.Pool
}

// InitTransactionManager initializes and returns a new transaction manager instance.
// It takes a pgxpool.Pool as an argument, which is used to manage database connections.
//
// Parameters:
//   - dbPool: A pointer to a pgxpool.Pool that provides the database connection pool.
//
// Returns:
//   - TxManager: An instance of TxManager that can be used to manage transactions.
func InitTransactionManager(dbPool *pgxpool.Pool) TxManager {
	return &resources{
		dbPool: dbPool,
	}
}

// WithTransaction executes a handler function within a database transaction context.
// It begins a new transaction, passes a service provider initialized with the transaction
// to the handler, and commits or rolls back the transaction based on the outcome of the handler.
//
// If the handler returns an error, the transaction is rolled back. If the handler executes
// successfully, the transaction is committed. In case of a panic within the handler, the
// transaction is rolled back and the panic is recovered.
//
// Parameters:
//
//	ctx - The context for the transaction.
//	handler - The function to execute within the transaction context.
//
// Returns:
//
//	An error if the transaction could not be started, committed, or if the handler returns an error.
func (instance *resources) WithTransaction(
	ctx context.Context,
	handler Handler,
) error {
	// Инициализируем соединение
	transaction, err := instance.dbPool.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return err
	}

	// Настраиваем функцию отсрочки для отката или коммита транзакции.
	defer func() {
		// Восстанавливаемся после паники
		recoverResult := recover()
		if recoverResult != nil {
			err = errors.Errorf("panic recovered: %v", recoverResult)
		}

		// Откатываем транзакцию, если произошла ошибка
		if err != nil {
			errRollback := transaction.Rollback(ctx)
			if errRollback != nil {
				err = errors.Wrapf(err, "errRollback: %v", errRollback)
			}

			return
		}
	}()

	// Инициализируем сервис-провайдер транзакцией
	serviceProvider := service_provider.NewWithTransaction(&transaction)
	// А можем и с коннекшеном
	// serviceProvider := getServiceProviderWithConnection(&connection)

	// Выполняем бизнес-логику
	err = handler(ctx, serviceProvider)
	if err != nil {
		err = errors.Wrap(err, "failed executing code inside transaction")
	}

	// если ошибок не было, коммитим транзакцию
	if err == nil {
		err = transaction.Commit(ctx)
		if err != nil {
			err = errors.Wrap(err, "tx commit failed")
		}
	}

	return err
}
