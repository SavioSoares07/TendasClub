package database

func UserExists(email string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)"
	err := DB.QueryRow(query, email).Scan(&exists)
	return exists, err
}