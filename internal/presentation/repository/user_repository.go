package repository

import (
	"api-wave-bot/internal/presentation/domain"
	"context"
	"database/sql"
	"errors"

	"github.com/Masterminds/squirrel"
)

// Interface define as operações do repositório de usuários
type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) (string, error)
	FindUserByEmail(ctx context.Context, email string) (*domain.User, error)
	FindUserByID(ctx context.Context, id string) (*domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) error
	DeleteUser(ctx context.Context, id string) error
	ListUsersByCompany(ctx context.Context, companyID string) ([]*domain.User, error)
}

// Implementação concreta do repositório
type userRepository struct {
	db *sql.DB
	sq squirrel.StatementBuilderType
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
		sq: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

// Criação de um novo usuário
func (r *userRepository) CreateUser(ctx context.Context, user *domain.User) (string, error) {
	query, args, err := r.sq.Insert("users").
		Columns("name", "email", "password_hash", "company_id").
		Values(user.Name, user.Email, user.PasswordHash, user.CompanyID).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return "", err
	}

	var id string
	err = r.db.QueryRowContext(ctx, query, args...).Scan(&id)
	return id, err
}

// Busca usuário pelo e-mail (login)
func (r *userRepository) FindUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	query, args, err := r.sq.Select("id", "name", "email", "password_hash", "company_id").
		From("users").
		Where(squirrel.Eq{"email": email}).
		ToSql()
	if err != nil {
		return nil, err
	}

	row := r.db.QueryRowContext(ctx, query, args...)
	var user domain.User
	err = row.Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.CompanyID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	return &user, err
}

// Busca por ID
func (r *userRepository) FindUserByID(ctx context.Context, id string) (*domain.User, error) {
	query, args, err := r.sq.Select("id", "name", "email", "password_hash", "company_id").
		From("users").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return nil, err
	}

	row := r.db.QueryRowContext(ctx, query, args...)
	var user domain.User
	err = row.Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.CompanyID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	return &user, err
}

// Atualiza os dados do usuário
func (r *userRepository) UpdateUser(ctx context.Context, user *domain.User) error {
	query, args, err := r.sq.Update("users").
		Set("name", user.Name).
		Set("email", user.Email).
		Set("password_hash", user.PasswordHash).
		Where(squirrel.Eq{"id": user.ID}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, query, args...)
	return err
}

// Deleta usuário
func (r *userRepository) DeleteUser(ctx context.Context, id string) error {
	query, args, err := r.sq.Delete("users").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, query, args...)
	return err
}

// Lista todos os usuários da mesma empresa
func (r *userRepository) ListUsersByCompany(ctx context.Context, companyID string) ([]*domain.User, error) {
	query, args, err := r.sq.Select("id", "name", "email", "password_hash", "company_id").
		From("users").
		Where(squirrel.Eq{"company_id": companyID}).
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*domain.User
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.CompanyID); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}
