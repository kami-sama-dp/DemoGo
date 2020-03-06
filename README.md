# DemoGo

# 使用grpc 生成proto

    go get -u github.com/golang/protobuf/protoc-gen-go
    cp protoc-gen-go  /usr/local/bin/
    vim ~/.bash_profile
    export GOPATH=$HOME/go
    PATH=$PATH:$GOPATH/bin
    protoc --go_out=.  ./proto/*.proto

