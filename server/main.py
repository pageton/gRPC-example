import os
from concurrent import futures

import cv2
import grpc
import numpy as np
import service_pb2
import service_pb2_grpc


class ImageProcessorServicer(service_pb2_grpc.ImageProcessorServicer):
    def ProcessImage(self, request, context):
        try:
            image_path = request.path
            size = request.size if request.size > 0 else 256
            if not os.path.exists(image_path):
                print(f"Image not found: {image_path}")
                return service_pb2.Response(status=False, path="", size=0)
            image = cv2.imread(image_path, cv2.IMREAD_UNCHANGED)
            if image is None:
                print(f"Failed to load image: {image_path}")
                return service_pb2.Response(status=False, path=image_path, size=0)

            height, width = image.shape[:2]
            mask = np.zeros((height, width), dtype=np.uint8)
            center = (width // 2, height // 2)
            radius = min(center[0], center[1], width - center[0], height - center[1])
            cv2.circle(mask, center, radius, 255, -1)

            circular_image = cv2.bitwise_and(image, image, mask=mask)
            output_path = f"{os.path.splitext(image_path)[0]}_processed.png"
            cv2.imwrite(output_path, circular_image)

            return service_pb2.Response(status=True, path=output_path, size=size)

        except Exception as e:
            print(f"Error during processing: {e}")
            return service_pb2.Response(status=False, path="", size=0)


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    service_pb2_grpc.add_ImageProcessorServicer_to_server(ImageProcessorServicer(), server)
    port = "50051"
    server.add_insecure_port(f"[::]:{port}")
    print(f"🚀 gRPC Server is running on port {port}...")
    server.start()
    server.wait_for_termination()

if __name__ == "__main__":
    serve()
