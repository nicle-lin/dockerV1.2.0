package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/nicle-lin/dockerV1.2.0/docker/api"
	"github.com/nicle-lin/dockerV1.2.0/docker/api/client"
	"github.com/nicle-lin/dockerV1.2.0/docker/dockerversion"
	flag "github.com/nicle-lin/dockerV1.2.0/docker/pkg/mflag"
	"github.com/nicle-lin/dockerV1.2.0/docker/reexec"
	"github.com/nicle-lin/dockerV1.2.0/docker/utils"
)

const (
	defaultCaFile   = "ca.pem"
	defaultKeyFile  = "key.pem"
	defaultCertFile = "cert.pem"
)

func main() {
	//与一些驱动有关，第一次运行没有，所以返回false
	if reexec.Init() {
		return
	}
	//docker自己重写的flag,与go原生的flag唯一区别是可以解析多个参数
	/*
		docker: flVersion     = flag.Bool([]string{"v", "-version"}, false, "Print version information and quit")
		go:   	flVersion     = flag.Bool("v", false, "Print version information and quit")
		docker把类型string改为string切片
	*/

	//当解析参数出错时，会把出错之后的所有参数存在flag.Args(),以便之后执行Docker client具体请求时使用
	flag.Parse()
	// FIXME: validate daemon flags here

	//下面的参数都定义在同目录下flags.go里
	if *flVersion {
		showVersion()
		return
	}
	if *flDebug {
		os.Setenv("DEBUG", "1")
	}

	if len(flHosts) == 0 {
		defaultHost := os.Getenv("DOCKER_HOST")
		if defaultHost == "" || *flDaemon {
			// If we do not have a host, default to unix socket
			defaultHost = fmt.Sprintf("unix://%s", api.DEFAULTUNIXSOCKET)
		}
		if _, err := api.ValidateHost(defaultHost); err != nil {
			log.Fatal(err)
		}
		flHosts = append(flHosts, defaultHost)
	}

	//
	if *flDaemon {
		mainDaemon()
		return
	}

	if len(flHosts) > 1 {
		log.Fatal("Please specify only one -H")
	}
	//flHosts like: tcp://host:port, unix:///path/to/socket, fd://* or fd://socketfd.
	protoAddrParts := strings.SplitN(flHosts[0], "://", 2)
	//把tcp://host:port 分成两部分tcp及 host:port
	//但为什么flHosts是[]string类型而不是string类型，难道docker daemon要监听多个地址?
	//TODO: to find out
	//Answer:哈哈，这是用于客户端的,客户端可以同时向多个地址请求?

	var (
		cli       *client.DockerCli
		tlsConfig tls.Config
	)
	tlsConfig.InsecureSkipVerify = true

	// If we should verify the server, we need to load a trusted ca
	if *flTlsVerify {
		*flTls = true
		certPool := x509.NewCertPool()
		file, err := ioutil.ReadFile(*flCa)
		if err != nil {
			log.Fatalf("Couldn't read ca cert %s: %s", *flCa, err)
		}
		certPool.AppendCertsFromPEM(file)
		tlsConfig.RootCAs = certPool
		tlsConfig.InsecureSkipVerify = false
	}

	// If tls is enabled, try to load and send client certificates
	if *flTls || *flTlsVerify {
		_, errCert := os.Stat(*flCert)
		_, errKey := os.Stat(*flKey)
		if errCert == nil && errKey == nil {
			*flTls = true
			cert, err := tls.LoadX509KeyPair(*flCert, *flKey)
			if err != nil {
				log.Fatalf("Couldn't load X509 key pair: %s. Key encrypted?", err)
			}
			tlsConfig.Certificates = []tls.Certificate{cert}
		}
	}

	if *flTls || *flTlsVerify {
		cli = client.NewDockerCli(os.Stdin, os.Stdout, os.Stderr, protoAddrParts[0], protoAddrParts[1], &tlsConfig)
	} else {
		cli = client.NewDockerCli(os.Stdin, os.Stdout, os.Stderr, protoAddrParts[0], protoAddrParts[1], nil)
	}
	//当解析参数出错时，会把出错之后的所有参数存在flag.Args(),以便之后执行Docker client具体请求时使用
	//解析参数出错时执行
	//如命令:docker -H unix:///var/run/socket pull ubuntu:latest,这样保存在flag.Args()的参数是:pull ubuntu:latest
	if err := cli.Cmd(flag.Args()...); err != nil {
		if sterr, ok := err.(*utils.StatusError); ok {
			if sterr.Status != "" {
				log.Println(sterr.Status)
			}
			os.Exit(sterr.StatusCode)
		}
		log.Fatal(err)
	}
}

func showVersion() {
	fmt.Printf("Docker version %s, build %s\n", dockerversion.VERSION, dockerversion.GITCOMMIT)
}
