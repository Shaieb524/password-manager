package user

import (
	"time"

	"html"
	"strings"

	"github.com/gofrs/uuid"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthenticationRepo struct {
	logger *zap.Logger
	db     *gorm.DB
}

type User struct {
	gorm.Model
	Id        uuid.UUID `gorm:"size: 128; not null;unique" json:"id"`
	Username  string    `gorm:"size: 255;not null;unique" json:"username"`
	Password  string    `gorm:"size: 255;not null;" json:"password"`
	CreatedAt time.Time `gorm:"size 255;not null;"`
}

func (repo *AuthenticationRepo) Save(user *User) (*User, error) {
	user.Id = uuid.Must(uuid.NewV4())
	user.CreatedAt = time.Now()
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	hashedPassword, err := repo.hashPassword(user.Password)
	if err != nil {
		return &User{}, err
	}
	user.Password = hashedPassword

	err = repo.db.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (repo *AuthenticationRepo) hashPassword(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	strHashedPw := string(passwordHash)
	return strHashedPw, nil
}

func (repo *AuthenticationRepo) ValidatePassword(user *User, inputpassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(inputpassword))
}

func (repo *AuthenticationRepo) FindUserByUsername(username string) (*User, error) {
	var user User

	err := repo.db.Where("username=?", username).Find(&user).Error
	if err != nil {
		return &User{}, err
	}
	return &user, nil
}

func (repo *AuthenticationRepo) FindUserById(id uuid.UUID) (User, error) {
	var user User

	err := repo.db.Where("id=?", id).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// DI
func NewAuthenticationRepoModule(logger *zap.Logger, db *gorm.DB) *AuthenticationRepo {
	return &AuthenticationRepo{
		logger: logger,
		db:     db,
	}
}
