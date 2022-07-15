package user

import (
	"errors"
	"sync"

	"github.com/example/project/app/addresses"
	"github.com/example/project/app/configs"
)

var lockUS = &sync.Mutex{}
var userService *UserService

func GetUserService() *UserService {
	lockUS.Lock()
	defer lockUS.Unlock()
	if userService == nil {
		userService = newUserService(configs.Cfg.Persistence)
	}
	return userService
}

type UserService struct {
	userOnline map[string]string
	ur         UserRepoIf
}

func newUserService(persistance bool) *UserService {
	return &UserService{make(map[string]string, 0), newUserRepo(persistance)}
}
func (s *UserService) logout(token string) {
	if _, ok := s.userOnline[token]; !ok {
		delete(s.userOnline, token)
	}
}
func (s *UserService) login(deviceId string, deviceName string, mobile string, pwd string) (*User, error) {
	//TODO: not check device info yet.
	user := s.findUser(mobile, pwd)
	if user == nil {
		return nil, errors.New("can not find user, which have mobile:" + mobile + "  pwd:" + pwd)
	}
	user.inflate()
	s.userOnline[mobile] = mobile //take moble as a token
	return user, nil
}
func (s *UserService) isOnline(token string) bool {
	_, ok := s.userOnline[token]
	return ok
}
func (s *UserService) FindUserByToken(token string) *User {
	mobileNumber := s.userOnline[token]
	if mobileNumber == "" {
		return nil
	}
	return s.ur.retriveUserByMobile(mobileNumber)
}

func (s *UserService) findUser(mobile string, pwd string) *User {
	user := s.ur.findUser(mobile, pwd)
	return user
}
func (s *UserService) RegisterNewUser(mobile string, pwd string, nickname string, autologin string) (*User, error) {
	//TODO: not check device info yet.
	if s.ur.findUser(mobile, pwd) != nil {
		return nil, errors.New("该手机号码已被占用！")
	}
	newUser, err := s.ur.CreateUser(mobile, pwd, nickname, autologin, genUId)
	if err != nil {
		return nil, errors.New("注册失败，内部错误。请重新尝试。")
	}
	s.userOnline[mobile] = mobile
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
