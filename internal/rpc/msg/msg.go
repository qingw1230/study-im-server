package msg

import (
	"net"
	"strconv"

	"github.com/qingw1230/study-im-server/pkg/common/config"
	"github.com/qingw1230/study-im-server/pkg/common/kafka"
	"github.com/qingw1230/study-im-server/pkg/common/log"
	pbMsg "github.com/qingw1230/study-im-server/pkg/proto/msg"
	"github.com/qingw1230/study-im-server/pkg/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type msgServer struct {
	pbMsg.UnimplementedMsgServer
	rpcPort         int
	rpcRegisterName string
	zkSchema        string
	zkAddr          []string
	producer        *kafka.Producer
}

func NewMsgServer(port int) *msgServer {
	log.NewPrivateLog("msg")
	rpc := &msgServer{
		rpcPort:         port,
		rpcRegisterName: config.Config.RpcRegisterName.OfflineMessageName,
		zkSchema:        config.Config.Zookeeper.ZKSchema,
		zkAddr:          config.Config.Zookeeper.ZKAddr,
	}
	rpc.producer = kafka.NewKafkaProducer(config.Config.Kafka.Ws2mschat.Addr, config.Config.Kafka.Ws2mschat.Topic)
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

	// TODO(qingw1230): 将 rpc 服务注册进 zk
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
