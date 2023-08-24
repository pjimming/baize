//go:generate statik -src=./static
//go:generate go fmt statik/statik.go
//go:generate go-bindata -o=core-ymal.go etc

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/pjimming/baize/core/internal/config"
	"github.com/pjimming/baize/core/internal/handler"
	"github.com/pjimming/baize/core/internal/svc"
	_ "github.com/pjimming/baize/core/statik"

	"github.com/rakyll/statik/fs"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

//var configFile = flag.String("c", "etc/core-api.yaml", "the config file")

func main() {
	//flag.Parse()

	var c config.Config
	f, _ := Asset("etc/core-api.yaml")
	if err := conf.LoadFromYamlBytes(f, &c); err != nil {
		panic(err)
	}
	logx.MustSetup(c.Log)

	c.Timeout = int64(5 * time.Minute)

	server := rest.MustNewServer(c.RestConf, rest.WithCors())
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	go func() {
		statikFS, _ := fs.New()
		http.Handle("/", http.FileServer(statikFS))
		fmt.Println("Server started at http://localhost:5250")
		_ = http.ListenAndServe(":5250", nil)
	}()

	//fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()

}
