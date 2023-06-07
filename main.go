//package main

//import "fmt"

//func main(){
//	fmt.Println("hoge")
// }

package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()
  router.LoadHTMLGlob("templates/*")
  router.GET("/index", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.tmpl", gin.H{
      "title": "Main website",
    })
  })
  router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}


