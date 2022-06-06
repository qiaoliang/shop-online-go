package user

import (
	"bookstore/app/addresses"
	"sync"
)

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
	if us.userOnline[token] != "" {
		delete(us.userOnline, token)
	}
}
func (us *UserService) login(deviceId string, deviceName string, mobile string, pwd string) *User {
	user := us.findUser(mobile, pwd)
	if user == nil {
		return nil
	}
	us.userOnline[mobile] = mobile //take moble as a token
	return user
}
func (r *UserService) isOnline(token string) bool {
	mobileNumber := r.userOnline[token]
	return mobileNumber != ""
}
func (r *UserService) FindUserByToken(token string) *User {
	mobileNumber := r.userOnline[token]
	if mobileNumber == "" {
		return nil
	}
	return GetUserRepoInstance().retriveUserByMobile(mobileNumber)
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
	s.userOnline[mobile] = mobile
	return newUser
}
func (s *UserService) GetDeliveryAddressesFor(token string) []addresses.Address {
	return nil
}

func (s *UserService) GetDefaultDeliveryAddress(token string) []addresses.Address {
	return nil
}
