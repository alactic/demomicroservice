package main

import (
	"fmt"
	"github.com/alactic/demomicroservice/library/utils/connection"
	"github.com/alactic/demomicroservice/library/utils/router"
	"gopkg.in/couchbase/gocb.v1"
)

var bucket *gocb.Bucket

func main() {
	fmt.Println("starting application")
	bucket = connection.Connection()
	router.Router()
}
