package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
"gopkg.in/gomail.v2"
)


func main() {
	r := gin.Default()
	r.LoadHTMLGlob("*.html")
//	r.StaticFile("/index.html", "index.html")
	r.StaticFS("/css", http.Dir("css"))
	r.StaticFS("/images", http.Dir("images"))
	r.StaticFS("/scripts", http.Dir("scripts"))
	r.Use(Cors())
	r.GET("/", func(c *gin.Context) {
			    c.HTML(http.StatusOK, "index.html", gin.H{})
			})
	r.POST("/sendMail", SendMail)
	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{
				"title": "Main website",
		})
	})
	r.Run(":8080")
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func SendMail(c *gin.Context) {
	name := c.PostForm("name")
	from := c.PostForm("email")
	message := c.PostForm("message")
	if name!="" && from != "" && message !=""{
		m := gomail.NewMessage()
    m.SetAddressHeader("From", "example@ya.ru", name)
    m.SetAddressHeader("To", "wtcute.info@gmail.com", "Me")
    m.SetHeader("Subject", "Order request")
		m.SetHeader("Reply-to", from)
    m.SetBody("text/plain", message)

    d := gomail.NewPlainDialer("smtp.yandex.ru", 587, "login", "pass")

    if err := d.DialAndSend(m); err != nil {
        panic(err)
    }
		// d := gomail.Dialer{Host: "localhost", Port: 587}
		// if err := d.DialAndSend(m); err != nil {
    // 	panic(err)
		// }
	}
}
