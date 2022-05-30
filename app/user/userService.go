package user

import "sync"

var lockUS = &sync.Mutex{}
var userService *UserService

func GetUserService() *UserService {
	lockUS.Lock()
	defer lockUS.Unlock()
	if userService == nil {
		userService = &UserService{}
		userService.userOnline = make(map[string]string, 10)
	}
	return userService
}

type UserService struct {
	userOnline map[string]string
}

func (us *UserService) logout(token string) {
	if us.userOnline["token"] != "" {
		delete(us.userOnline, token)
	}
}
func (us *UserService) login(deviceId string, deviceName string, mobile string, pwd string) *User {
	user := us.findUser(mobile, pwd)
	if user == nil {
		return nil
	}
	us.userOnline["token"] = mobile //should be token, rather than Mobile
	return user
}
func (s *UserService) findUser(mobile string, pwd string) *User {
	user := GetUserRepoInstance().findUser(mobile, pwd)
	return user
}
func (s *UserService) RegisterNewUser(mobile string, pwd string, nickname string) *User {
	if GetUserRepoInstance().findUser(mobile, pwd) != nil {
		return nil
	}

	newUser, err := GetUserRepoInstance().CreateUser(mobile, pwd, nickname)
	if err != nil {
		return nil
	}
	s.userOnline["token"] = mobile
	return newUser
}
