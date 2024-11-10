package models

import (
	sq "github.com/Masterminds/squirrel"

	"web_lab/internal/storage/pg"
)

var qb = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

const (
	uniqueViolationCode = "23505"
)

type PG struct {
	db *pg.Storage
}

func NewPG(db *pg.Storage) *PG {
	return &PG{
		db: db,
	}
}

type PGStd struct {
	db *pg.StorageStd
}

func NewPGStd(db *pg.StorageStd) *PGStd {
	return &PGStd{
		db: db,
	}
}
