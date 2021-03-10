module micro-service-shippy/shippy-cli-consignment

go 1.15

require (
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/micro/micro/v3 v3.1.2-0.20210310091306-ef1881f4e1c7
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110 // indirect
	golang.org/x/sys v0.0.0-20210309074719-68d13333faf2 // indirect
	golang.org/x/text v0.3.5 // indirect
	google.golang.org/genproto v0.0.0-20210309190941-1aeedc14537d // indirect
	google.golang.org/grpc v1.36.0
	micro-service-shippy/shippy-service-consignment v0.1.0
)

replace micro-service-shippy/shippy-service-consignment => ../shippy-service-consignment
