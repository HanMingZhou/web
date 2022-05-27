package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
func sayhello(w http.ResponseWriter,r *http.Request) {
	b, _ := ioutil.ReadFile("./hello.txt")
	_, _ = fmt.Fprintln(w,	string(b))
}


func main() {
	http.HandleFunc("/hello",sayhello)
	err := http.ListenAndServe(":9090",nil)
	if err != nil {
		fmt.Printf("http sever failed: %v\n",err)
		return
	}
} 
*/

func sayhello(c *gin.Context) {
	c.JSON(200,gin.H {
		"message" : "hello golang!",
	})
	
}
func main() {
	r := gin.Default() //返回默认的路由引擎

	//指定用户使用GET请求访问/hello时，执行sayhello这个函数
	r.GET("/hello",sayhello)

	r.GET("/book",	func (c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "get",
		})
	})
	r.POST("/book", func(c *gin.Context) {
		c.JSON(http.StatusAccepted, gin.H{
			"method": "POST",
		})
	})
	r.PUT("/book",	func(c *gin.Context) {
		c.JSON(http.StatusAccepted,gin.H{
			"method" : "PUT",
		})
	})
	r.DELETE("/book",	func(c *gin.Context) {
		c.JSON(http.StatusAccepted,gin.H{
			"method" : "DELETE",
		})
	})

	//启动服务
	r.Run(":9090")
}