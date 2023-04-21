//source: https://doc.qt.io/qt-5/qtwebview-minibrowser-example.html

package main

import (
	"os"

	"github.com/ilius/qt/core"
	"github.com/ilius/qt/gui"
	"github.com/ilius/qt/qml"
	"github.com/ilius/qt/webview"
)

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	gui.NewQGuiApplication(len(os.Args), os.Args)
	webview.QtWebView_Initialize()

	var view = qml.NewQQmlApplicationEngine(nil)
	view.Load(core.NewQUrl3("qrc:/main.qml", 0))

	gui.QGuiApplication_Exec()
}
