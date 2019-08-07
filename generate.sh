protoc Greet/greetpb/greet.proto --go_out=plugins=grpc:.
go run Greet/greet_server/server.go
go run Greet/greet_client/client.go

protoc Calculator/calpb/calcu.proto --go_out=plugins=grpc:.
go run Calculator/cal_server/server.go
go run Calculator/cal_client/client.go