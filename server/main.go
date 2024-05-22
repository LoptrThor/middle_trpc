package main

import (
	"context"
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

type accImpl struct {
	pb.UnimplementedAcc
}

func (s *accImpl) OpenCheck(
	ctx context.Context,
	req *pb.Req,
) (*pb.Rsp, error) {
	log.Info(req.Req)
	rsp := &pb.Rsp{}
	rsp.Rsp = "222"
	rsp.ErrCode = 0
	rsp.ErrMsg = "ok"
	return rsp, nil
}
