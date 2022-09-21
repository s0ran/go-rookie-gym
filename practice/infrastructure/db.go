package infrastructure

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/s0ran/go-rookie-gym/domain/user"
)

func NewDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/database")
	if err != nil {
		return nil, err
	}
	if err := db.PingContext(context.Background()); err != nil {
		return nil, err
	}
	return db, nil
}

type repository struct {
	db *sql.DB
}

func (r *repository) DB() *sql.DB {
	return r.db
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db: db}
}

func (r *repository) PostUser(ctx context.Context, user *user.User) (int64, error) {
	res, err := r.DB().ExecContext(ctx, "INSERT INTO `users` (name) VALUES (?)", user.Name)
	if err != nil {
		return 0, err
	} else {
		fmt.Println(res)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return id, err
	}

	return id, nil
}
