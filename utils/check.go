package utils

const maxUserPassSaltLen = 32

// CheckName 检查用户名
func CheckName(username string) bool {
	return len(username) != 0 && len(username) <= maxUserPassSaltLen
}

// CheckPass 检查密码
func CheckPass(pass string) bool {
	return len(pass) != 0 && len(pass) <= maxUserPassSaltLen
}
