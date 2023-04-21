package top

import (
	"github.com/ilius/qt/quick"

	_ "github.com/ilius/qt/internal/examples/showcases/wallet/view/top/controller"
)

func init() { colorTemplate_QmlRegisterType2("TopTemplate", 1, 0, "ColorTemplate") }

type colorTemplate struct {
	quick.QQuickItem

	_ func() `signal:"change,->(controller.NewColorController(nil))"`
}
