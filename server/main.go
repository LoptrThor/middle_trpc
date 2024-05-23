package main

import (
	pb "github.com/LoptrThor/middle_trpc/pb"
	"trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/log"
)

func main() {
	s := trpc.NewServer()
	pb.RegisterAccService(s.Service("trpc.app.server.Acc"), &accImpl{})
	if err := s.Serve(); err != nil {
		log.Fatal(err)
	}
}
