package dialog

import (
	"github.com/ilius/qt/core"

	_ "github.com/ilius/qt/internal/examples/showcases/wallet/wallet/dialog/controller"
)

func init() { sendTemplate_QmlRegisterType2("DialogTemplate", 1, 0, "SendTemplate") }

type sendTemplate struct {
	dialogTemplate

	_ func(string, string) *core.QVariant `slot:"send,->(controller.Controller)"`
}
