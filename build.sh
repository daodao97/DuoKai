#!/bin/bash

gobuild="go build"
project="DuoKai"
execute="duokai"

function buildMac() {
  name=$1$([ -n "$2" ] && echo -$2 || echo )
  echo "start build mac-${name}"
 
  rm -rf build/$project.app
  cp -rf build/meta/$project.app build
  
  $(env GOOS=darwin GOARCH=$1 $([ -n "$2" ] && echo GOAMD64=$2 || echo ) CGO_ENABLED=1 $gobuild -o build/$project.app/Contents/MacOS/$execute .)

  (cd build && zip -r $project-${name}.zip $project.app 1>/dev/null)

  rm -rf build/$project.app
  echo "success !"
}

buildMac amd64
buildMac arm64

open build