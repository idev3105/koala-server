package userdomain

import (
	"context"
	"time"

	"org.idev.koala/backend/common/logger"
	commonenum "org.idev.koala/backend/domain/common/enum"
	userentity "org.idev.koala/backend/domain/user/entity"
)

type userUseCase struct {
	repo UserRepository
}

func NewUserUseCase(repo UserRepository) *userUseCase {
	return &userUseCase{
		repo: repo,
	}
}

func (s *userUseCase) ExistsByUserId(ctx context.Context, userId string) (bool, error) {
	return s.repo.ExistsByUserId(ctx, userId)
}

func (s *userUseCase) FindByUserId(ctx context.Context, userId string) (*userentity.User, error) {
	return s.repo.FindByUserId(ctx, userId)
}

func (s *userUseCase) Create(ctx context.Context, userId string, username string) (*userentity.User, error) {
	log := logger.New("User Use Case", "Create")
	log.Infof("Creating user with user id and username: %s %s", userId, username)
	user := &userentity.User{
		UserId:   userId,
		Username: username,
	}
	now := time.Now()
	actor := commonenum.ACTOR_SYSTEM.String()
	user.CreatedAt = &now
	user.CreatedBy = &actor
	user.UpdatedBy = &actor
	return s.repo.Save(ctx, user)
}
