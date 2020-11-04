/*
Unary gRPC: Server Code
*/

package service

import (
	"bytes"
	"context"
	"gitlab.com/techschool/pcbook/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
)

// Max Image Size : 1MB
const maxImageSize = 1 << 20

// LaptopServer is the server that provides laptop services
type LaptopServer struct {
	imageStore ImageStore
}

// NewLaptopServer returns a new LaptopServer
func NewLaptopServer(imageStore ImageStore) *LaptopServer {
	return &LaptopServer{imageStore}
}

func (server *LaptopServer) UploadImage(stream pb.LaptopService_UploadImageServer) error {
	req, err := stream.Recv()
	if err != nil {
		return logError(status.Errorf(codes.Unknown, "cannot receive image info"))
	}

	imageName := req.GetInfo().GetImageName()
	imageType := req.GetInfo().GetImageType()
	log.Printf("receive an upload-image request for imageName %s with image type %s", imageName, imageType)

	imageData := bytes.Buffer{}
	imageSize := 0
	log.Printf("Start Image Download, Name: %s", imageName)

	for {
		if err := contextError(stream.Context()); err != nil {
			return err
		}
		//log.Print("waiting to receive more data")

		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("no more data")
			break
		}
		if err != nil {
			return logError(status.Errorf(codes.Unknown, "cannot receive chunk data: %v", err))
		}

		chunk := req.GetChunkData()
		size := len(chunk)

		//log.Printf("received a chunk with size: %d", size)
		imageSize += size
		if imageSize > maxImageSize {
			return logError(status.Errorf(codes.InvalidArgument, "image is too large: %d > %d", imageSize, maxImageSize))
		}
		_, err = imageData.Write(chunk)
		if err != nil {
			return logError(status.Errorf(codes.Internal, "cannot write chunk data: %v", err))
		}
	}
	log.Printf("End Image Download, Name: %s", imageName)
	imageID, err := server.imageStore.Save(imageName, imageType, imageData)
	if err != nil {
		return logError(status.Errorf(codes.Internal, "cannot save image to the store: %v", err))
	}

	res := &pb.UploadImageResponse{
		Id:   imageID,
		Size: uint32(imageSize),
	}

	err = stream.SendAndClose(res)
	if err != nil {
		return logError(status.Errorf(codes.Unknown, "cannot send response: %v", err))
	}

	log.Printf("saved image with id: %s, size: %d", imageID, imageSize)

	return nil
}

func contextError(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		return logError(status.Error(codes.Canceled, "request is canceled"))
	case context.DeadlineExceeded:
		return logError(status.Error(codes.DeadlineExceeded, "deadline is exceeded"))
	default:
		return nil
	}
}

func logError(err error) error {
	if err != nil {
		log.Print(err)
	}
	return err
}
