# go-ahk
## Home of go-ahk. Uses go-ole to wrap AutoHotkey.dll ##

*Windows only*. (Tested on Win10 Pro, but should work on Win7+.) 

Go version: 1.17.7

## First commit February 19, 2022
IMPORTANT: To use this module you need AutoHotkey.dll. See the link below.
The DLL needs to be registered for COM on your system. 
```
regsvr32.exe <path-to-dll>\AutoHotkey.dll
```
You should only need to do this once unless you un-register it, e.g.
```
regsvr32.exe AutoHotkey.dll
```
There will be support for most AutoHotkey.dll methods, hopefully. However, you can call all those methods anyway with e.g.
```
oleutil.MustCallMethod(ahk, "ahkexec", `var1 = 5 + 27.8`)
```
(You don't necessarily *need* this package to access AutoHotkey.dll from Go. You can do everything this package does with [go-ole](https://github.com/go-ole/go-ole) and a little more typing. )

This module uses [AutoHotkey.dll v.1](https://github.com/HotKeyIt/ahkdll-v1-release/archive/master.zip).

### In the future I might upload a few gists of examples NOT using this package for those inclined to experiment. ###

_As this is still pre-alpha code, it might be better to call the engine's built-in methods rather than the wrapping methods in 'engine.go'._

Here's an example of a GUI
that requires engine.EngineReady(engine), with some commented-out code that accomplishes the same thing. If you uncomment the variables at the top and the for loop, be sure to comment out 'engine.EngineReady(ahk2)).
### GUI example: ###
```
package main

import (
	"fmt"

	"github.com/burque505/go-ahk/engine"
	"github.com/go-ole/go-ole/oleutil"
)

/*
var ahkTrue ole.VARIANT = ole.VARIANT{VT: ole.VT_I4, Val: 1}
var ahkFalse ole.VARIANT = ole.VARIANT{VT: ole.VT_I4, Val: 0}
*/

func main() {
	ahk2 := engine.CreateEngine()
	// only call if you'll be using a GUI! You have to be very careful with this.
	// If you

	script := `global labelText
	Gui, Add, Text, vlabelText w80 h30, Change me
	Gui, Add, Button, vButton1 Default, Change
	Gui, Show, w250 h150, Change
	
	return
	
	GuiClose:
	Esc::
	ExitApp
	
	ButtonChange:
		Gui, submit, nohide
		GuiControl, , labelText, Changed!
		Gui, submit, nohide
		return
	`
	fmt.Println(script)
	// engine.EngineReady(ahk)
	oleutil.MustCallMethod(ahk2, "ahkready")
	oleutil.MustCallMethod(ahk2, "ahktextdll", script)

	engine.EngineReady(ahk2) // This basically implements the code below
	/*
		for {
			time.Sleep(100 * time.Millisecond)
			// Now check again to see what "ahkReady" returns
			ready := oleutil.MustCallMethod(ahk2, "ahkReady")

			if *ready == ahkFalse {
				//fmt.Printf("%v\n", ready.Val)
				break
			}
		}
	*/

	oleutil.MustCallMethod(ahk2, "ahkTerminate")

}
```

#### If you get in an infinite loop, it's very likely you used *engine.EngineReady(engine)* or the commented-out 'for' loop where it isn't required. If your code execution never reaches a point where the engine isn't ready and you've made this method call, it'll loop forever. If this happens use Task Manager to kill the process.<br><br>If your code is in a file 'example.go' and you do get in an infinite loop, look in Task Manager for 'example.exe'. You may even have to stop 'go.exe'. (As you're experimenting with this it's almost certain to happen to you at least once.)
