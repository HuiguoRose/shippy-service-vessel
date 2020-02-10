module github.com/HuiguoRose/shippy-service-vessel

go 1.13

require (
	github.com/golang/protobuf v1.3.3
	github.com/micro/go-micro/v2 v2.0.0
	go.mongodb.org/mongo-driver v1.3.0
	golang.org/x/net v0.0.0-20200202094626-16171245cfb2 // indirect
)

replace sigs.k8s.io/yaml => github.com/kubernetes-sigs/yaml v1.1.0
