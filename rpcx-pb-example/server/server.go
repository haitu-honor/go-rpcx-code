package main

import (
	"context"
	"flag"
	"fmt"

	"gorpcx-model/pb"

	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

type Arith int

// 接受3个参数，第一个是 context.Context类型，其他2个都是可导出（或内置）的类型。
func (t *Arith) Mul(ctx context.Context, args *pb.ProtoArgs, reply *pb.ProtoReply) error {
	reply.C = args.A * args.B
	fmt.Printf("call: %d * %d = %d\n", args.A, args.B, reply.C)
	return nil
}

func main() {
	// flag.Parse()
	// 注册一个服务命为 Arith 的服务
	s := server.NewServer()
	//s.RegisterName("Arith", new(example.Arith), "")
	s.Register(new(Arith), "")
	s.Serve("tcp", *addr)
}
