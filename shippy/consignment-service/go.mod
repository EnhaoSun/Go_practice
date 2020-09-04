module Go_practice/shippy/consignment-service

go 1.15

require google.golang.org/grpc v1.31.1

require (
	google.golang.org/protobuf v1.25.0 // indirect
	shippy/consignment-service/proto/consignment v1.2.3
)

replace shippy/consignment-service/proto/consignment => ./proto/consignment
