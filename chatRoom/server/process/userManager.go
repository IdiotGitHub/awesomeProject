package process

import (
	"fmt"
)

var (
	userManager *UserManager
)

type UserManager struct {
	onlineUser map[string]*UserProcessor
}

//init userManager
func init() {
	userManager = &UserManager{
		onlineUser: make(map[string]*UserProcessor, 1024),
	}
}

//add online user
func (u *UserManager) AddOnlineUser(processor *UserProcessor) {
	u.onlineUser[processor.UserId] = processor
}

//delete online user
func (u *UserManager) DeleteOnlineUser(userId string) {
	delete(u.onlineUser, userId)
}

//get online user
func (u *UserManager) GetOnlineUser() map[string]*UserProcessor {
	return u.onlineUser
}

//modify online user
func (u *UserManager) ModifyOnlineUser(processor *UserProcessor) {
	u.onlineUser[processor.UserId] = processor
}

//get online user by id
func (u *UserManager) GetOnlineUserById(userId string) (up *UserProcessor, err error) {
	up, ok := u.onlineUser[userId]
	if !ok {
		err = fmt.Errorf("user %s not found", userId)
		return
	}
	return
}
