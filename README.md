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

As this is still pre-alpha code, it's probably better to rely on the engine's built-in methods rather than the methods in 'engine.go'. Here's an example of a GUI
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
