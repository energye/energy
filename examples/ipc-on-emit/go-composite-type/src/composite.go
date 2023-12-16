package src

type DemoUser struct {
	Name     string
	Age      int
	Income   float64
	Sex      bool
	UserInfo DemoUserInfo
}

type DemoUserInfo struct {
	Phone       string
	Addr        string
	HeadPicture string
	Height      int
	Weight      int
}
