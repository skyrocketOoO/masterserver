package postgres

import (
	"context"

	"gorm.io/gorm"
)

type OrmRepository struct {
	db *gorm.DB
}

func NewOrmRepository(db *gorm.DB) *OrmRepository {
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

func (r *OrmRepository) GetUsers(c context.Context,
	filter map[string]interface{}) ([]User, error) {
	users := []User{}
	if err := r.db.Where(filter).Find(&users).Error; err != nil {
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

func (r *OrmRepository) CreateUser(c context.Context, user *User) error {
	return r.db.Create(user).Error
}

func (r *OrmRepository) UpdateUser(c context.Context, id uint,
	updates map[string]interface{}) error {
	user := User{ID: id}
	return r.db.Model(&user).Updates(updates).Error
}

func (r *OrmRepository) DeleteUser(c context.Context, id uint) error {
	return r.db.Delete(&User{}, id).Error
}
