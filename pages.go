package main
import (
    //"io/ioutil"

    //"net/http"

    "github.com/gin-gonic/gin"

    //"strconv"

    //"fmt"

    //"database/sql"

    _ "github.com/go-sql-driver/mysql"

    //"strings"
    //"github.com/itsjamie/gin-cors"
    //"github.com/gin-contrib/cors"
    //"github.com/gin-gonic/gin"
)


func getPage1(c *gin.Context) {
	c.HTML(200, "homepage.html", gin.H{"code": "Success", "message": "Page found"})
}
func getPage2(c *gin.Context) {
	c.HTML(200, "alerts.html", gin.H{"code": "Success", "message": "Page found"})
}
type Calend struct {
	Month     string
	Day  string
	Date string
}
var calen = []Calend{
	{Month: "December", Day: "Friday", Date: "9"},
    {Month: "December", Day: "Saturday", Date: "10"},
    {Month: "December", Day: "Sunday", Date: "11"},
    {Month: "December", Day: "Monday", Date: "12"},
    {Month: "December", Day: "Tuesday", Date: "13"},
    {Month: "December", Day: "Wednesday", Date: "14"},
    {Month: "December", Day: "Thursday", Date: "15"},
    {Month: "December", Day: "Friday", Date: "16"},
    {Month: "December", Day: "Saturday", Date: "17"},
    {Month: "December", Day: "Sunday", Date: "18"},
    {Month: "December", Day: "Monday", Date: "19"},
    {Month: "December", Day: "Tuesday", Date: "20"},
    {Month: "December", Day: "Wednesday", Date: "21"},
    {Month: "December", Day: "Thursday", Date: "22"},
    {Month: "December", Day: "Friday", Date: "23"},
    {Month: "December", Day: "Saturday", Date: "24"},
    {Month: "December", Day: "Sunday", Date: "25"},
    {Month: "December", Day: "Monday", Date: "26"},
    {Month: "December", Day: "Tuesday", Date: "27"},
    {Month: "December", Day: "Wednesday", Date: "28"},
    {Month: "December", Day: "Thursday", Date: "29"},
    {Month: "December", Day: "Friday", Date: "30"},
    {Month: "December", Day: "Saturday", Date: "31"},
}
func getPage3(c *gin.Context) {
	c.HTML(200, "slider-calendar.html", gin.H{"calend": calen})
}
func getPage4(c *gin.Context) {
	c.HTML(200, "insight.html", gin.H{"code": "Success", "message": "Page found"})
}
//5 to 11
func getPage5(c *gin.Context) {
	c.HTML(200, "video-layout-new.html", gin.H{"code": "Success", "message": "Page found"})
}
func getPage12(c *gin.Context) {
	c.HTML(200, "health-chart.html", gin.H{"code": "Success", "message": "Page found"})
}
func getPage13(c *gin.Context) {
	c.HTML(200, "health-chart-table.html", gin.H{"code": "Success", "message": "Page found"})
}
func getPage14(c *gin.Context) {
	c.HTML(200, "health-chart-dashboard.html", gin.H{"code": "Success", "message": "Page found"})
}
func getPage15(c *gin.Context) {
	c.HTML(200, "calendar.html", gin.H{"code": "Success", "message": "Page found"})
}
func getPage17(c *gin.Context) {
	c.HTML(200, "journal-day.html", gin.H{"code": "Success", "message": "Page found"})
}
func getPage18(c *gin.Context) {
	c.HTML(200, "bespoke-overview.html", gin.H{"code": "Success", "message": "Page found"})
}
func getPage19(c *gin.Context) {
	c.HTML(200, "bespoke-sleep-prescription.html", gin.H{"code": "Success", "message": "Page found"})
}
func getPage20(c *gin.Context) {
	c.HTML(200, "bespoke-sleep-score.html", gin.H{"code": "Success", "message": "Page found"})
}
func getPage21(c *gin.Context) {
	c.HTML(200, "resource-overview.html", gin.H{"code": "Success", "message": "Page found"})
}
func getPage22(c *gin.Context) {
	c.HTML(200, "resources-a.html", gin.H{"code": "Success", "message": "Page found"})
}
func getPage24(c *gin.Context) {
	c.HTML(200, "excercise.html", gin.H{"code": "Success", "message": "Page found"})
}
func getPage25(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{"code": "Success", "message": "Page found"})
}
func getPage26(c *gin.Context) {
	c.HTML(200, "faqs.html", gin.H{"code": "Success", "message": "Page found"})
}
func getPage27(c *gin.Context) {
	c.HTML(200, "stress.html", gin.H{"code": "Success", "message": "Page found"})
}
func getPage28(c *gin.Context) {
	c.HTML(200, "stress.html", gin.H{"code": "Success", "message": "Page found"})
}
func getPage30(c *gin.Context) {
	c.HTML(200, "anxiety-insight-1.html", gin.H{"code": "Success", "message": "Page found"})
}
func getPage33(c *gin.Context) {
	c.HTML(200, "support-walkin.html", gin.H{"code": "Success", "message": "Page found"})
}
func getPage34(c *gin.Context) {
	c.HTML(200, "support-area-look-at.html", gin.H{"code": "Success", "message": "Page found"})
}
func getPage35(c *gin.Context) {
	c.HTML(200, "support-wellbeing.html", gin.H{"code": "Success", "message": "Page found"})
}
func getPage36(c *gin.Context) {
	c.HTML(200, "support-communicate.html", gin.H{"code": "Success", "message": "Page found"})
}