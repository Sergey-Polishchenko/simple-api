package postgres

import (
	"context"
	"errors"

	"gorm.io/gorm"

	app "github.com/Sergey-Polishchenko/simple-api/internal/application"
	"github.com/Sergey-Polishchenko/simple-api/internal/domain"
	perrors "github.com/Sergey-Polishchenko/simple-api/internal/pkg/errors"
)

type UserPG struct {
	ID   string `gorm:"primaryKey"`
	Name string
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) (app.UserRepository, error) {
	repo := &UserRepo{db: db}
	err := repo.migrate()
	return repo, err
}

func (ur *UserRepo) migrate() error {
	return ur.db.AutoMigrate(&UserPG{})
}

func (ur *UserRepo) Create(ctx context.Context, user *domain.User) error {
	pgUser := &UserPG{
		ID:   user.ID(),
		Name: user.Name(),
	}

	return ur.db.WithContext(ctx).Create(pgUser).Error
}

func (ur *UserRepo) GetAll(ctx context.Context) ([]*domain.User, error) {
	var pgUsers []UserPG
	if err := ur.db.WithContext(ctx).Find(&pgUsers).Error; err != nil {
		return nil, err
	}

	users := make([]*domain.User, 0, len(pgUsers))
	for _, u := range pgUsers {
		users = append(users, domain.NewUser(u.ID, u.Name))
	}

	return users, nil
}

func (ur *UserRepo) GetByID(ctx context.Context, id string) (*domain.User, error) {
	var pgUser UserPG
	err := ur.db.WithContext(ctx).Where("id = ?", id).First(&pgUser).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, perrors.ErrUserNotFound
		}
		return nil, err
	}

	return domain.NewUser(pgUser.ID, pgUser.Name), nil
}

func (ur *UserRepo) Update(ctx context.Context, user *domain.User) error {
	result := ur.db.WithContext(ctx).
		Model(&UserPG{}).
		Where("id = ?", user.ID()).
		Updates(UserPG{
			Name: user.Name(),
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return perrors.ErrUserNotFound
	}

	return nil
}

func (ur *UserRepo) Remove(ctx context.Context, id string) error {
	result := ur.db.WithContext(ctx).Where("id = ?", id).Delete(&UserPG{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return perrors.ErrUserNotFound
	}

	return nil
}
