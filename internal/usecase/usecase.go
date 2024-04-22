package usecase

import (
	"context"

	"github.com/skyrocketOoO/masterserver/domain"
	"github.com/skyrocketOoO/masterserver/internal/infra/postgres"
)

type Usecase struct {
	repo *postgres.OrmRepository
}

func NewUsecase(ormRepo *postgres.OrmRepository) *Usecase {
	return &Usecase{
		repo: ormRepo,
	}
}

func (u *Usecase) Healthy(c context.Context) error {
	// do something check like db connection is established
	if err := u.repo.Ping(c); err != nil {
		return err
	}

	return nil
}

func (u *Usecase) GetUsers(ctx context.Context, filter map[string]interface{},
	sort domain.Sort, pagination domain.Pagination) (
	[]postgres.User, domain.PageInfo, error) {
	return u.repo.GetUsers(ctx, filter, sort, pagination)
}

func (u *Usecase) GetUser(ctx context.Context, id uint) (postgres.User, error) {
	user, err := u.repo.GetUserById(ctx, id)
	if err != nil {
		return postgres.User{}, err
	}
	return user, nil
}

func (u *Usecase) GetManyReference(ctx context.Context, target string, id uint,
	filter map[string]interface{}, sort domain.Sort, rang domain.Pagination) {
}

func (u *Usecase) CreateUser(ctx context.Context, user postgres.User) (
	postgres.User, error) {
	return u.repo.CreateUser(ctx, user)
}

func (u *Usecase) UpdateUser(ctx context.Context, id uint, preData postgres.User,
	updates map[string]interface{}) (postgres.User, error) {
	return u.repo.UpdateUser(ctx, id, preData, updates)
}

func (u *Usecase) UpdateUsers(ctx context.Context, ids []uint,
	updates map[string]interface{}) error {
	return u.repo.UpdateUsers(ctx, ids, updates)
}

func (u *Usecase) DeleteUser(ctx context.Context, id uint,
	preData postgres.User) (postgres.User, error) {
	return u.repo.DeleteUser(ctx, id, preData)
}

func (u *Usecase) DeleteUsers(ctx context.Context, ids []uint) error {
	return u.repo.DeleteUsers(ctx, ids)
}
