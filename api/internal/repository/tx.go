package repository

import (
	"context"
	"database/sql"

	ctxkey "github.com/keis8221/surprise/api/internal"
	"gorm.io/gorm"
)


type TxRepository struct {
    db *gorm.DB
}

func NewTransaction(db *gorm.DB) *TxRepository {
    return &TxRepository{db: db}
}

func (t *TxRepository) DoInTx(ctx context.Context, f func(ctx context.Context) (interface{}, error)) (interface{}, error) {
    var v interface{}
    var err error
    t.db.Transaction(func(tx *gorm.DB) error {
        ctx = context.WithValue(ctx, ctxkey.TxCtxKey, tx)
        v, err = f(ctx)
        if err != nil {
            tx.Rollback()
            return err 
        }
        tx.Commit()
        return err
    })
    return v, err
}

func GetTx(ctx context.Context) (*sql.Tx, bool) {
    tx, ok := ctx.Value(ctxkey.TxCtxKey).(*sql.Tx)
    return tx, ok
}

