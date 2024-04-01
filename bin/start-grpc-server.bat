@echo off
echo Compiling Go code...
cd ..
go build -o bin/start-grpc-server.exe servers/grpcServer/main.go
echo Running Go code...
cd bin
start-grpc-server.exe