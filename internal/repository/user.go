package repository

import (
	"context"
	//"database/sql"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prem0x01/hospital/internal/database/queries"
	"github.com/prem0x01/hospital/internal/domain"
)

type UserRepository struct {
	db      *pgxpool.Pool
	queries *queries.Queries
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db:      pool,
		queries: queries.New(pool),
	}
}

func (r *UserRepository) Create(ctx context.Context, user *domain.User) error {
	params := queries.CreateUserParams{
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		Role:         user.Role,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
	}

	res, err := r.queries.CreateUser(ctx, params)
	if err != nil {
		return err
	}

	user.ID = res.ID
	user.CreatedAt = res.CreatedAt
	user.UpdatedAt = res.UpdatedAt
	return nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	res, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:           res.ID,
		Email:        res.Email,
		PasswordHash: res.PasswordHash,
		Role:         res.Role,
		FirstName:    res.FirstName,
		LastName:     res.LastName,
		CreatedAt:    res.CreatedAt,
		UpdatedAt:    res.UpdatedAt,
	}, nil
}

func (r *UserRepository) GetByID(ctx context.Context, id int32) (*domain.User, error) {
	res, err := r.queries.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:           res.ID,
		Email:        res.Email,
		PasswordHash: res.PasswordHash,
		Role:         res.Role,
		FirstName:    res.FirstName,
		LastName:     res.LastName,
		CreatedAt:    res.CreatedAt,
		UpdatedAt:    res.UpdatedAt,
	}, nil
}

func (r *UserRepository) Update(ctx context.Context, user *domain.User) error {
	params := queries.UpdateUserParams{
		ID:           user.ID,
		Email:        &user.Email,
		PasswordHash: &user.PasswordHash,
		Role:         &user.Role,
		FirstName:    &user.FirstName,
		LastName:     &user.LastName,
	}

	res, err := r.queries.UpdateUser(ctx, params)
	if err != nil {
		return err
	}

	user.Email = res.Email
	user.PasswordHash = res.PasswordHash
	user.Role = res.Role
	user.FirstName = res.FirstName
	user.LastName = res.LastName
	user.CreatedAt = res.CreatedAt
	user.UpdatedAt = res.UpdatedAt
	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id int32) error {
	return r.queries.DeleteUser(ctx, id)
}

func (r *UserRepository) GetDoctors(ctx context.Context) ([]domain.User, error) {
	results, err := r.queries.GetDoctors(ctx)
	if err != nil {
		return nil, err
	}

	var doctors []domain.User
	for _, d := range results {
		doctors = append(doctors, domain.User{
			ID:        d.ID,
			Email:     d.Email,
			Role:      d.Role,
			FirstName: d.FirstName,
			LastName:  d.LastName,
			CreatedAt: d.CreatedAt,
			UpdatedAt: d.UpdatedAt,
		})
	}

	return doctors, nil
}
