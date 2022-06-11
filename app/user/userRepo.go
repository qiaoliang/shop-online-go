package user

import (
	"bookstore/app/configs"
	"bookstore/app/utils"
	"errors"
	"fmt"
	"sync"
)

var lockUR = &sync.Mutex{}

type MemoryUserRepo struct {
	userlist map[string]*User
}

var userRepo UserRepoIf

func GetUserRepo() UserRepoIf {
	return NewUserRepo(configs.Cfg.Persistence)
}
func NewUserRepo(persistence bool) UserRepoIf {
	lockUR.Lock()
	defer lockUR.Unlock()

	if userRepo == nil {
		if persistence {
			userRepo = GetUserRepoDB(configs.Cfg.DBConnection())

		} else {

			userRepo = &MemoryUserRepo{make(map[string]*User, 10)}
			userRepo.CreateAdmin("13900007997", "1234")
		}
	}
	return userRepo
}

func GetMemoryUserRepo() UserRepoIf {
	lockUR.Lock()
	defer lockUR.Unlock()
	if userRepo == nil {
		userRepo = &MemoryUserRepo{make(map[string]*User, 10)}
		userRepo.CreateAdmin("13900007997", "1234")
	}
	return userRepo
}

func (r *MemoryUserRepo) TotalUsers() int {
	return len(r.userlist)
}

func (r *MemoryUserRepo) DeleteByMobile(mobile string) {
	//TODO: 未实现
}

func (r *MemoryUserRepo) findUser(mobile string, pwd string) *User {
	found := r.retriveUserByMobile(mobile)
	if found == nil || found.Password != pwd {
		return nil
	}
	return found
}
func (r *MemoryUserRepo) retriveUserByMobile(mobile string) *User {
	return r.userlist[mobile]
}

type UserIdGen func() string

func genUId() string {
	return fmt.Sprintf("userId%v", utils.NewRandom().GenStr())
}

func (r *MemoryUserRepo) CreateUser(mobile string, pwd string, nickname string, genUserId UserIdGen) (user *User, err error) {
	if r.findUser(mobile, pwd) != nil {
		return nil, errors.New("hello,error")
	}
	userId := genUserId()
	avatarUrl := utils.NewRandom().GenAavatarStr()
	r.userlist[mobile] = &User{
		Id:          userId,
		Password:    pwd,
		Mobile:      mobile,
		Nickname:    nickname,
		AvatarUrl:   avatarUrl,
		Province:    "未知",
		City:        "未知",
		AutoLogin:   0,
		UserInfo:    "FakeUserInfo",
		UserLevelId: 1,
		UserLevel:   &UserLevel{GREENTYPE, GREENTYPE.String()},
	}
	return r.userlist[mobile], nil
}
func (r *MemoryUserRepo) CreateAdmin(mobile string, pwd string) {
	r.CreateUser(mobile, pwd, "超级塞亚人", genUId)
}
