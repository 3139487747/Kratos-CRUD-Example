package data

import (
	"context"
	"gorm.io/gorm"

	"shop/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type CrudRepo struct {
	data  *Data
	log   *log.Helper
	table *gorm.DB
}

func NewCrudRepo(data *Data, logger log.Logger) biz.CrudRepo {
	return &CrudRepo{
		data:  data,
		log:   log.NewHelper(logger),
		table: data.db.Table("user"),
	}
}

func (r *CrudRepo) Save(ctx context.Context, g *biz.User) (*biz.User, error) {
	result := r.table.WithContext(ctx).Create(g)
	if result.Error != nil {
		return nil, result.Error
	}
	return g, nil
}

func (r *CrudRepo) Update(ctx context.Context, g *biz.User) (*biz.User, error) {
	result := r.table.Model(g).Updates(biz.User{Name: g.Name, Pwd: g.Pwd})
	if result.Error != nil {
		return nil, result.Error
	}
	return g, nil
}

func (r *CrudRepo) FindByName(ctx context.Context, name string) (*biz.User, error) {
	user := &biz.User{}
	r.table.WithContext(ctx).First(user, "name = ?", name)
	if user.Id == 0 {
		return &biz.User{}, nil
	}
	return user, nil
}

func (r *CrudRepo) ListAll(context.Context) ([]*biz.User, error) {

	var users []biz.User
	result := r.table.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	g := make([]*biz.User, 0)
	for _, p := range users {
		g = append(g, &biz.User{
			Id:   p.Id,
			Name: p.Name,
			Pwd:  p.Pwd,
		})
	}
	return g, nil
}

func (r *CrudRepo) Delete(ctx context.Context, id int64) (*biz.User, error) {
	result := r.table.Delete(&biz.User{Id: id})
	if result.Error != nil {
		return nil, result.Error
	}
	return &biz.User{}, nil
}
