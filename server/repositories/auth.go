package repositories

import (
	"cosmic-gate-chat/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// Password Encryption
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// Check whether the password provided is correct
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func AuthUser(userRequest models.User) (*models.User, error) {
	foundUser, err := GetUserByEmail(userRequest.Email)
	if foundUser.Email == "" || err != nil {
		return &models.User{}, errors.New("User Not Found")
	}

	ok := CheckPasswordHash(userRequest.Password, foundUser.Password)
	if !ok {
		return &models.User{}, errors.New("Wrong Password")
	}

	userData := &models.User{
		ID:             foundUser.ID,
		Email:          foundUser.Email,
		Username:       foundUser.Username,
		FriendRequests: foundUser.FriendRequests,
	}
	return userData, nil
}
