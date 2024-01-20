package cruiser

import (
	"cruiser/route"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() /* 创建一个默认的 Gin 引擎 */

	route.SetupRoutes(r) /* 设置路由 */

	r.Run(":4200") /* 启动服务器，监听端口 8080 */
}
