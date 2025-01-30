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
		AccountPort        []int `yaml:"accountPort"`
		FriendPort         []int `yaml:"friendPort"`
		OfflineMessagePort []int `yaml:"offlineMessagePort"`
		GroupPort          []int `yaml:"groupPort"`
	} `yaml:"rpcport"`

	RpcRegisterName struct {
		AccountName        string `yaml:"accountName"`
		FriendName         string `yaml:"friendName"`
		OfflineMessageName string `yaml:"offlineMessageName"`
		GroupName          string `yaml:"groupName"`
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

	Zookeeper struct {
		ZKSchema string   `yaml:"zkSchema"`
		ZKAddr   []string `yaml:"zkAddr"`
	} `yaml:"zookeeper"`

	Mysql struct {
		DBAddress      []string `yaml:"dbMysqlAddress"`
		DBUserName     string   `yaml:"dbMysqlUserName"`
		DBPassword     string   `yaml:"dbMysqlPassword"`
		DBDatabaseName string   `yaml:"dbMysqlDatabaseName"`
	} `yaml:"mysql"`

	Redis struct {
		DBAddress  string `yaml:"dbAddress"`
		DBPassword string `yaml:"dbPassword"`
	} `yaml:"redis"`

	Kafka struct {
		Ws2mschat struct {
			Addr  []string `yaml:"addr"`
			Topic string   `yaml:"topic"`
		} `yaml:"ws2mschat"`
		Ms2pschat struct {
			Addr  []string `yaml:"addr"`
			Topic string   `yaml:"topic"`
		} `yaml:"ms2pschat"`
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
