package main

import (
	"github.com/ilius/qt/core"
)

type Person struct {
	core.QObject

	_ string `property:"name"`
	_ int    `property:"shoeSize"`
}
