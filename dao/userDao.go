package dao

import (
	"gorm.io/gorm"
	"sync"
)

var (
	userDao  *UserDAO
	userOnce sync.Once
)

type UserDAO struct {
}

// User 用户表
type User struct {
	gorm.Model
	FollowCount   int64  `json:"follow_count,omitempty"`                // 关注总数
	FollowerCount int64  `json:"follower_coun,omitemptyt"`              // 粉丝总数
	ID            int64  `json:"id,omitempty"`                          // 用户id
	Salt          []byte `gorm:"not null; type:varbinary(32)" json:"-"` // 密码盐
	Name          string `json:"name,omitempty"`                        // 用户名称
	Password      []byte `json:"password"`                              // 用户密码
}

func NewUserInfoDAO() *UserDAO {
	userOnce.Do(func() {
		userDao = new(UserDAO)
	})
	return userDao
}

// QueryUserById 根据id查询用户信息
func (u *UserDAO) QueryUserById(id int64) (*User, error) {
	user := &User{}
	err := DB.First(user, "Id = ?", id).Error
	return user, err
}

// AddUser 添加用户
func (u *UserDAO) AddUser(user *User) int {
	if u.userCheck(user) {
		return 1
	}
	if err := DB.Create(user).Error; err != nil {
		return 2
	}
	return 0
}

// QueryUserByName 根据用户名查询是否存在
func (u *UserDAO) QueryUserByName(name string) (*User, bool) {
	var user User
	err := DB.Where(&User{Name: name}, "name").First(&user).Error
	if err != nil {
		return nil, false
	}
	return &user, true
}

// userCheck 查询用户名是否重复
func (u *UserDAO) userCheck(user *User) bool {
	return DB.Model(&User{}).Where("name =?", user.Name).First(&User{}).Error == nil
}
