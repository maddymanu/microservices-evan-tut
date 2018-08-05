
build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/maddymanu/microservices-evan-tut/user-service \
			proto/user/user.proto
	docker build --no-cache -t user-service .

run:
	docker run -p 50053:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns user-service
