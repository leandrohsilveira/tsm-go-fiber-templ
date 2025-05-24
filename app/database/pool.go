package database

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/leandrohsilveira/tsm/dao"
	"github.com/rs/zerolog/log"
)

type DatabasePool interface {
	Text(string) pgtype.Text
	Acquire(context.Context) (*dao.Queries, func(), error)
	Close()
}

type databasePool struct {
	pool *pgxpool.Pool
}

func NewPool(ctx context.Context) (DatabasePool, error) {
	connString, isSet := os.LookupEnv("DATABASE_URL")
	if !isSet {
		connString = "postgres://app:password@localhost:5432/app?sslmode=disable"
	}

	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	log.Info().
		Str("host", config.ConnConfig.Host).
		Uint16("port", config.ConnConfig.Port).
		Msg("Database connected")

	return &databasePool{pool}, nil
}

func (*databasePool) Text(text string) pgtype.Text {
	return pgtype.Text{String: text, Valid: true}
}

func (d *databasePool) Acquire(ctx context.Context) (*dao.Queries, func(), error) {
	conn, err := d.pool.Acquire(ctx)

	if err != nil {
		return nil, nil, err
	}

	queries := dao.New(conn)

	return queries, conn.Release, nil
}

func (d *databasePool) Close() {
	d.pool.Close()
	log.Info().Msg("Database disconnected")
}
