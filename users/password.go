package users

// HashPwd hashes a password.
func HashPwd(password string) (string, error) {
	//TODO no hash password
	//bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	//return string(bytes), err
	return password,nil
}

// CheckPwd checks if a password is correct.
func CheckPwd(password, hash string) bool {
	//TODO compare source
	//err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	//return err == nil
	return password == hash
}
