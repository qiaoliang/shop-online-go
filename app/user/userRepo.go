package user

import (
	"bookstore/app/configs"
	"bookstore/app/utils"
	"errors"
	"fmt"
	"strconv"
	"sync"
)

var lockUR = &sync.Mutex{}

type UserRepoMem struct {
	userlist map[string]*User
}

var userRepo UserRepoIf

func GetUserRepo() UserRepoIf {
	return newUserRepo(configs.Cfg.Persistence)
}
func newUserRepo(persistence bool) UserRepoIf {
	lockUR.Lock()
	defer lockUR.Unlock()

	if userRepo == nil {
		if persistence {
			userRepo = GetUserRepoDB(configs.Cfg.DBConnection())

		} else {

			userRepo = &UserRepoMem{make(map[string]*User, 10)}
			userRepo.CreateAdmin("13900007997", "1234")
		}
	}
	return userRepo
}

func GetMemoryUserRepo() UserRepoIf {
	lockUR.Lock()
	defer lockUR.Unlock()
	if userRepo == nil {
		userRepo = &UserRepoMem{make(map[string]*User, 10)}
		userRepo.CreateAdmin("13900007997", "1234")
	}
	return userRepo
}

func (r *UserRepoMem) TotalUsers() int {
	return len(r.userlist)
}

func (r *UserRepoMem) DeleteByMobile(mobile string) {
	//TODO: 未实现
}

func (r *UserRepoMem) findUser(mobile string, pwd string) *User {
	found := r.retriveUserByMobile(mobile)
	if found == nil || found.Password != pwd {
		return nil
	}
	return found
}
func (r *UserRepoMem) retriveUserByMobile(mobile string) *User {
	return r.userlist[mobile]
}

type UserIdGen func() string

func genUId() string {
	return fmt.Sprintf("userId%v", utils.RandomImpl{}.GenStr())
}

func (r *UserRepoMem) CreateUser(mobile string, pwd string, nickname string, autologin string, genUserId UserIdGen) (user *User, err error) {
	if r.findUser(mobile, pwd) != nil {
		return nil, errors.New("hello,error")
	}
	al, _ := strconv.Atoi(autologin)
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
		AutoLogin:   uint(al),
		UserInfo:    "FakeUserInfo",
		UserLevelId: GREENTYPE,
		UserLevel:   &UserLevel{GREENTYPE, GREENTYPE.String()},
	}
	return r.userlist[mobile], nil
}
func (r *UserRepoMem) CreateAdmin(mobile string, pwd string) {
	r.CreateUser(mobile, pwd, "超级塞亚人", "1", genUId)
}
