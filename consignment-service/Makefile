
build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/maddymanu/microservices-evan-tut/consignment-service \
			proto/consignment/consignment.proto
	docker build --no-cache -t consignment-service .

run:
	docker run --net="host" \
			-p 50052 \
			-e MICRO_SERVER_ADDRESS=:50052 \
			-e MICRO_REGISTRY=mdns \
			consignment-service
