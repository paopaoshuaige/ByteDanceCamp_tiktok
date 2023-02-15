package dao

import "sync"

var (
	userDao  *UserDAO
	userOnce sync.Once
)

type UserDAO struct {
}

// User
type User struct {
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
