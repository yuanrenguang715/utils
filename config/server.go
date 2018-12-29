package config

import (
	"crypto/tls"
	"log"
	"os"

	"github.com/go-ini/ini"
	consulapi "github.com/hashicorp/consul/api"
)

// SrvConfig 服务配置
type SrvConfig struct {
	ServerName string
	ConsulConf *consulapi.Config
	TlsConf    *tls.Config
	LogTrack   bool
}

var SrvCfg SrvConfig

func SrvCfgInit() {
	cfg, err := ini.Load("./config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	cfgSec := cfg.Section("server")

	err = os.Setenv("MICRO_SERVER_ADDRESS", cfgSec.Key("srv_addrs").String())
	if err != nil {
		log.Fatalln(err)
	}

	cer, err := tls.LoadX509KeyPair(cfgSec.Key("grpc_tls_certfile").String(), cfgSec.Key("grpc_tls_keyfile").String())
	if err != nil {
		log.Fatalln(err)
	}
	SrvCfg.TlsConf = &tls.Config{
		MinVersion:   tls.VersionTLS12,
		ClientAuth:   tls.NoClientCert,
		Certificates: []tls.Certificate{cer},
	}

	consulCfg := consulapi.DefaultConfig()
	consulCfg.Address = cfgSec.Key("consul_addrs").String()
	//consulCfg.Token = cfgSec.Key("consul_acl_token").String()

	SrvCfg.ConsulConf = consulCfg
	SrvCfg.ServerName = cfgSec.Key("srv_name").String()
	SrvCfg.LogTrack, _ = cfgSec.Key("log_track_enable").Bool()
}
