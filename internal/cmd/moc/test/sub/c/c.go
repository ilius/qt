package c

import (
	"github.com/ilius/qt/core"

	_ "github.com/ilius/qt/internal/cmd/moc/test/sub/conf"
)

type StructSubGoC struct{}
type StructSubMocC struct{ core.QObject }
