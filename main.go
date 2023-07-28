package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"math"
	"net/http"
	"time"
)

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}

func PhuongTrinhBac2() {
	type ThamSo struct {
		a float64
		b float64
		c float64
	}

	type nghiem struct {
		x1 float64
		x2 float64
	}

	cacThamSo := ThamSo{
		a: 12,
		b: 500,
		c: 40,
	}

	delta := tinhDelta(cacThamSo.a, cacThamSo.b, cacThamSo.c)

	if delta < 0 {
		fmt.Println("Phương trình vô nghiệm")
	} else if delta > 0 {
		ketQua := nghiem{
			x1: (-(cacThamSo.b) + math.Sqrt(delta)) / 2 * cacThamSo.a,
			x2: (-(cacThamSo.b) - math.Sqrt(delta)) / 2 * cacThamSo.a,
		}
		fmt.Println("Phương trình có 2 nghiệm là ", ketQua)
	} else {
		ketQua := nghiem{
			x1: -(cacThamSo.b) / 2 * cacThamSo.a,
			x2: -(cacThamSo.b) / 2 * cacThamSo.a,
		}
		fmt.Println("Phương trình có nghiệm kép", ketQua)
	}

	fmt.Println("Gia Tri delta la", tinhDelta(cacThamSo.a, cacThamSo.b, cacThamSo.c))
}

// Giai Phuong Trinh bac 2

func tinhDelta(a float64, b float64, c float64) float64 {

	return b*b - (4 * a * c)
}

func fetchUsers() string {
	data := "Anh Thiện Đẹp Trai"
	return data
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
			"strucdemo": fetchUsers(),
		})
	})

	router.Run(":7000")
}
