import os

import requests


def send_image_request(image_path, size):
    url = 'http://localhost:3000/process'

    full_path = os.path.abspath(image_path)

    with open(full_path, 'rb') as img_file:
        files = {'file': img_file}
        data = {'size': size, 'path': full_path}

        response = requests.post(url, files=files, data=data)

    if response.status_code == 200:
        print(f"Status: {response.json()['status']}")
        print(f"Path: {response.json()['path']}")
        print(f"Size: {response.json()['size']}")
    else:
        print(f"Failed to process image. Status code: {response.status_code}")

if __name__ == '__main__':
    send_image_request("image.jpg", 200)
