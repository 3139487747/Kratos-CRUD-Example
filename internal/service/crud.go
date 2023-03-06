package service

import (
	"context"
	v1 "shop/api/crud/v1"
	"shop/internal/biz"
)

type CrudService struct {
	v1.UnimplementedCrudServer

	uc *biz.CrudUsecase
}

func NewCrudService(uc *biz.CrudUsecase) *CrudService {
	return &CrudService{uc: uc}
}

func (s *CrudService) Save(ctx context.Context, in *v1.CreateRequest) (*v1.CreateReply, error) {
	g, err := s.uc.Save(ctx, &biz.User{
		Name: in.Name,
		Pwd:  in.Pwd,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateReply{
		Name: g.Name,
		Pwd:  g.Pwd,
		Id:   g.Id,
	}, nil
}

func (s *CrudService) Update(ctx context.Context, in *v1.UpdateRequest) (*v1.UpdateReply, error) {
	g, err := s.uc.Update(ctx, &biz.User{
		Name: in.Name,
		Pwd:  in.Pwd,
		Id:   in.Id,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UpdateReply{
		Name: g.Name,
		Pwd:  g.Pwd,
		Id:   g.Id,
	}, nil
}

func (s *CrudService) ListAll(ctx context.Context, in *v1.ListRequest) (*v1.ListReply, error) {
	ps, err := s.uc.ListAll(ctx)
	if err != nil {
		return nil, err
	}
	reply := &v1.ListReply{}
	for _, p := range ps {
		reply.Users = append(reply.Users, &v1.User{
			Id:   p.Id,
			Name: p.Name,
			Pwd:  p.Pwd,
		})
	}
	return reply, nil
}

func (s *CrudService) FindByName(ctx context.Context, in *v1.FindRequest) (*v1.FindReply, error) {
	g, err := s.uc.FindByName(ctx, in.Name)
	if err != nil {
		return nil, err
	}
	return &v1.FindReply{
		Name: g.Name,
		Pwd:  g.Pwd,
		Id:   g.Id,
	}, nil
}

func (s *CrudService) Delete(ctx context.Context, in *v1.DeleteRequest) (*v1.DeleteReply, error) {
	g, err := s.uc.Delete(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteReply{
		Name: g.Name,
		Pwd:  g.Pwd,
		Id:   g.Id,
	}, nil
}
