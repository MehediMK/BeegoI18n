<h1 align='center'>BeegoI18n project using goleng Beego Framework </h1>


## Prerequisites
  - Inatall go version 1.19
  - Beego version 2.0.4
 
## First create those directory 
- [x] `mkdir go` `cd go`
- [x] `mkdir bin src pkg`

```mermaid
graph TD;
    go-->bin;
    go-->src;
    go-->pkg;
```
  
## Setup GOPATH, GOBIN and GOENV variables:
> For linux:
 - `export GOPATH=/usr/local/go`
 - `export GOBIN=/usr/local/go/bin`
 - `export GOENV=/usr/local/go`
> For Windlow:
 - `set GOPATH=D:\go`
 - `set GOBIN=D:\go\bin`
 - `set GOENV=D:\go`
 
## Create mod file
  - (Outside GOPATH) run `go mod init go` and then `go mod tidy` on terminal 
  
## Install Beego
  - `go get github.com/astaxie/beego`
  - for windows `go install github.com/beego/bee/v2@latest;`
  
# Install dependencies:
  - `go get -u github.com/beego/beego/v2@latest`
  - `go get -u github.com/beego/bee/v2@latest`
  - `go get github.com/beego/beego/v2/server/web@v2.0.4`
  - `go get github.com/beego/i18n` for **i18n/beego language management**

## How to Run this Project

  - First of all clone this Project
  - Now go to the download directory
  - Extract this File
  - move this project `/usr/local/go/src` directory
  - go to project directory
  - run `bee run`
  - `http://localhost:8080` type on your browser

## Helpful commands
  - `go version` "for check go version"
  - `bee version` "for check bee version"
  - `bee new project-name` "For create new bee project"
