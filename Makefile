build-protoc:
	protoc -I. --micro_out=.  --go_out=.  proto/vessel/vessel.proto
build:
	docker build -t shippy-service-vessel .
run:
	docker run  --rm=true -e MICRO_REGISTRY=mdns -e MICRO_ADDRESS=":50051" \
 				-e DB_HOST="mongodb://mongo:27017" --link mongo \
			 	--name shippy-service-vessel \
			 	shippy-service-vessel
	#docker run --net="host" --rm=true -e MICRO_REGISTRY=mdns -e MICRO_ADDRESS=":50051" -e DB_HOST="mongodb://localhost:27017" -p 50052:50051 shippy-service-vessel
