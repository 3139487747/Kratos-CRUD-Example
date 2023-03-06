package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Id   int64
	Name string
	Pwd  string
}

type CrudRepo interface {
	Save(context.Context, *User) (*User, error)
	Update(context.Context, *User) (*User, error)
	FindByName(context.Context, string) (*User, error)
	ListAll(context.Context) ([]*User, error)
	Delete(context.Context, int64) (*User, error)
}

type CrudUsecase struct {
	repo CrudRepo
	log  *log.Helper
}

func NewCrudUsecase(repo CrudRepo, logger log.Logger) *CrudUsecase {
	return &CrudUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *CrudUsecase) Save(ctx context.Context, g *User) (*User, error) {
	return uc.repo.Save(ctx, g)
}

func (uc *CrudUsecase) Update(ctx context.Context, g *User) (*User, error) {
	return uc.repo.Update(ctx, g)
}

func (uc *CrudUsecase) FindByName(ctx context.Context, name string) (*User, error) {
	return uc.repo.FindByName(ctx, name)
}

func (uc *CrudUsecase) ListAll(ctx context.Context) ([]*User, error) {
	return uc.repo.ListAll(ctx)
}

func (uc *CrudUsecase) Delete(ctx context.Context, id int64) (*User, error) {
	return uc.repo.Delete(ctx, id)
}
