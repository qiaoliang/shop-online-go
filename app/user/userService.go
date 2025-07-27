package user

import (
	"errors"
	"sync"

	"bookstore/app/addresses"
)

var lockUS = &sync.Mutex{}
var userService *UserService

func GetUserService() *UserService {
	lockUS.Lock()
	defer lockUS.Unlock()
	if userService == nil {
		userService = newUserService(false)
	}
	return userService
}

type UserService struct {
	cache map[string]string
	ur    UserRepo
}

func NewUserServiceWithRepo(repo UserRepo) *UserService {
	return &UserService{make(map[string]string, 0), repo}
}

func newUserService(persistance bool) *UserService {
	return &UserService{make(map[string]string, 0), newUserRepo()}
}
func (s *UserService) logout(token string) {
	if _, ok := s.cache[token]; !ok {
		delete(s.cache, token)
	}
}
func (s *UserService) login(deviceId string, deviceName string, mobile string, pwd string) (*User, error) {
	//TODO: not check device info yet.
	user := s.findUser(mobile, pwd)
	if user == nil {
		return nil, errors.New("can not find user, which have mobile:" + mobile + "  pwd:" + pwd)
	}
	user.inflate()
	s.cache[mobile] = mobile //take moble as a token
	return user, nil
}
func (s *UserService) isOnline(token string) bool {
	_, ok := s.cache[token]
	return ok
}
func (s *UserService) FindUserByToken(token string) *User {
	//TODO: 现在的token 就是用户的手机号码。这里应该使用更好的TOKEN加密机制
	mobileNumber := token
	if mobileNumber == "" {
		return nil
	}
	return s.ur.RetriveUserByMobile(mobileNumber)
}

func (s *UserService) findUser(mobile string, pwd string) *User {
	user := s.ur.findUser(mobile, pwd)
	return user
}
func (s *UserService) RegisterNewUser(mobile string, pwd string, nickname string, autologin string) (*User, error) {
	//TODO: not check device info yet.
	if s.ur.findUser(mobile, pwd) != nil {
		return nil, errors.New("该用户已注册！")
	}
	newUser, err := s.ur.CreateUser(mobile, pwd, nickname, autologin, genUId)
	if err != nil {
		return nil, errors.New("注册失败，内部错误。请重新尝试。")
	}
	s.cache[mobile] = mobile
	return newUser, nil
}
func (s *UserService) GetDeliveryAddressesFor(token string) []addresses.Address {
	//TODO: Not implemented yet.
	return nil
}

func (s *UserService) GetDefaultDeliveryAddress(token string) []addresses.Address {
	//TODO: Not implemented yet.
	return nil
}

func (s *UserService) UpdateUserByToken(token string, data User) *User {
	// 通过token找到用户
	user := s.FindUserByToken(token)
	if user == nil {
		return nil
	}

	// 更新用户信息
	if data.Nickname != "" {
		user.Nickname = data.Nickname
	}
	if data.AvatarUrl != "" {
		user.AvatarUrl = data.AvatarUrl
	}
	if data.Province != "" {
		user.Province = data.Province
	}
	if data.City != "" {
		user.City = data.City
	}

	// 保存更新后的用户信息
	s.ur.updateUser(user)

	return user
}

