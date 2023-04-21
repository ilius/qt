// +build !qml

package main

import (
	"os"

	"github.com/ilius/qt/core"
	"github.com/ilius/qt/widgets"

	"github.com/ilius/qt/internal/examples/sql/masterdetail_qml/controller"

	"github.com/ilius/qt/internal/examples/sql/masterdetail_qml/view"
)

func main() {
	qApp := widgets.NewQApplication(len(os.Args), os.Args)

	controller.NewController(nil).InitWith(core.NewQFile2(":/albumdetails.xml"), qApp)

	view.NewViewController(nil, 0).Show()

	qApp.Exec()
}
