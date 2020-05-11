package handlers

import (
	"errors"
	"net/http"

	"github.com/bahattincinic/messenger-challenge/domain/models"
	"github.com/bahattincinic/messenger-challenge/domain/repositories"
	"github.com/jinzhu/gorm"
)

// GetUser returns current user
func GetUser(r *http.Request) (user models.User, err error) {
	user, found := r.Context().Value("user").(models.User)

	if found == false {
		err = errors.New("Forbidden")
	}

	return
}

// BaseHandler will hold everything that controller needs
type BaseHandler struct {
	userRepo    repositories.IUserRepository
	authRepo    repositories.IAuthRepository
	messageRepo repositories.IMessageRepository
}

// NewBaseHandler returns a new BaseHandler
func NewBaseHandler(db *gorm.DB) *BaseHandler {
	userRepo := repositories.NewUserRepo(db)
	authRepo := repositories.NewAuthRepo(db)
	messageRepo := repositories.NewMessageRepo(db)

	return &BaseHandler{
		userRepo:    userRepo,
		authRepo:    authRepo,
		messageRepo: messageRepo,
	}
}
