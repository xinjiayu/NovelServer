#!/bin/sh

BuildVersion=`git describe --abbrev=0 --tags`
BuildTime=`date +%FT%T%z`
CommitID=`git rev-parse HEAD`


function help() {
    echo "$0 linux|windows|mac"
}

function linux(){
    CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC=x86_64-linux-musl-gcc CGO_LDFLAGS="-static" go build -a -ldflags "-w -s -X main.BuildVersion=${BuildVersion} -X main.CommitID=${CommitID} -X main.BuildTime=${BuildTime}"

    copyFile
    cp curl.sh bin/
    cp NovelServer bin/

    rm -f NovelServer

}
function windows(){
    CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build -ldflags "-w -s -X main.BuildVersion=${BuildVersion} -X main.CommitID=${CommitID} -X main.BuildTime=${BuildTime}"

    copyFile
    cp NovelServer.exe bin/

    rm -f NovelServer.exe

}
function mac(){
    go build -ldflags "-w -s -X main.BuildVersion=${BuildVersion} -X main.CommitID=${CommitID} -X main.BuildTime=${BuildTime}"

    copyFile
    cp curl.sh bin/
    cp NovelServer bin/

    rm -f NovelServer

}

function copyFile() {
    rm -rf bin
    mkdir bin
    cp -r document/. bin/document/
    cp -r public/. bin/public/
    cp -r config/. bin/config/
}

if [ "$1" == "" ]; then
    help
elif [ "$1" == "linux" ];then
    linux
elif [ "$1" == "windows" ];then
    windows
elif [ "$1" == "mac" ];then
    mac
fi