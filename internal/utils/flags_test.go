package utils

import (
	"testing"

	"github.com/ilius/is/v2"
)

func TestAppendToFlag(t *testing.T) {
	is := is.New(t)

	flags := []string{"-p", "-pkgdir=/usr/lib/testdir", "-mod", "-modfile go.mod"}

	flags = AppendToFlag(flags, "-extldflags", "-v")
	is.AddMsg("Flag was not inserted").Equal("-extldflags=-v", flags[4])

	flags = AppendToFlag(flags, "-extldflags", "-Wl,soname,libname.so")
	is.AddMsg("Flag was not appended with new content").Equal("-extldflags=-v -Wl,soname,libname.so", flags[4])

	flags = AppendToFlag(flags, "-p", "test")
	is.AddMsg("Content was not appended to simple flag").Equal("-p=test", flags[0])
}
