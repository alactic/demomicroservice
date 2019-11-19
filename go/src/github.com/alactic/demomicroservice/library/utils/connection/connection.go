package connection

import (
	"fmt"

	"gopkg.in/couchbase/gocb.v1"
)

func Connection() *gocb.Bucket {
	var bucket *gocb.Bucket
	// cluster, _ := gocb.Connect("couchbase://localhost")
	cluster, _ := gocb.Connect("http://192.168.0.100:31542")
	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: "elvis",
		Password: "password",
	})
	bucket, _ = cluster.OpenBucket("default", "")
	fmt.Println("host bucket:: ", bucket)

	return bucket
}
