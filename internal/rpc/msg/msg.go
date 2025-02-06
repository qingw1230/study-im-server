package msg

import (
	"net"
	"strconv"

	"github.com/qingw1230/study-im-server/pkg/common/config"
	"github.com/qingw1230/study-im-server/pkg/common/kafka"
	"github.com/qingw1230/study-im-server/pkg/common/log"
	"github.com/qingw1230/study-im-server/pkg/etcdv3"
	pbMsg "github.com/qingw1230/study-im-server/pkg/proto/msg"
	"github.com/qingw1230/study-im-server/pkg/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type msgServer struct {
	pbMsg.UnimplementedMsgServer
	rpcPort         int
	rpcRegisterName string
	etcdSchema      string
	etcdAddr        []string
	producer        *kafka.Producer
}

func NewMsgServer(port int) *msgServer {
	log.NewPrivateLog("msg")
	rpc := &msgServer{
		rpcPort:         port,
		rpcRegisterName: config.Config.RpcRegisterName.OfflineMessageName,
		etcdSchema:      config.Config.Etcd.EtcdSchema,
		etcdAddr:        config.Config.Etcd.EtcdAddr,
	}
	rpc.producer = kafka.NewKafkaProducer(config.Config.Kafka.OfflineMsgToMongoMysql.Addr, config.Config.Kafka.OfflineMsgToMongoMysql.Topic)
	return rpc
}

func (s *msgServer) Run() {
	log.Info("rpc msg start...")
	address := utils.ServerIP + ":" + strconv.Itoa(s.rpcPort)
	ln, err := net.Listen("tcp", address)
	if err != nil {
		log.Error("listen network failed ", err.Error(), address)
		return
	}
	defer ln.Close()

	server := grpc.NewServer()
	defer server.GracefulStop()

	pbMsg.RegisterMsgServer(server, s)
	reflection.Register(server)
	printRegisteredServices(server)

	err = etcdv3.RegisterEtcd(s.etcdSchema, s.etcdAddr, utils.ServerIP, s.rpcPort, s.rpcRegisterName, etcdv3.TIME_TO_LIVE)
	if err != nil {
		log.Error("msg RegisterEtcd failed", err.Error())
		return
	}
	log.Info("rpc account register success")
	err = server.Serve(ln)
	if err != nil {
		log.Error("Server failed ", err.Error())
		return
	}
}

func printRegisteredServices(srv *grpc.Server) {
	services := srv.GetServiceInfo()
	log.Info("已注册的 gRPC 服务:")
	for name, info := range services {
		log.Info("- 服务名称:", name)
		for _, method := range info.Methods {
			log.Info("  ⮡ 方法:", method.Name)
		}
	}
}
