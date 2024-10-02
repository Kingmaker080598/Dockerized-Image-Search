package main

import (
    "context"
    // "io/ioutil"
    "log"
    "net/http"
    // "strings"
    "encoding/base64"
    "github.com/gin-gonic/gin"
    "google.golang.org/grpc"
    "client/pb"
    "fmt"
)

func main() {
    router := gin.Default()
    router.LoadHTMLGlob("templates/*")

    router.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", nil)
    })

    router.POST("/", func(c *gin.Context) {
        keyword := c.PostForm("keyword")
        imageBase64, errMessage := searchImage(keyword)
        c.HTML(http.StatusOK, "result.html", gin.H{
            "imageBase64": imageBase64,
            "errorMessage": errMessage,
        })
    })

    log.Println("Starting web server on port 8080...")
    if err := router.Run(":8081"); err != nil {
        log.Fatalf("Failed to start web server: %v", err)
    }
}

func searchImage(keyword string) (string, string) {
    conn, err := grpc.Dial("grpc-server:8080", grpc.WithInsecure())
    if err != nil {
        return "", "Failed to connect to the gRPC server"
    }
    defer conn.Close()

    client := pb.NewImageSearchServiceClient(conn)
    ctx := context.Background()
    request := &pb.ImageServiceRequest{Keyword: keyword}
    response, err := client.SearchImage(ctx, request)
    fmt.Println(response)
    fmt.Println(err)
    if err != nil {
        return "", "No images found for the keyword"
    }

    imageBase64 := "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(response.ImageData)
    return imageBase64, ""
}
