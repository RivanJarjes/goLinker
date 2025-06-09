package memory

import (
	"context"
	"errors"
	"sync"
	"time"
	"github.com/RivanJarjes/goLinker/internal/core"
)

type linkRepo struct {
	mtx   sync.RWMutex
	auto  int64
	store map[string]core.Link
}

func NewLinkRepo() core.LinkRepo {
	return &linkRepo{store: make(map[string]core.Link)}
}

func (r *linkRepo) Insert(_ context.Context, l *core.Link) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.auto++
	l.ID, l.Created = r.auto, time.Now()
	r.store[l.Code] = *l
	return nil
}

func (r *linkRepo) FindByCode(_ context.Context, code string) (*core.Link, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	if v, ok := r.store[code]; ok {
		return &v, nil
	}
	return nil, errors.New("not found")
}
