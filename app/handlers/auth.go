package handlers

import (
	"fmt"

	"github.com/caseyrwebb/auth-microservice/app/data"
	"github.com/caseyrwebb/auth-microservice/app/services"
	"github.com/caseyrwebb/auth-microservice/app/utils"
	"go.uber.org/zap"
)

// UserKey is used as a key for storing the User object in context at middleware
type UserKey struct{}

// UserIDKey is used as a key for storing the UserID in context at middleware
type UserIDKey struct{}

// VerificationDataKey is used as the key for storing the VerificationData in context at middleware
type VerificationDataKey struct{}

// UserHandler wraps instances needed to perform operations on user object
type AuthHandler struct {
	logger      zap.Logger
	configs     *utils.Configurations
	validator   *data.Validation
	repo        data.Repository
	authService services.Authentication
	mailService services.MailService
}

// NewUserHandler returns a new UserHandler instance
func NewAuthHandler(l zap.Logger, c *utils.Configurations, v *data.Validation, r data.Repository, auth services.Authentication, mail services.MailService) *AuthHandler {
	return &AuthHandler{
		logger:      l,
		configs:     c,
		validator:   v,
		repo:        r,
		authService: auth,
		mailService: mail,
	}
}

// GenericResponse is the format of our response
type GenericResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Errors []string `json:"errors"`
}

// Below data types are used for encoding and decoding b/t go types and json
type TokenResponse struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

type AuthResponse struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
	Username     string `json:"username"`
}

type UsernameUpdate struct {
	Username string `json:"username"`
}

type CodeVerificationReq struct {
	Code string `json: "code"`
	Type string `json" "type"`
}

type PasswordResetReq struct {
	Password   string `json: "password"`
	PasswordRe string `json: "password_re"`
	Code       string `json: "code"`
}

var ErrUserAlreadyExists = fmt.Sprintf("User already exists with the given email")
var ErrUserNotFound = fmt.Sprintf("No user account exists with given email. Please sign in first")
var UserCreationFailed = fmt.Sprintf("Unable to create user.Please try again later")

var PgDuplicateKeyMsg = "duplicate key value violates unique constraint"
var PgNoRowsMsg = "no rows in result set"
