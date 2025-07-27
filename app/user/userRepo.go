package user

import (
	"errors"
	"fmt"
	"strconv"

	"bookstore/app/utils"

	"gorm.io/gorm"
)

// UserRepo 接口
//go:generate mockgen -source=userRepo.go -destination=mock_userRepo.go -package=user
// 方便后续 mock
//
type UserRepo interface {
	TotalUsers() int
	DeleteByMobile(mobile string)
	findUser(mobile, pwd string) *User
	RetriveUserByMobile(mobile string) *User
	RetriveUserByID(userID string) *User
	CreateUser(mobile, pwd, nickname, autologin string, genUserId UserIdGen) (*User, error)
	CreateAdmin(mobile, pwd string)
	updateUser(user *User) // 添加更新用户信息的方法
}

type UserIdGen func() string

func genUId() string {
	return fmt.Sprintf("userId%v", utils.RandomImpl{}.GenStr())
}

// UserRepoDB 实现

// UserRepoDB 用于数据库实现
//
type UserRepoDB struct {
	db *gorm.DB
}

func NewUserRepoDB(db *gorm.DB) *UserRepoDB {
	return &UserRepoDB{db: db}
}

func (r *UserRepoDB) TotalUsers() int {
	var count int64
	r.db.Model(&User{}).Count(&count)
	return int(count)
}

func (r *UserRepoDB) DeleteByMobile(mobile string) {
	r.db.Where("mobile = ?", mobile).Delete(&User{})
}

func (r *UserRepoDB) findUser(mobile, pwd string) *User {
	var user User
	if err := r.db.Where("mobile = ? AND pwd = ?", mobile, pwd).First(&user).Error; err != nil {
		return nil
	}
	return &user
}

func (r *UserRepoDB) RetriveUserByMobile(mobile string) *User {
	var user User
	if err := r.db.Where("mobile = ?", mobile).First(&user).Error; err != nil {
		return nil
	}
	return &user
}

func (r *UserRepoDB) RetriveUserByID(userID string) *User {
	var user User
	if err := r.db.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil
	}
	return &user
}

func (r *UserRepoDB) CreateUser(mobile, pwd, nickname, autologin string, genUserId UserIdGen) (*User, error) {
	// 验证手机号不能为空
	if mobile == "" {
		return nil, errors.New("手机号不能为空")
	}

	if r.RetriveUserByMobile(mobile) != nil {
		return nil, errors.New("hello,error")
	}
	al, _ := strconv.Atoi(autologin)
	userId := genUId()
	avatarUrl := "default_avatar.jpeg"
	if mock, ok := interface{}(utils.RandomMock{}).(interface{ GenAavatarStr() string }); ok {
		if v := mock.GenAavatarStr(); v != "" {
			avatarUrl = v
		}
	}
	user := &User{
		Id:          userId,
		Pwd:         pwd,
		Mobile:      mobile,
		Nickname:    nickname,
		AvatarUrl:   avatarUrl,
		Province:    "未知",
		City:        "未知",
		AutoLogin:   uint(al),
		UserInfo:    "FakeUserInfo",
		UserLevelId: GREENTYPE,
		UserLevel:   &UserLevel{GREENTYPE, GREENTYPE.String()},
	}
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepoDB) CreateAdmin(mobile, pwd string) {
	r.CreateUser(mobile, pwd, "超级塞亚人", "1", genUId)
}

func (r *UserRepoDB) updateUser(user *User) {
	if user == nil || user.Mobile == "" {
		return
	}
	// 在数据库实现中，使用GORM更新用户信息
	r.db.Save(user)
}

