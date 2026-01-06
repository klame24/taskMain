package password

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes := []byte(password)

	hash, err := bcrypt.GenerateFromPassword(bytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
