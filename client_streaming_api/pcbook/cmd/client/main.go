package main

import (
	"bufio"
	"context"
	"gitlab.com/techschool/pcbook/pb"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

// Client Stream gRPC: Server App
func main() {
	conn, err := grpc.Dial("localhost:4343", grpc.WithInsecure())
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	laptopClient := pb.NewLaptopServiceClient(conn)
	testUploadImage(laptopClient)
}

func testUploadImage(laptopClient pb.LaptopServiceClient) {
	//To-Do: Make Test Image Information
	imageName := "testImg"
	uploadImage(laptopClient, imageName, "tmp/laptop.jpg")
}

func uploadImage(laptopClient pb.LaptopServiceClient, imageName string, imagePath string) {
	file, err := os.Open(imagePath)
	if err != nil {
		log.Fatal("cannot open image file: ", err)
	}
	defer file.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := laptopClient.UploadImage(ctx)
	if err != nil {
		log.Fatal("cannot upload image: ", err)
	}

	req := &pb.UploadImageRequest{
		Data: &pb.UploadImageRequest_Info{
			Info: &pb.ImageInfo{
				ImageName: imageName,
				ImageType: filepath.Ext(imagePath),
			},
		},
	}

	err = stream.Send(req)
	if err != nil {
		log.Fatal("cannot send image info to server: ", err)
	}

	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)

	for {
		//File Chunk by buffer size(1kb)
		_, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("cannot read chunk to buffer: ", err)
		}

		req := &pb.UploadImageRequest{
			Data: &pb.UploadImageRequest_ChunkData{
				ChunkData: buffer[:],
			},
		}

		err = stream.Send(req)
		if err != nil {
			log.Fatal("cannot send chunk to server: ", err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("cannot receive response: ", err)
	}

	log.Printf("image uploaded with id: %s, size: %d", res.GetId(), res.GetSize())
}
