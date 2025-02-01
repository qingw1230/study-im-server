package etcdv3

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/qingw1230/study-im-server/pkg/common/log"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

const (
	CONST_DURATION_GRPC_TIMEOUT_SECOND = 5 * time.Second
)

type Resolver struct {
	cc                 resolver.ClientConn
	schema             string
	serviceName        string
	grpcClientConn     *grpc.ClientConn
	cli                *clientv3.Client
	watchStartRevision int64
}

var (
	resolvers     = make(map[string]*Resolver)
	resolverMutex sync.RWMutex
)

func NewResolver(schema string, endpoints []string, serviceName string) (*Resolver, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: endpoints,
	})
	if err != nil {
		log.Error("etcd new failed", err.Error())
		return nil, err
	}

	r := Resolver{
		schema:      schema,
		serviceName: serviceName,
		cli:         cli,
	}
	resolverMutex.Lock()
	resolver.Register(&r)
	resolverMutex.Unlock()
	r.grpcClientConn, err = r.newGrpcClientConn()
	return &r, err
}

func (r *Resolver) newGrpcClientConn() (conn *grpc.ClientConn, err error) {
	log.Info("call newGrpcClientConn")
	var (
		opts []grpc.DialOption
		ctx  context.Context
	)
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)))
	ctx, _ = context.WithTimeout(context.Background(), CONST_DURATION_GRPC_TIMEOUT_SECOND)
	resolverMutex.Lock()
	conn, err = grpc.DialContext(ctx, GetPrefix(r.schema, r.serviceName), opts...)
	resolverMutex.Unlock()
	if err != nil {
		log.Error("DialContext failed", err.Error())
		return
	}
	log.Info("newGrpcClientConn return")
	return
}

func GetConn(schema string, endpoints []string, serviceName string) *grpc.ClientConn {
	log.Info("call GetConn")
	resolverMutex.RLock()
	r, ok := resolvers[schema+serviceName]
	resolverMutex.RUnlock()
	if ok {
		return r.grpcClientConn
	}

	r, err := NewResolver(schema, endpoints, serviceName)
	if err != nil {
		return nil
	}
	resolverMutex.Lock()
	resolvers[schema+serviceName] = r
	resolverMutex.Unlock()
	return r.grpcClientConn
}

func (r *Resolver) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	if r.cli == nil {
		return nil, fmt.Errorf("etcd clientv3 client failed, etcd: %v", target)
	}
	r.cc = cc
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	prefix := GetPrefix(r.schema, r.serviceName)
	resp, err := r.cli.Get(ctx, prefix, clientv3.WithPrefix())
	if err == nil {
		var addrList []resolver.Address
		for i := range resp.Kvs {
			addrList = append(addrList, resolver.Address{Addr: string(resp.Kvs[i].Value)})
		}
		r.cc.UpdateState(resolver.State{Addresses: addrList})
		r.watchStartRevision = resp.Header.Revision + 1
		go r.watch(prefix, addrList)
	} else {
		return nil, fmt.Errorf("etcd get failed, prefix: %s", prefix)
	}
	return r, nil
}

func (r *Resolver) watch(prefix string, addrList []resolver.Address) {
	rch := r.cli.Watch(context.Background(), prefix, clientv3.WithPrefix())
	for n := range rch {
		update := false
		for _, ev := range n.Events {
			switch ev.Type {
			case mvccpb.PUT:
				if !exists(addrList, string(ev.Kv.Value)) {
					update = true
					addrList = append(addrList, resolver.Address{Addr: string(ev.Kv.Value)})
				}
			case mvccpb.DELETE:
				i := strings.LastIndexAny(string(ev.Kv.Key), "/")
				if i < 0 {
					return
				}
				t := string(ev.Kv.Key)[i+1:]
				if s, ok := remove(addrList, t); ok {
					update = true
					addrList = s
				}
			}
		}

		if update {
			r.cc.UpdateState(resolver.State{Addresses: addrList})
		}
	}
}

func exists(addrList []resolver.Address, addr string) bool {
	for _, v := range addrList {
		if v.Addr == addr {
			return true
		}
	}
	return false
}

func remove(s []resolver.Address, addr string) ([]resolver.Address, bool) {
	for i := range s {
		if s[i].Addr == addr {
			s[i] = s[len(s)-1]
			return s[:len(s)-1], true
		}
	}
	return nil, false
}

func (r *Resolver) Scheme() string {
	return r.schema
}

func (r *Resolver) ResolveNow(rn resolver.ResolveNowOptions) {
}

func (r *Resolver) Close() {
}
