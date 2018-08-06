
build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/maddymanu/microservices-evan-tut/user-service \
			proto/user/user.proto
	docker build --no-cache -t user-service .

run:
	docker run --net="host" \
			-p 50053:50051 \
			-e DB_HOST=localhost \
			-e DB_NAME=postgres \
			-e DB_PASS=postgres \
			-e DB_PORT=5432 \
			-e DB_USER=postgres \
			-e MICRO_SERVER_ADDRESS=:50051 \
			-e MICRO_REGISTRY=mdns \
			user-service
