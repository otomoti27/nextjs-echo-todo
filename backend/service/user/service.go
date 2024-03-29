package user

import (
	"echo-api/domain"
	"os"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	GetByEmail(user *domain.User, email string) error
	GetByID(user *domain.User, id uint) error
	Create(user *domain.User) error
	Update(user *domain.User) error
	Delete(user *domain.User) error
}

type Service struct {
	userRepo UserRepository
}

func NewService(ur UserRepository) *Service {
	return &Service{
		userRepo: ur,
	}
}

func (s *Service) SignUp(user domain.User) (domain.UserResponse, error) {
	err := validation.ValidateStruct(&user,
		validation.Field(&user.Email, domain.UserEmailRule...),
		validation.Field(&user.Password, domain.UserPasswordRule...),
		validation.Field(&user.Name, domain.UserNameRule...),
	)
	if err != nil {
		return domain.UserResponse{}, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return domain.UserResponse{}, err
	}

	newUser := domain.User{Email: user.Email, Password: string(hashedPassword), Name: user.Name}
	err = s.userRepo.Create(&newUser)
	if err != nil {
		return domain.UserResponse{}, err
	}

	resUser := domain.UserResponse{ID: newUser.ID, Email: newUser.Email, Name: newUser.Name}
	return resUser, nil
}

func (s *Service) LogIn(user domain.User) (string, error) {
	err := validation.ValidateStruct(&user,
		validation.Field(&user.Email, domain.UserEmailRule...),
		validation.Field(&user.Password, domain.UserPasswordRule...),
	)
	if err != nil {
		return "", err
	}

	var u domain.User
	err = s.userRepo.GetByEmail(&u, user.Email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": u.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	return token.SignedString([]byte(os.Getenv("SECRET")))
}

func (s *Service) FetchMe(id uint) (domain.UserResponse, error) {
	var u domain.User
	err := s.userRepo.GetByID(&u, id)
	if err != nil {
		return domain.UserResponse{}, err
	}

	resUser := domain.UserResponse{ID: u.ID, Email: u.Email, Name: u.Name}
	return resUser, nil
}
