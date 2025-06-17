package repository

import (
	"context"
	"database/sql"

	"github.com/prem0x01/hospital/internal/database/queries"
	"github.com/prem0x01/hospital/internal/models"
)

type UserRepository struct {
	db      *sql.DB
	queries *queries.Queries
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db:      db,
		queries: queries.New(db),
	}
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
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
	user.CreatedAt = res.CreatedAt.Time
	user.UpdatedAt = res.UpdatedAt.Time
	return nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	res, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:           res.ID,
		Email:        res.Email,
		PasswordHash: res.PasswordHash,
		Role:         res.Role,
		FirstName:    res.FirstName,
		LastName:     res.LastName,
		CreatedAt:    res.CreatedAt.Time,
		UpdatedAt:    res.UpdatedAt.Time,
	}, nil
}

func (r *UserRepository) GetByID(ctx context.Context, id int32) (*models.User, error) {
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
		CreatedAt:    res.CreatedAt.Time,
		UpdatedAt:    res.UpdatedAt.Time,
	}, nil
}

func (r *UserRepository) Update(ctx context.Context, user *models.User) error {
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
	user.CreatedAt = res.CreatedAt.Time
	user.UpdatedAt = res.UpdatedAt.Time
	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id int32) error {
	return r.queries.DeleteUser(ctx, id)
}

func (r *UserRepository) GetDoctors(ctx context.Context) ([]models.User, error) {
	results, err := r.queries.GetDoctors(ctx)
	if err != nil {
		return nil, err
	}

	var doctors []models.User
	for _, d := range results {
		doctors = append(doctors, models.User{
			ID:        d.ID,
			Email:     d.Email,
			Role:      d.Role,
			FirstName: d.FirstName,
			LastName:  d.LastName,
			CreatedAt: d.CreatedAt.Time,
			UpdatedAt: d.UpdatedAt.Time,
		})
	}

	return doctors, nil
}
