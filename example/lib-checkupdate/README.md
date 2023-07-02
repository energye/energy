### Energy liblcl auto update example

#### Demonstrating the liblcl dynamic link library module
* liblcl

#### Using the `autoupdate` tool function in the `cmd` package, check update.
#### This example demonstrates the use of UI window design
#### Reference Example Customization Update

### Check update steps
> 1. We copy a dynamic library and customize the import of dynamic link library functions to avoid conflicts with the currently updated ones
> 2. Injecting dynamic library proc by oneself `imports.SetEnergyImportDefs(version)`
> 3. initialize `inits.Init(nil, &resources)`
> 4. i18n config
> 5. open autoupdate `autoupdate.IsCheckUpdate(true)`
> 6. Set the callback function, and if there is an update, the function will be executed `autoupdate.CanUpdateLiblcl`
>> 6.1 `CanUpdateLiblcl` Obtain information on the updated module and the updated version level
> 
>> 6.2. Update processing here, coding...
> 7. call `autoupdate.CheckUpdate()` Check for updates