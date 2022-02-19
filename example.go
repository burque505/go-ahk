package main

import (
	// "os"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/burque505/go-ahk/engine"

	// "path/filepath"
	"reflect"

	ole "github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
	// "io/ioutil"
)

// Required if you're going to use AHK booleans from Go.
// engine.go also defines ahkFalse at the moment.
// It may be removed from engine.go in the future
var ahkTrue ole.VARIANT = ole.VARIANT{VT: ole.VT_I4, Val: 1}
var ahkFalse ole.VARIANT = ole.VARIANT{VT: ole.VT_I4, Val: 0}

// These variables are purely for my own convenience
// Feel free to rename them
var Print = fmt.Println
var Call = oleutil.MustCallMethod
var Sleep = time.Sleep

func main() {

	ahk := engine.CreateEngine()
	defer ahk.Release()
	Print(reflect.TypeOf(ahk))

	// Later in the example we'll use functions from a local .ahk file.
	// (In your own code you may want to embed .ahk scripts.
	// See e.g. this SO post:
	// https://stackoverflow.com/questions/17796043/how-to-embed-files-into-go-binaries)
	content, err := ioutil.ReadFile(`functions.ahk`)
	script := string(content)

	if err != nil {
		log.Fatal(err)
	}

	// Call this to start ablank thread. GUIs require a different approach.
	// A GUI example will be added later.

	Call(ahk, "ahktextdll", "")

	time.Sleep(1 * time.Second)

	Call(ahk, "addScript", script)

	Sleep(1 * time.Second)
	var sayHelloFunction = `SayHello(name) {
		msgbox, ,Monkey Business, Hello %name%, 3
	}
	`
	Call(ahk, "ahkexec", sayHelloFunction)
	Sleep(100 * time.Millisecond)
	Call(ahk, "ahkexec", `SayHello("Mario") `)

	//m(ahk, "ahkExec", "Add5(3)")

	Call(ahk, "ahkexec", `msgbox, ,Function Tests, hello world from golang, 3`)

	/*
		// NOT NEEDED FOR THIS SCRIPT
		// You can call this with ahk.EngineReady(ahk) // or whatever you named the engine
		for {
			time.Sleep(100 * time.Millisecond)
			// Now check again to see what "ahkReady" returns
			ready := oleutil.MustCallMethod(ahk, "ahkReady")

			if *ready == ahkFalse {
				p("AHK engine is finished, exiting")
				//fmt.Printf("%v\n", ready.Val)
				break
			}

		}
		//You will need this if you run a script that has e.g. a GUI. Otherwise Go won't show the GUI
		//A similar 'wait' loop is required for ALL languages that call Autohotkey.dll
	*/

	Call(ahk, "ahkexec", "var1 = 5")
	Call(ahk, "ahkexec", "var2 = 10")
	Call(ahk, "ahkexec", `var5 := "Fat monkeys"`)
	Print("var2 as string is", oleutil.MustCallMethod(ahk, "ahkgetvar", "var2").Value())

	Call(ahk, "ahkexec", `var3 := var1+var2`)
	Sleep(1 * time.Second)
	Call(ahk, "ahkexec", `msgbox, ,Monkeys have succeeded in capturing variables, var3 = %var3%`)

	var3 := Call(ahk, "ahkgetvar", `var3`).Value()
	Print("var 3 is now", var3)

	Print("I am still around")
	// Terminate the AutoHotkey engine
	Call(ahk, "ahkTerminate")
	// ahk.Release() if not deferred above
	Print("Done")

}
