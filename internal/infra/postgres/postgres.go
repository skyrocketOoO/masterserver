package postgres

import (
	"context"
	"fmt"

	"github.com/skyrocketOoO/masterserver/domain"
	"gorm.io/gorm"
)

type OrmRepository struct {
	db *gorm.DB
}

func NewOrmRepository(db *gorm.DB) *OrmRepository {
	db.AutoMigrate(&User{})
	return &OrmRepository{
		db: db,
	}
}

func (r *OrmRepository) Ping(c context.Context) error {
	db, err := r.db.DB()
	if err != nil {
		return err
	}
	if err := db.PingContext(c); err != nil {
		return err
	}

	return nil
}

func (r *OrmRepository) GetUsers(c context.Context, filter map[string]interface{},
	sort domain.Sort, pagination domain.Pagination) ([]User, error) {

	users := []User{}
	if err := r.db.Order(fmt.Sprintf("%s %s", sort.Field, sort.Order)).
		Offset((pagination.Page - 1) * pagination.PerPage).Limit(pagination.PerPage).
		Where(filter).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *OrmRepository) GetUserById(c context.Context, id uint) (*User, error) {
	user := User{ID: id}
	if err := r.db.Take(&user).Error; err != nil {
		return &user, err
	}
	return &user, nil
}

func (r *OrmRepository) CreateUser(c context.Context, user User) (User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func (r *OrmRepository) UpdateUser(c context.Context, id uint,
	updates map[string]interface{}) error {
	user := User{ID: id}
	return r.db.Model(&user).Updates(updates).Error
}

func (r *OrmRepository) DeleteUser(c context.Context, id uint) error {
	return r.db.Delete(&User{}, id).Error
}
