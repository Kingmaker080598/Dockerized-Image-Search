package main

import (
    "context"
    "server/pb"
    "log"
    "net"
	"io/ioutil"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "math/rand"
    "fmt"
)

type server struct {
    pb.ImageSearchServiceServer
}

var imageDatabase = map[string][]string{
    "dog":       {"dog_image1.jpg", "dog_image2.jpg", "dog_image3.jpg"},
    "cat":       {"cat_image1.jpg", "cat_image2.jpg", "cat_image3.jpg"},
    "home":      {"home_image1.jpg", "home_image2.jpg", "home_image3.jpg"},
    "transport": {"car_image1.jpg", "train_image2.jpg", "bus_image3.jpg", "ship_image4.jpg", "aeroplane_image5.jpg"},
}

func (s *server) SearchImage(ctx context.Context, req *pb.ImageServiceRequest) (*pb.ImageServiceReply, error) {
    keyword := req.Keyword
    fmt.Println(keyword)

    imageFilenames, found := imageDatabase[keyword]
    if !found {
        return nil, status.Error(codes.NotFound, "Keyword not found")
    }

    // Choose a random image filename
    selectedImage := imageFilenames[rand.Intn(len(imageFilenames))]

    imageData, err := loadImage(selectedImage)
    if err != nil {
        return nil, status.Error(codes.Internal, "Failed to load image data")
    }
	if imageData == nil {
        return nil, status.Error(codes.Internal, "Failed to load image data1")
    }

    return &pb.ImageServiceReply{ImageData: imageData}, nil
}

func loadImage(filename string) ([]byte, error) {
    // Replace this with your image loading logic
    // Read and return the image data from the specified filename.
    // For example, read the image file from a local directory.
    imageBytes, err := ioutil.ReadFile("images/" + filename)
    if err != nil {
        return nil, err
    }
    fmt.Println("image found")
    return imageBytes, nil
}

func main() {
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        panic(err)
    }

    s := grpc.NewServer()
    pb.RegisterImageSearchServiceServer(s, &server{})
    if err := s.Serve(listener); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
