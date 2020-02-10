build:
	protoc -I. --go_out=plugins=micro:.  proto/vessel/vessel.proto
	docker build -t shippy-service-vessel .
run:
	docker run --net="host" --rm=true -e MICRO_REGISTRY=mdns -e MICRO_ADDRESS=":50051" -e DB_HOST="mongodb://localhost:27017" -p 50052:50051 shippy-service-vessel
