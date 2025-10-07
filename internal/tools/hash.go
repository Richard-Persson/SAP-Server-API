package tools
import (
  "golang.org/x/crypto/bcrypt"
)

// HashPassword takes a plaintext password and returns the bcrypt hash.
func HashPassword(password string) (string, error) {
  hashBytes, err := bcrypt.GenerateFromPassword(
    []byte(password),
    bcrypt.DefaultCost, // ~12 by default
  )
  return string(hashBytes), err
}

// CheckPassword compares a bcrypt hash with its possible plaintext equivalent.
// Returns true on match, false otherwise.
func CheckPassword(hash, password string) bool {
  err := bcrypt.CompareHashAndPassword(
    []byte(hash),
    []byte(password),
  )
  return err == nil
}
