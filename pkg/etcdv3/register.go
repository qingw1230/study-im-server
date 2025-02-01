package etcdv3

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/qingw1230/study-im-server/pkg/common/log"
	clientv3 "go.etcd.io/etcd/client/v3"
)

const (
	TIME_TO_LIVE = 10
)

type RegEtcd struct {
	cli          *clientv3.Client
	endpoints    []string
	serviceValue string
	serviceKey   string
	ttl          int
	schema       string
	serviceName  string
	reChan       chan struct{} // 控制重连
	connected    bool
}

var rEtcd *RegEtcd

// "%s:///%s/"
func GetPrefix(schema, serviceName string) string {
	return fmt.Sprintf("%s:///%s/", schema, serviceName)
}

func RegisterEtcd(schema string, endpoints []string, myHost string, myPort int, serviceName string, ttl int) error {
	// schema:///serviceName/ip:port -> ip:port
	serviceValue := net.JoinHostPort(myHost, strconv.Itoa(myPort))
	serviceKey := GetPrefix(schema, serviceName) + serviceValue
	rEtcd = &RegEtcd{
		endpoints:    endpoints,
		serviceValue: serviceValue,
		serviceKey:   serviceKey,
		ttl:          ttl,
		schema:       schema,
		serviceName:  serviceName,
		reChan:       make(chan struct{}),
	}
	rEtcd.ReRegister()
	return rEtcd.Register()
}

func (r *RegEtcd) Register() (err error) {
	defer func() {
		if err != nil {
			r.reChan <- struct{}{}
		}
	}()

	var (
		cli    *clientv3.Client
		ctx    context.Context
		cancel context.CancelFunc
		resp   *clientv3.LeaseGrantResponse
		kresp  <-chan *clientv3.LeaseKeepAliveResponse
	)

	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   r.endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Error("etcd New failed", err.Error())
		return
	}

	ctx, cancel = context.WithCancel(context.Background())
	resp, err = cli.Grant(ctx, int64(r.ttl))
	if err != nil {
		log.Error("etcd Grant failed", err.Error())
		return
	}

	if _, err = cli.Put(ctx, r.serviceKey, r.serviceValue, clientv3.WithLease(resp.ID)); err != nil {
		log.Error("etcd Push key value failed", err.Error())
		return
	}

	kresp, err = cli.KeepAlive(ctx, resp.ID)
	if err != nil {
		log.Error("etcd KeepAlive failed", err.Error())
		return
	}
	rEtcd.cli = cli

	go func() {
		for {
			select {
			case _, ok := <-kresp:
				r.connected = ok
				if !ok {
					log.Error("etcd 租约失败")
					cancel()
					r.reChan <- struct{}{}
					return
				}
			}
		}
	}()
	return
}

func (r *RegEtcd) ReRegister() {
	go func() {
		for {
			select {
			case _, ok := <-r.reChan:
				if !ok {
					return
				}
				time.Sleep(1 * time.Second)
				r.Register()
			}
		}
	}()
}
