package main

import (
	"os"

	"github.com/ilius/qt/core"
	"github.com/ilius/qt/gui"
	"github.com/ilius/qt/qml"
	"github.com/ilius/qt/quickcontrols2"
)

func main() {
	os.Setenv("QT_IM_MODULE", "qtvirtualkeyboard")
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	gui.NewQGuiApplication(len(os.Args), os.Args)

	quickcontrols2.QQuickStyle_SetStyle("Material")
	engine := qml.NewQQmlApplicationEngine(nil)
	engine.Load(core.NewQUrl3("qrc:/qml/main.qml", 0))

	gui.QGuiApplication_Exec()
}
