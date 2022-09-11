package center

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/edte/erpc/log"
	"github.com/gin-gonic/gin"
)

var (
	defaultCenter = NewCenter()
)

type ResigerArg struct {
	Server   string
	Addr     string
	Request  interface{}
	Response interface{}
}

// 服务注册
// 一台机器向注册中心发送本地地址、注册的服务和接口，以及接口的参数、返回值
// TODO: 这里考虑优化：把一个 server 的所有 func 一次性请求过来注册？
// 这里注册中心主要存 req、rsp type 吗？有必要吗？
func Register(server string, addr string, request interface{}, response interface{}) (err error) {
	r := &ResigerArg{
		Server:   server,
		Addr:     addr,
		Request:  request,
		Response: response,
	}

	b, err := json.Marshal(r)
	if err != nil {
		return
	}

	resp, err := http.Post("http://127.0.0.1:8080/register", "application/json", strings.NewReader(string(b)))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if resp.StatusCode == 500 {
		return errors.New(string(body))
	}

	return
}

// 服务发现
// 给定请求服务名，然后负债均衡返回其中一个部署的机器 ip 地址
// TODO: 每次响应一个 ip，那其他集群内怎么同步的？
func Discovery(server string) (addr string, err error) {
	resp, err := http.Get("http://127.0.0.1:8080/discovery?server=" + server)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if resp.StatusCode == 500 {
		return "", errors.New(string(body))
	}

	return string(body), nil
}

func Listen(addr string) {
	r := gin.Default()

	r.GET("/discovery", func(ctx *gin.Context) {
		s := ctx.Query("server")
		res, err := defaultCenter.discovery(s)
		if err != nil {
			log.Errorf("server %s discovery failed, err:%s", s, err)
			ctx.String(500, err.Error())
			return
		}
		log.Debugf("serve %s discover succ, res:%s", s, res)
		ctx.String(200, res)
	})

	r.POST("/register", func(ctx *gin.Context) {
		r := &ResigerArg{}

		if err := ctx.BindJSON(&r); err != nil {
			log.Errorf("registe failed, err:%s", err)
			ctx.String(500, err.Error())
			return
		}

		if err := defaultCenter.register(r.Server, r.Addr, r.Request, r.Response); err != nil {
			log.Errorf("registe failed, err:%s", err)
			ctx.String(500, err.Error())
			return
		}

		log.Debugf("register %s succ", r.Server)
	})

	go Ping()

	r.Run()
}