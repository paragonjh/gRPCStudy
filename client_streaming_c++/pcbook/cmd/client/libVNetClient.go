package main

import (
	"C"
	"context"
	"gitlab.com/techschool/pcbook/pb"
	"google.golang.org/grpc"
	"log"
	"time"
	"unsafe"
)

//export UploadImage
func UploadImage(receivedImage unsafe.Pointer, imageSize C.int) {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	//Make ImageFile to Golang Byte
	log.Printf("requested image size: %d", imageSize)
	imageData := C.GoBytes(receivedImage, imageSize)

	//Creates a client connection to the VNet server
	conn, err := grpc.Dial("localhost:4343", grpc.WithInsecure())
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	laptopClient := pb.NewLaptopServiceClient(conn)
	uploadImage(laptopClient, imageData)
}

func uploadImage(laptopClient pb.LaptopServiceClient, imageData []byte) {

	log.Printf("requested image size: %d", len(imageData))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := laptopClient.UploadImage(ctx)
	if err != nil {
		log.Fatal("cannot upload image: ", err)
	}

	req := &pb.UploadImageRequest{
		Data: &pb.UploadImageRequest_Info{
			Info: &pb.ImageInfo{
				//To-Do: Add more image informations
				ImageName: "testImage",
				ImageType: ".jpg",
			},
		},
	}

	log.Printf("Image Send Start")
	err = stream.Send(req)
	if err != nil {
		log.Fatal("cannot send image info to server: ", err)
	}

	//File chunk by 1kb and send image
	chunkSize := 1024
	for i := 0; i < len(imageData); i += chunkSize {
		end := i + chunkSize
		log.Printf("i: %d, end: %d", i, end)
		if end > len(imageData) {
			end = len(imageData)
		}
		buffer := imageData[i:end]
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

	log.Printf("Image Send End")
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("cannot receive response: ", err)
	}

	log.Printf("image uploaded with id: %s, size: %d", res.GetId(), res.GetSize())
}

// Client Stream gRPC: Server App
func main() {

}
