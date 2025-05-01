package repositories

import (
	dbmodel "4it428-newsletter-api/services/user-service/internal/infrastructure/persistence/model"
	"4it428-newsletter-api/services/user-service/internal/infrastructure/persistence/query"
	"4it428-newsletter-api/services/user-service/internal/service/errors"
	"4it428-newsletter-api/services/user-service/internal/service/model"
	"context"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool: pool}
}

func (r *UserRepository) List(ctx context.Context) ([]*model.User, error) {
	var users []dbmodel.UserEntity
	err := pgxscan.Select(ctx, r.pool, &users, query.SelectUsers)
	if err != nil {
		return nil, err
	}

	response := make([]*model.User, len(users))
	for i, user := range users {
		response[i] = &model.User{
			ID:          user.ID,
			FirebaseUID: user.FirebaseUID,
			Email:       user.Email,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
		}
	}
	return response, nil
}

func (r *UserRepository) Create(ctx context.Context, user *model.User) error {
	_, err := r.pool.Exec(ctx, query.InsertUser,
		user.ID, user.Email, user.FirstName, user.LastName, user.FirebaseUID,
	)
	return err
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	var user dbmodel.UserEntity
	err := pgxscan.Get(ctx, r.pool, &user, query.SelectUserByEmail, email)
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:          user.ID,
		Email:       user.Email,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		FirebaseUID: user.FirebaseUID,
	}, nil
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (*model.User, error) {
	var user dbmodel.UserEntity
	err := pgxscan.Get(ctx, r.pool, &user, query.SelectUserById, id)
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:          user.ID,
		Email:       user.Email,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		FirebaseUID: user.FirebaseUID,
	}, nil
}

func (r *UserRepository) Update(ctx context.Context, user *model.User) (*model.User, error) {
	_, err := r.pool.Exec(ctx, query.UpdateUser, user.FirstName, user.LastName, user.Email)
	if err != nil {
		return nil, err
	}

	updatedUser, err := r.GetByEmail(ctx, user.Email)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}

func (r *UserRepository) Delete(ctx context.Context, email string) error {
	commandTag, err := r.pool.Exec(ctx, query.DeleteUser, email)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return errors.ErrUserNotFound
	}
	return err
}
