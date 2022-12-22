package data

import (
	"context"
	"fmt"
	"time"

	"github.com/caseyrwebb/auth-microservice/app/models"
	uuid "github.com/satori/go.uuid"
)

// // PostgresRepository has the implementation of the db methods.
// type PostgresRepository struct {
// 	db     *sqlx.DB
// 	logger zap.Logger
// }

// // NewPostgresRepository returns a new PostgresRepository instance
// func NewPostgresRepository(db *sqlx.DB, logger zap.Logger) *PostgresRepository {
// 	return &PostgresRepository{db, logger}
// }

// Create inserts the given user into the database
func (d *DB) Create(ctx context.Context, user *models.User) error {
	user.ID = uuid.NewV4().String()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	d.logger.Info(fmt.Sprintf("%s %v", "creating user", user))
	query := "insert into users (id, email, username, password, tokenhash, createdat, updatedat) values ($1, $2, $3, $4, $5, $6, $7)"
	_, err := d.db.ExecContext(ctx, query, user.ID, user.Email, user.Username, user.Password, user.Token, user.CreatedAt, user.UpdatedAt)
	return err
}

// GetUserByEmail retrieves the user object having the given email, else returns error
func (d *DB) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	d.logger.Debug(fmt.Sprintf("querying for user with email", email))
	query := "select * from users where email = $1"
	var user models.User
	if err := d.db.GetContext(ctx, &user, query, email); err != nil {
		return nil, err
	}
	d.logger.Debug(fmt.Sprintf("%s %v", "read users", user))
	return &user, nil
}

// GetUserByID retrieves the user object having the given ID, else returns error
func (repo *PostgresRepository) GetUserByID(ctx context.Context, userID string) (*User, error) {
	repo.logger.Debug("querying for user with id", userID)
	query := "select * from users where id = $1"
	var user User
	if err := repo.db.GetContext(ctx, &user, query, userID); err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUsername updates the username of the given user
func (repo *PostgresRepository) UpdateUsername(ctx context.Context, user *User) error {
	user.UpdatedAt = time.Now()

	query := "update users set username = $1, updatedat = $2 where id = $3"
	if _, err := repo.db.ExecContext(ctx, query, user.Username, user.UpdatedAt, user.ID); err != nil {
		return err
	}
	return nil
}

// UpdateUserVerificationStatus updates user verification status to true
func (repo *PostgresRepository) UpdateUserVerificationStatus(ctx context.Context, email string, status bool) error {

	query := "update users set isverified = $1 where email = $2"
	if _, err := repo.db.ExecContext(ctx, query, status, email); err != nil {
		return err
	}
	return nil
}

// StoreMailVerificationData adds a mail verification data to db
func (repo *PostgresRepository) StoreVerificationData(ctx context.Context, verificationData *VerificationData) error {

	query := "insert into verifications(email, code, expiresat, type) values($1, $2, $3, $4)"
	_, err := repo.db.ExecContext(ctx, query, verificationData.Email, verificationData.Code, verificationData.ExpiresAt, verificationData.Type)
	return err
}

// GetMailVerificationCode retrieves the stored verification code.
func (repo *PostgresRepository) GetVerificationData(ctx context.Context, email string, verificationDataType VerificationDataType) (*VerificationData, error) {

	query := "select * from verifications where email = $1 and type = $2"

	var verificationData VerificationData
	if err := repo.db.GetContext(ctx, &verificationData, query, email, verificationDataType); err != nil {
		return nil, err
	}
	return &verificationData, nil
}

// DeleteMailVerificationData deletes a used verification data
func (repo *PostgresRepository) DeleteVerificationData(ctx context.Context, email string, verificationDataType VerificationDataType) error {

	query := "delete from verifications where email = $1 and type = $2"
	_, err := repo.db.ExecContext(ctx, query, email, verificationDataType)
	return err
}

// UpdatePassword updates the user password
func (repo *PostgresRepository) UpdatePassword(ctx context.Context, userID string, password string, tokenHash string) error {

	query := "update users set password = $1, tokenhash = $2 where id = $3"
	_, err := repo.db.ExecContext(ctx, query, password, tokenHash, userID)
	return err
}
