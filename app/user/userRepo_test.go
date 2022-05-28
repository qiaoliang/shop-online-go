package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_create_user(t *testing.T) {
	ur := GetUserRepoInstance()
	assert.Equal(t, 0, ur.TotalUsers(), "should has 0 user")
	ur.CreateUser("mobile", "pwd", "nickname")
	assert.Equal(t, 1, ur.TotalUsers(), "should has 1 user")
	ur.CreateUser("mobile", "pwd", "nickname")
	assert.Equal(t, 2, ur.TotalUsers(), "should has 2 users")
}
