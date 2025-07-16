package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

var RouteList = make([]*RouteInfo, 0)

// RouteInfo 路由信息
type RouteInfo struct {
	Method HttpType
	Url    string
	Func   gin.HandlerFunc
}

// AddRoute 添加路由
func AddRoute(method HttpType, url string, handler gin.HandlerFunc) {
	RouteList = append(RouteList, &RouteInfo{
		Method: method,
		Url:    url,
		Func:   handler,
	})
}

// RegisterRoute 注册路由
func RegisterRoute(r *gin.Engine) {
	for _, route := range RouteList {
		switch route.Method {
		case HTTP_TYPE_GET:
			r.GET(route.Url, route.Func)
		case HTTP_TYPE_POST:
			r.POST(route.Url, route.Func)
		case HTTP_TYPE_PUT:
			r.PUT(route.Url, route.Func)
		case HTTP_TYPE_DELETE:
			r.DELETE(route.Url, route.Func)
		default:
			fmt.Printf("Not support method: %s\n", route.Method)
			os.Exit(1)
		}
	}
}

// Run 运行服务
func Run(port int) {
	if len(RouteList) == 0 {
		fmt.Println("Not found route info, exit!")
		os.Exit(2)
	}
	r := gin.Default()
	RegisterRoute(r)
	if err := r.Run(fmt.Sprintf(":%d", port)); err != nil {
		fmt.Println("Run http server error:", err)
		os.Exit(3)
	}
}
