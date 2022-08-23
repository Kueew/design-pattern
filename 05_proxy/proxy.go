package proxy

type IUser interface {
	Login(username, password string) error
}

type User struct {
}

func (u *User) Login(username, password string) error {
	//todo:登录
	return nil
}

type UserProxy struct {
	user *User
}

func NewUserProxy(u *User) *UserProxy {
	return &UserProxy{user: u}
}

func (p *UserProxy) Login(username, password string) error {
	//before 方法执行前的逻辑
	ret := p.user.Login(username, password)
	//after 方法执行后的逻辑
	return ret
}
