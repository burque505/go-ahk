package engine

import (
	"fmt"
	"log"
	"time"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)


func CreateEngine() *ole.IDispatch {

	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	unknown, err := oleutil.CreateObject("AutoHotkey.Script")
	defer unknown.Release()

	if err != nil {
		log.Fatal("CreateObject: ", err)
	}
	engine, err := unknown.QueryInterface(ole.IID_IDispatch)
	//defer engine.Release()

	if err != nil {
		log.Fatal("QueryInterface: ", err)
	}
	return engine
}

func EngineReady(engine *ole.IDispatch) {
	//var ahkTrue ole.VARIANT = ole.VARIANT{VT: ole.VT_I4, Val: 1}
	var ahkFalse ole.VARIANT = ole.VARIANT{VT: ole.VT_I4, Val: 0}
	for {
		time.Sleep(100 * time.Millisecond)
		// Now check again to see what "ahkReady" returns
		ready := oleutil.MustCallMethod(engine, "ahkReady")

		if *ready == ahkFalse {
			fmt.Println("AHK engine is finished, exiting")
			//fmt.Printf("%v\n", ready.Val)
			break
		}

	}
}


func HelloEngine() {
	fmt.Println("Hello from the engine")
}