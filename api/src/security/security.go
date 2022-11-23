package security

import "golang.org/x/crypto/bcrypt"

func Hash(password string) ([]byte, error){
    return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, stringPassword string) error {
    return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(stringPassword))
}