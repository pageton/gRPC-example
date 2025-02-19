package main

import (
	"context"
	"flag"
	"log"
	pb "microservice/gen"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"
)

type gRPCClient struct {
	client pb.ImageProcessorClient
}

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewImageProcessorClient(conn)

	app := fiber.New()

	app.Post("/process", func(c *fiber.Ctx) error {
		var request struct {
			Path string `json:"path"`
			Size int32  `json:"size"`
		}

		if err := c.BodyParser(&request); err != nil {
			log.Printf("could not parse request body: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status": false,
				"error":  "Invalid request body",
			})
		}

		req := &pb.Request{
			Path: request.Path,
			Size: request.Size,
		}

		resp, err := client.ProcessImage(context.Background(), req)
		if err != nil {
			log.Printf("could not process image: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status": false,
				"error":  "Invalid request body",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": resp.Status,
			"path":   resp.Path,
			"size":   int32(resp.Size),
		})
	})

	log.Fatal(app.Listen(":3000"))
}
