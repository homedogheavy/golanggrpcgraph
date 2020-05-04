# go
```
cd go
go run main.go
```

# js
```
cd js
yarn serve
```

# protofuniki
```
brew install protobuf
protoc --go_out=./ proto/task_service.proto
protoc --plugin=protoc-gen-go=path/protoc-gen-go --go_out=plugins=grpc:../src/proto ./task_service.proto
protoc hoge.proto --js_out=import_style=commonjs:. --grpc-web_out=import_style=commonjs,mode=grpcwebtext:.
sudo mv ~/Downloads/protoc-gen-grpc-web-1.0.4-darwin-x86_64 /usr/local/bin/protoc-gen-grpc-web
yarn add google-protobuf grpc-web
yarn add -D babel-plugin-add-module-exports
```

