package utils

func HashPassword(password string) string {
	return "hashed_" + password
}

func CheckPasswordHash(password, hash string) bool {
	return hash == HashPassword(password)
}