package core

import "context"

type LinkRepo interface {
	Insert(ctx context.Context, l *Link) error
	FindByCode(ctx context.Context, code string) (*Link, error)
}
