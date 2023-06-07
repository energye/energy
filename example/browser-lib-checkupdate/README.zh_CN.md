### Energy liblcl 自动更新示例

#### 演示liblcl动态链接库模块
* liblcl

#### 使用`cmd`包中的`autoupdate`工具功能，检查更新。
#### 此示例使用了UI窗口设计
#### 参考该示例自定义更新, 或直接使用

### 检查更新步骤
> 1. 我们复制了一个动态库，并自定义了动态链接库函数的导入，以避免与当前更新的函数发生冲突
> 2. 自行导入动态库proc `imports.SetEnergyImportDefs(version)`
> 3. 初始化 `inits.Init(nil, &resources)`
> 4. i18n 配置
> 5. 打开自动更新 `autoupdate.IsCheckUpdate(true)`
> 6. 设置回调函数，如果有更新，函数将被执行 `autoupdate.CanUpdateLiblcl`
>> 6.1 `CanUpdateLiblcl` 回调函数返回更新模块信息和更新版本级别
> 
>> 6.2. 在这里编写更新代码, coding...
> 7. 调用 `autoupdate.CheckUpdate()` 检查更新