package tools
import (
  "golang.org/x/crypto/bcrypt"
)

// Takes a plaintext password and returns the bcrypt hash.
func HashPassword(password string) (string, error) {
  hashBytes, err := bcrypt.GenerateFromPassword(
    []byte(password),
    bcrypt.DefaultCost, 
  )
  return string(hashBytes), err
}

// Compares a bcrypt hash with its possible plaintext equivalent.
func CheckPassword(hash, password string) bool {
  err := bcrypt.CompareHashAndPassword(
    []byte(hash),
    []byte(password),
  )
  return err == nil
}
