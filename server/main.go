package main

import (
	"context"
	pb "github.com/LoptrThor/middle_trpc/pb"
	"log"
	"trpc.group/trpc-go/trpc-go"
)

func main() {
	s := trpc.NewServer()
	pb.RegisterAccService(s.Service("trpc.app.server.Acc"), &accImpl{})
	if err := s.Serve(); err != nil {
		log.Fatal(err)
	}
}

type accImpl struct {
	pb.UnimplementedAcc
}

func (s *accImpl) OpenCheck(
	ctx context.Context,
	req *pb.Req,
) (*pb.Rsp, error) {

	log.Fatal(req)
	rsp := &pb.Rsp{}
	rsp.Rsp = "222"
	return rsp, nil
}
