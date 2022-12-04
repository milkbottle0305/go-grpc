# Go Grpc - FSSN Project
This repository is final project for fullstack service networking.
## Overview
Four server-client source codes are included.
- <a href="https://github.com/milkbottle0305/go-grpc/tree/master/unary">unary</a>
- <a href="https://github.com/milkbottle0305/go-grpc/tree/master/bidirectional-streaming">bidirectional-streaming</a>
- <a href="https://github.com/milkbottle0305/go-grpc/tree/master/client-streaming">client-streaming</a>
- <a href="https://github.com/milkbottle0305/go-grpc/tree/master/server-streaming">server-streaming </a>

## How to Install
### Install Go language 

Windows
Install Go language in https://golang.org/

Linux
```
$ sudo apt update && sudo apt upgrade -y 
$ sudo apt install golang
```

Mac OS
```
$ brew install go
```

### Install go-grpc package
Install go-grpc, protocompiler
```
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```
Add environment variable for Go directory
```
$ export PATH="$PATH:$(go env GOPATH)/bin"
```

## How to Starting
1. Clone this repository
```
$ git clone https://github.com/milkbottle0305/go-grpc.git
$ cd go-grpc
```
2. Run server first
```
$ go run $(directory name)/server/server.go
```
3. Then client
```
$ go run $(directory name)/cleint/client.go
```

## Who made?
경희대학교 컴퓨터공학과 2019102200 유병우