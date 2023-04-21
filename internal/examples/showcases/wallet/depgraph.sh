#!/bin/bash

set -ev

$GOPATH/bin/godepgraph -horizontal -s -o github.com/ilius/qt/internal/examples/showcases/wallet github.com/ilius/qt/internal/examples/showcases/wallet | dot -Tpng -o depgraph.png