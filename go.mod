module github.com/ilius/qt

go 1.20

require (
	github.com/gopherjs/gopherjs v1.17.2
	github.com/ilius/is/v2 v2.3.2
	github.com/sirupsen/logrus v1.9.0
	github.com/therecipe/env_darwin_amd64_513 v0.0.0-20190626001412-d8e92e8db4d0
	github.com/therecipe/env_linux_amd64_513 v0.0.0-20190626000307-e137a3934da6
	github.com/therecipe/env_windows_amd64_513 v0.0.0-20190626000028-79ec8bd06fb2
	golang.org/x/crypto v0.8.0
	golang.org/x/tools v0.8.0
)

require (
	github.com/therecipe/env_windows_amd64_513/Tools v0.0.0-20190626000028-79ec8bd06fb2 // indirect
	golang.org/x/sys v0.7.0 // indirect
)

replace golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190418165655-df01cb2cc480

replace golang.org/x/net => github.com/golang/net v0.0.0-20190420063019-afa5a82059c6

replace golang.org/x/sys => github.com/golang/sys v0.0.0-20190419153524-e8e3143a4f4a

replace golang.org/x/tools => github.com/golang/tools v0.0.0-20190420181800-aa740d480789

replace golang.org/x/text => github.com/golang/text v0.3.1-0.20190410012825-f4905fbd45b6
