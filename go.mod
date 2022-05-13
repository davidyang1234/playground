module playground

go 1.13

require (
	cloud.google.com/go/pubsub v1.21.1
	github.com/gin-gonic/gin v1.7.7
)

// replace google.golang.org/grpc v1.45.0 => google.golang.org/grpc v1.26.0

replace github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.6
