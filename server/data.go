package main

import (
	pb "github.com/kuuhaku86/simple-go-rpc/dataserver"
)

var Datas []pb.Data = []pb.Data{
	pb.Data{
		Key:   "a",
		Value: 10,
	},
	pb.Data{
		Key:   "b",
		Value: 100,
	},
	pb.Data{
		Key:   "c",
		Value: 1000,
	},
}