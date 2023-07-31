package main

import (
	"fmt"
	"html/template"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}

// Variable

var DetailNghiem string

var ketQua []any

func PhuongTrinhBac2(delta float64, x float64, y float64, z float64) {

	fmt.Println("Giá trị delta là:", delta)

	if delta < 0 {
		fmt.Println("Phương trình vô nghiệm")
		ketQua = []any{
			"vô nghiệm",
			"vô nghiệm",
		}
	} else if delta > 0 {
		ketQua = []any{
			(-(y) + math.Sqrt(delta)) / 2 * x,
			(-(y) - math.Sqrt(delta)) / 2 * x,
		}
		fmt.Println("Phương trình có 2 nghiệm là ", ketQua)
	} else {
		ketQua = []any{
			-(y) / 2 * x,
			-(y) / 2 * x,
		}
		fmt.Println("Phương trình có nghiệm kép", ketQua)
	}

}

// Giai Phuong Trinh bac 2

func tinhDelta(a float64, b float64, c float64) float64 {

	return b*b - (4 * a * c)
}

func fetchUsers() string {
	data := "GIẢI PHƯƠNG TRÌNH BẬC 2"
	return data
}

func main() {
	router := gin.Default()
	router.Delims("{[{", "}]}")
	router.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	router.LoadHTMLFiles("./static/index.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", map[string]interface{}{
			"strucdemo": fetchUsers(),
		})
	})

	router.POST("/giai-phuong-trinh", func(c *gin.Context) {

		type ThamSo struct {
			x float64
			y float64
			z float64
		}
		var cacThamSo ThamSo

		// Get form values and convert them to float64
		x, err := strconv.ParseFloat(c.PostForm("x"), 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid value for 'x'"})
			return
		}
		cacThamSo.x = x

		y, err := strconv.ParseFloat(c.PostForm("y"), 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid value for 'y'"})
			return
		}
		cacThamSo.y = y

		z, err := strconv.ParseFloat(c.PostForm("z"), 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid value for 'z'"})
			return
		}
		cacThamSo.z = z

		fmt.Println(cacThamSo)
		delta := tinhDelta(cacThamSo.x, cacThamSo.y, cacThamSo.z)

		if delta < 0 {
			DetailNghiem = "Vô Nghiệm"
		} else if delta > 0 {
			DetailNghiem = "Hai Nghiệm"
		} else {
			DetailNghiem = "Nghiệm kép"
		}

		PhuongTrinhBac2(delta, cacThamSo.x, cacThamSo.y, cacThamSo.z)

		fmt.Println("Kết quả nghiệm là:", ketQua)

		c.HTML(http.StatusOK, "index.html", map[string]interface{}{
			"strucdemo": fetchUsers(),
			"delta":     tinhDelta(cacThamSo.x, cacThamSo.y, cacThamSo.z),
			"nghiem":    DetailNghiem,
			"nghiem1":   ketQua[0],
			"nghiem2":   ketQua[1],
		})
	})

	router.Run(":7000")
}
