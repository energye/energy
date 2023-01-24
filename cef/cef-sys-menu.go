package cef

type SysMenu struct {
	Label string
	Items []*SysMenuItem
}

type SysMenuItem struct {
	ChildMenu   *SysMenu
	IsSeparator bool
	Label       string
	Action      func()
	Disabled    bool
	Checked     bool
	Icon        []byte
	Shortcut    string
}

// NewMenu 创建一个新菜单，给定指定的标签和要显示的项目列表
func NewMenu(label string, items ...*SysMenuItem) *SysMenu {
	return &SysMenu{Label: label, Items: items}
}

// NewMenuItem 根据传递的标签和操作参数创建一个新菜单项
func NewMenuItem(label string, action func()) *SysMenuItem {
	return &SysMenuItem{Label: label, Action: action}
}

// NewMenuItemSeparator 创建将用作分隔符的菜单项
func NewMenuItemSeparator() *SysMenuItem {
	return &SysMenuItem{IsSeparator: true}
}
