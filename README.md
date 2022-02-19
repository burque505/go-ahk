# go-ahk
## Home of go-ahk. Uses go-ole to wrap AutoHotkey.dll ##

*Windows only*. (Tested on Win10 Pro, but should work on Win7+.) 

Go version: 1.17.7

## First commit February 19, 2022

There will be support for most AutoHotkey.dll methods, hopefully. However, you can call all those methods anyway with e.g.
```
oleutil.MustCallMethod(ahk, "ahkexec", `var1 = 5 + 27.8`)
```
You will not *need* this package to access AutoHotkey.dll. You can do everything this package does with [go-ole](https://github.com/go-ole/go-ole). No files yet, sorry!

This package targets [AutoHotkey.dll v.1](https://github.com/HotKeyIt/ahkdll-v1-release/archive/master.zip).

### In the future I might upload a few gists of examples NOT using this package for those inclined to experiment. ###
