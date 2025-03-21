package config

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

var (
	_, b, _, _ = runtime.Caller(0)
	Root       = filepath.Join(filepath.Dir(b), "../../..")
)

var Config config

type config struct {
	ServerIP string `yaml:"serverip"`

	RpcPort struct {
		AccountPort            []int `yaml:"accountPort"`
		FriendPort             []int `yaml:"friendPort"`
		OfflineMessagePort     []int `yaml:"offlineMessagePort"`
		OnlineMessageRelayPort []int `yaml:"onlineMessageRelayPort"`
		GroupPort              []int `yaml:"groupPort"`
		PushPort               []int `yaml:"pushPort"`
		ConversationPort       []int `yaml:"conversationPort"`
	} `yaml:"rpcport"`

	RpcRegisterName struct {
		AccountName            string `yaml:"accountName"`
		FriendName             string `yaml:"friendName"`
		OfflineMessageName     string `yaml:"offlineMessageName"`
		OnlineMessageRelayName string `yaml:"onlineMessageRelayName"`
		GroupName              string `yaml:"groupName"`
		PushName               string `yaml:"pushName"`
		ConversationName       string `yaml:"conversationName"`
	} `yaml:"rpcregistername"`

	Admin struct {
		UserIds []string `yaml:"userIds"`
	} `yaml:"admin"`

	TokenPolicy struct {
		Secret       string `yaml:"secret"`
		AccessExpire int64  `yaml:"accessExpire"`
	} `yaml:"tokenpolicy"`

	Log struct {
		StorageLocation     string `yaml:"storageLocation"`
		RotationTime        int    `yaml:"rotationTime"`
		RemainRotationCount uint   `yaml:"remainRotationCount"`
		RemainLogLevel      uint   `yaml:"remainLogLevel"`
	} `yaml:"log"`

	LongConnSvr struct {
		WebsocketPort       []int `yaml:"wsPort"`
		WebsocketMaxConnNum int   `yaml:"websocketMaxConnNum"`
		WebsocketMaxMsgLen  int   `yaml:"websocketMaxMsgLen"`
		WebsocketTimeout    int   `yaml:"websocketTimeout"`
	} `yaml:"longconnsvr"`

	Etcd struct {
		EtcdSchema string   `yaml:"etcdSchema"`
		EtcdAddr   []string `yaml:"etcdAddr"`
	} `yaml:"etcd"`

	Mysql struct {
		DBAddress      []string `yaml:"dbMysqlAddress"`
		DBUserName     string   `yaml:"dbMysqlUserName"`
		DBPassword     string   `yaml:"dbMysqlPassword"`
		DBDatabaseName string   `yaml:"dbMysqlDatabaseName"`
	} `yaml:"mysql"`

	Mongo struct {
		DBAddress           []string `yaml:"dbAddress"`
		DBUserName          string   `yaml:"dbMongoUserName"`
		DBPassword          string   `yaml:"dbMongoPassword"`
		DBDirect            bool     `yaml:"dbDirect"`
		DBTimeout           int      `yaml:"dbTimeout"`
		DBDatabase          string   `yaml:"dbDatabase"`
		DBSource            string   `yaml:"dbSource"`
		DBMaxPoolSize       int      `yaml:"dbMaxPoolSize"`
		DBRetainChatRecords int      `yaml:"dbRetainChatRecords"`
	} `yaml:"mongo"`

	Redis struct {
		DBAddress  string `yaml:"dbAddress"`
		DBPassword string `yaml:"dbPassword"`
	} `yaml:"redis"`

	Kafka struct {
		OfflineMsgToMongoMysql struct {
			Addr  []string `yaml:"addr"`
			Topic string   `yaml:"topic"`
		} `yaml:"offlineMsgToMongoMysql"`
		MsgToPush struct {
			Addr  []string `yaml:"addr"`
			Topic string   `yaml:"topic"`
		} `yaml:"msgToPush"`
		ConsumerGroupId struct {
			MsgToMongo string `yaml:"msgToMongo"`
			MsgToMySql string `yaml:"msgToMySql"`
			MsgToPush  string `yaml:"msgToPush"`
		}
	} `yaml:"kafka"`
}

func init() {
	cfgName := Root + "/config/config.yaml"
	viper.SetConfigFile(cfgName)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}
	bytes, err := os.ReadFile(cfgName)
	if err != nil {
		panic(err.Error())
	}
	if err = yaml.Unmarshal(bytes, &Config); err != nil {
		panic(err.Error())
	}
}
