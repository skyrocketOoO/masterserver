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
	sort domain.Sort, pagination domain.Pagination) ([]User, domain.PageInfo, error) {

	users := []User{}
	db := r.db.Model(&User{}).Where(filter)

	// Apply sorting
	if sort.Field != "" {
		db = db.Order(fmt.Sprintf("%s %s", sort.Field, sort.Order))
	}

	// Count total number of records
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, domain.PageInfo{}, err
	}

	// Apply pagination
	if pagination.Page > 0 && pagination.PerPage > 0 {
		db = db.Offset((pagination.Page - 1) * pagination.PerPage).Limit(pagination.PerPage)
	}

	// Fetch records
	if err := db.Find(&users).Error; err != nil {
		return nil, domain.PageInfo{}, err
	}

	// Calculate PageInfo
	pageInfo := domain.PageInfo{
		HasNextPage: int64(pagination.Page)*int64(pagination.PerPage) < total,
		HasPrevPage: pagination.Page > 1,
	}

	return users, pageInfo, nil
}

func (r *OrmRepository) GetUserById(c context.Context, id uint) (User, error) {
	user := User{ID: id}
	if err := r.db.Take(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *OrmRepository) CreateUser(c context.Context, user User) (User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func (r *OrmRepository) UpdateUser(c context.Context, id uint, preData User,
	updates map[string]interface{}) (User, error) {
	user := User{ID: id}
	if err := r.db.Model(&user).Where(preData).Updates(updates).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *OrmRepository) UpdateUsers(c context.Context, ids []uint,
	updates map[string]interface{}) error {
	if err := r.db.Model(User{}).Where("id IN ?", ids).Updates(updates).
		Error; err != nil {
		return err
	}
	return nil
}

func (r *OrmRepository) DeleteUser(c context.Context, id uint,
	preData User) (User, error) {
	user := User{ID: id}
	if err := r.db.Model(&user).Delete(&preData).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *OrmRepository) DeleteUsers(c context.Context, ids []uint) error {
	users := []User{}
	for _, id := range ids {
		users = append(users, User{ID: id})
	}
	if err := r.db.Delete(&users).Error; err != nil {
		return err
	}
	return nil
}
