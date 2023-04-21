package haxe

import "github.com/ilius/qt/interop"

func init() {
	interop.ReturnPointersAsStrings = true
	interop.SupportsSyncCallsIntoRemote = false
}
