package main
import (
    //"github.com/gin-contrib/static"

    //"io/ioutil"

    "net/http"

    "github.com/gin-gonic/gin"

    "strconv"

    "fmt"

    "database/sql"

    _ "github.com/go-sql-driver/mysql"

    "strings"
    //"github.com/itsjamie/gin-cors"
    //"github.com/gin-contrib/cors"
    //"github.com/gin-gonic/gin"
)
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []Book{
	{ID: "1", Title: "Harry Potter", Author: "J. K. Rowling"},
	{ID: "2", Title: "The Lord of the Rings", Author: "J. R. R. Tolkien"},
	{ID: "3", Title: "The Wizard of Oz", Author: "L. Frank Baum"},
}
func main() {
    

    fmt.Println("hello world")
    main2()
}
func main2() {
    router := gin.Default()
    router.LoadHTMLGlob("HTMLs/*")
    router.Static("/img","./img")
    router.Static("/js","./js")
    router.Static("/static-html","./static-html")
    router.Static("/client-logic","./client-logic")
    router.Static("/css","./css")
    router.Static("/fonts","./fonts")
    router.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, ResponseType, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
    
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        c.Next()
    })
    //router.Use(cors.Default())
    //router.GET("/albums", c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"}))
    router.GET("/books/:id", getBooks)
    router.POST("/albums", postAlbums)
    router.POST("/register", register)
    //1
    router.GET("/home", getPage1)
    //2
    router.GET("/2", getPage2)
    //3
    router.GET("/3", getPage3)
    //4
    router.GET("/4", getPage4)
    //5to11
    router.GET("/5", getPage5)
    //12
    router.GET("/12", getPage12)
    //13
    router.GET("/13", getPage13)
    //14
    router.GET("/14", getPage14)
    //15
    router.GET("/15", getPage15)
    //17
    router.GET("/17", getPage17)
    //18
    router.GET("/18", getPage18)
    //19
    router.GET("/19", getPage19)
    //20
    router.GET("/20", getPage20)
    //21
    router.GET("/21", getPage21)
    //22
    router.GET("/22", getPage22)
    //24
    router.GET("/24", getPage24)
    //25
    router.GET("/25", getPage25)
    //26
    router.GET("/26", getPage26)
    //27
    router.GET("/27", getPage27)
    //28
    router.GET("/28", getPage28)
    //30
    router.GET("/30", getPage30)
    //33
    router.GET("/33", getPage33)
    //34
    router.GET("/34", getPage34)
    //35
    router.GET("/35", getPage35)
    //36
    router.GET("/36", getPage36)
    

    router.Run("localhost:8080")
}
func getBooks(c *gin.Context) {
    id:=c.Param("id")
    ID,err:=strconv.Atoi(id)
    if err!=nil{
        c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
    }else
    {
        fmt.Println(ID)
        c.JSON(http.StatusOK, books)
    }
}
type user struct {
    Name     string
    Email  string
    Password string
}
func register(c *gin.Context){
    var newUser user

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newUser); err != nil {
        return
    }
    fmt.Println(newUser.Name)
    // Add the new album to the slice.
    //albums = append(albums, newAlbum)
    db, err := sql.Open("mysql", "gin:gin@tcp(127.0.0.1:3306)/gin")
    if err!=nil{
        fmt.Println("lalalala")
    }
    var sb strings.Builder
    fmt.Println("cucu")
    sb.WriteString("INSERT into gin_users VALUES(\"")
    fmt.Println("data")
    sb.WriteString(newUser.Name)
    fmt.Println("data2")
    sb.WriteString("\",\"")
    sb.WriteString(newUser.Email)
    sb.WriteString("\",\"")
    sb.WriteString(newUser.Password)
    sb.WriteString("\")")
    fmt.Println("data2")
    fmt.Println(sb.String())
    _, err = db.Query(sb.String())
    fmt.Println("data2")
    // handle error
    if err != nil {
        fmt.Println("error")
        panic(err)
    }else{
        fmt.Println("Account Created")
    }
    db.Close()
    //c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
    c.IndentedJSON(http.StatusCreated, newUser)
}
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}
func postAlbums(c *gin.Context) {
    var newAlbum album

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }
    fmt.Println(newAlbum.Artist)
    // Add the new album to the slice.
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}