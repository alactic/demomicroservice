package connection

import (
	"fmt"

	"gopkg.in/couchbase/gocb.v1"
)

func Connection() *gocb.Bucket {
	var bucket *gocb.Bucket
	// cluster, _ := gocb.Connect("couchbase://localhost")
	cluster, _ := gocb.Connect("couchbase://172.17.0.2")
	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: "Administrator",
		Password: "password",
	})
	bucket, _ = cluster.OpenBucket("default", "")
	fmt.Println("host bucket:: ", bucket)

	return bucket
}
