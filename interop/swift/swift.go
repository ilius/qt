package swift

import "github.com/ilius/qt/interop"

func init() {
	interop.ReturnPointersAsStrings = true
	interop.SupportsSyncCallsIntoRemote = false
}
