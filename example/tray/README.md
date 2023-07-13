Energy 提供了四种系统托盘使用, 为了更多使用方式吧

1. [lclceftray](lclceftray) lcl + cef, 使用lcl托盘 + CEF窗口, 效果是 html 实现
2. [lcltray](lcltray) 纯 lcl, 效果纯菜单形式
3. [lclvftray](lclvftray) lcl + vf, 使用lcl托盘 + VF窗口, 效果是 html 实现
4. [systray](systray) 原生系统托盘, 不如LCL好, 实际是为linux使用, 效果纯菜单形式
- > 因为linux cef106版本以后默认使用gtk3，但是gtk3在energy里还不完全支持,
  > 并且gtk2也无法输入中文, 以后完整支持gtk3后将移除该系统原生托盘组件,
  > 老版本cef可以直接使用gtk2 支持lcl, 但中文问题。。
