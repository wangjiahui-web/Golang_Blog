package main

import (
	_ "blogProject/docs"
	"blogProject/internal/router"
	"blogProject/pkg/color"
	"blogProject/pkg/config"
	"fmt"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// https://patorjk.com/software/taag/#p=display&f=ANSI%20Shadow&t=blog-project-api
const logo = `
██████╗ ██╗      ██████╗  ██████╗       ██████╗ ██████╗  ██████╗      ██╗███████╗ ██████╗████████╗    █████╗ ██████╗ ██╗
██╔══██╗██║     ██╔═══██╗██╔════╝       ██╔══██╗██╔══██╗██╔═══██╗     ██║██╔════╝██╔════╝╚══██╔══╝   ██╔══██╗██╔══██╗██║
██████╔╝██║     ██║   ██║██║  ███╗█████╗██████╔╝██████╔╝██║   ██║     ██║█████╗  ██║        ██║█████╗███████║██████╔╝██║
██╔══██╗██║     ██║   ██║██║   ██║╚════╝██╔═══╝ ██╔══██╗██║   ██║██   ██║██╔══╝  ██║        ██║╚════╝██╔══██║██╔═══╝ ██║
██████╔╝███████╗╚██████╔╝╚██████╔╝      ██║     ██║  ██║╚██████╔╝╚█████╔╝███████╗╚██████╗   ██║      ██║  ██║██║     ██║
╚═════╝ ╚══════╝ ╚═════╝  ╚═════╝       ╚═╝     ╚═╝  ╚═╝ ╚═════╝  ╚════╝ ╚══════╝ ╚═════╝   ╚═╝      ╚═╝  ╚═╝╚═╝     ╚═╝
                                                                                                                        `

//	@title			这里写标题
//	@version		1.0
//	@description	这里写描述信息
//	@termsOfService	https://swagger.io/terms/

//	@contact.name	这里写联系人信息
//	@contact.url	https://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	https://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description				Description for what is this security definition being used
func main() {

	// swagger 页面地址:   http://localhost:8080/swagger/index.html


	fmt.Println(color.Blue(logo))

	s := router.NewHttpServer()

	server := &http.Server{
		Addr:    ":" + config.Get().ServerConfig.Port,
		Handler: s,
	}

	// 开启协程监听网络请求
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.S().Errorf("listen: %s", err)
		}
	}()

	quit := make(chan os.Signal)

	// kill
	// kill -2 syscall.SIGINT
	// kill -9 syscall.SIGKILL
	// kill -15
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		zap.S().Errorf("shut down err: %s", err)
	}

	select {
	case <-ctx.Done():
		zap.L().Info("3秒超时退出")
	}

	zap.L().Info("服务器关闭")


}
