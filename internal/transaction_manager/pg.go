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


type TxManager interface {
	WithTransaction(ctx context.Context, fn Handler) error
}


// В перспективе тут может быть не только pg, но и другие подключения.
// Например, redis
type resources struct {
	dbPool *pgxpool.Pool
}


func InitTransactionManager(dbPool *pgxpool.Pool) TxManager {
	return &resources{
		dbPool: dbPool,
	}
}



// Разумеется, это будет лежать в оригинале в другом package
func (instance *resources) WithTransaction(
	ctx context.Context,
	fn Handler,
) error {
	// Инициализируем соединение
	transaction, err := instance.dbPool.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return err
	}

	// Настраиваем функцию отсрочки для отката или коммита транзакции.
	defer func() {
		// Восстанавливаемся после паники
		recoverResult := recover();
		if recoverResult != nil {
			err = errors.Errorf("panic recovered: %v", recoverResult)
		}

		// Откатываем транзакцию, если произошла ошибка
		if err != nil {
			errRollback := transaction.Rollback(ctx);
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
	err = fn( ctx, serviceProvider )
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
