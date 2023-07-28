package main 

import (
    "fmt"
    "html/template"
    "net/http"
    "time"
    "github.com/gin-gonic/gin"
)

func formatAsDate(t time.Time) string {
    year, month, day := t.Date()
    return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}

func main() {
    router := gin.Default()
    router.Delims("{[{", "}]}")
    router.SetFuncMap(template.FuncMap{
        "formatAsDate": formatAsDate,
    })
    router.LoadHTMLFiles("./static/index.html")

    router.GET("/index", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", map[string]interface{}{
			
		})
    })

    router.Run(":8080")
}
