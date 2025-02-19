
# Image Processing Microservice

This project provides a microservice that processes images by cropping them into a circular shape using a gRPC server and a Go Fiber client. The service is divided into two main parts:

1. **Go Client** - This component is responsible for sending image processing requests to the Python server.
2. **Python Server** - This component receives the image processing requests, processes the image, and sends back the result.

---

## Project Structure

```
project-root/
├── client/
│   └── main.go          # Go Fiber client that communicates with the gRPC server
├── server/
│   └── main.py          # Python gRPC server that processes images
├── service/
│   └── service.proto    # Protocol Buffers (Proto) definition for the gRPC service
├── gen/                 # Generated code from service.proto for Python and Go
├── example/
│   └── main.py          # Python script to send requests to the Go Fiber server
├── README.md            # This file
```

---

## Requirements

- **Go (client)**:
  - Go version 1.18 or higher.
  - Go Fiber package.
  - gRPC Go package.

- **Python (server)**:
  - Python 3.x.
  - `grpcio`, `grpcio-tools` for gRPC.
  - OpenCV (`cv2`) for image processing.
  - NumPy for image manipulation.

---

## Installation

### Python Setup

1. **Install required Python packages**:
   ```
   pip install grpcio grpcio-tools opencv-python numpy requests
   ```

2. **Generate Python gRPC code**:
   Use `grpcio-tools` to generate the Python server code from the `proto` file:
   ```
   python3 -m grpc_tools.protoc -I. --python_out=./gen --grpc_python_out=./gen service/service.proto
   ```

### Go Setup

1. **Install required Go packages**:
   ```
   go get github.com/gofiber/fiber/v2
   go get google.golang.org/grpc
   ```

2. **Generate Go gRPC code**:
   Use the following command to generate Go code from the `proto` file:
   ```
   protoc --go_out=gen --go-grpc_out=gen service/service.proto
   ```

---

## Running the Project

### 1. Start the Python Server

To start the gRPC server that processes images, navigate to the `server/` directory and run the following command:

```
python3 main.py
```

The server will listen on `localhost:50051`.

### 2. Start the Go Client

To start the Go Fiber client that sends image processing requests, navigate to the `client/` directory and run:

```
go run main.go
```

The client will listen on `localhost:3000`.

### 3. Send an Image Processing Request

To send an image for processing, use the `example/main.py` script by providing the path to an image file:

```
python3 example/main.py
```

This script sends an image with a specified size to the Go server, which forwards the request to the Python server to process the image.

---

## How It Works

### Go Client (client/main.go)

1. **Client Setup**: The Go client connects to the gRPC server on `localhost:50051`.
2. **Request Handling**: When the Go client receives a POST request at `/process`, it parses the request body (which includes the image path and size) and sends it to the Python server over gRPC.
3. **Response Handling**: The Go client receives the processed image response (status, path, size) and returns this information to the user.

### Python Server (server/main.py)

1. **Server Setup**: The Python server listens for incoming gRPC requests on `localhost:50051`.
2. **Image Processing**: Upon receiving an image processing request, the server uses OpenCV to crop the image into a circular shape and saves the output.
3. **Response**: The server returns the status, output image path, and size back to the client.

---

## Example Request

When you send an image through the Go client:

```json
{
  "path": "/absolute/path/to/image.jpg",
  "size": 200
}
```

The Python server will:
- Read the image from the provided path.
- If the size is greater than 0, it will use that size for processing.
- Crop the image into a circular shape.
- Return a response like:

```json
{
  "status": true,
  "path": "/absolute/path/to/image_processed.png",
  "size": 200
}
```

---

## Troubleshooting

### 1. Missing Dependencies
Ensure that all dependencies are installed for both the Python and Go components.

### 2. gRPC Connection Issues
If the Go client cannot connect to the Python server, ensure that the server is running on `localhost:50051` and that there are no firewall issues.

### 3. Image Not Found
If the provided image path is invalid or the image cannot be loaded, the Python server will return an error.

---

## Future Improvements

- **Image Format Support**: Add support for additional image formats (JPEG, GIF, etc.).
- **Error Handling**: Improve error handling for missing or invalid image files.
- **Performance**: Optimize image processing for large files.
