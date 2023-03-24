package pkg

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/resolver/manual"
)

func NewGrpcClient(addrSlice []string) (*grpc.ClientConn, error) {
	addrs := make([]resolver.Address, len(addrSlice))
	for i, addr := range addrSlice {
		addrs[i].Addr = addr
	}
	builder := manual.NewBuilderWithScheme("prod")
	conn, err := grpc.Dial(
		builder.Scheme()+":///prod.server",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithResolvers(builder),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`))
	builder.UpdateState(resolver.State{Addresses: addrs})
	return conn, err
}
