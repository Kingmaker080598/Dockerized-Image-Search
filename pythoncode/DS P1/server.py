import grpc
import image_search_pb2
import image_search_pb2_grpc
import random
from concurrent import futures

class ImageSearchServicer(image_search_pb2_grpc.ImageSearchServiceServicer):
    def SearchImage(self, request, context):
        # Implement the logic to search for an image based on the keyword
        # Replace this with your actual image search logic
        images = {
            "dog": ["dog_image1.jpg", "dog_image2.jpg", "dog_image3.jpg"],
            "cat": ["cat_image1.jpg", "cat_image2.jpg", "cat_image3.jpg"],
            "home":["home_image1.jpg", "home_image2.jpg", "home_image3.jpg"],
            "transport": ["car_image1.jpg", "train_image2.jpg", "bus_image3.jpg", "ship_image4.jpg", "aeroplane_image5.jpg"]
        }
        
        keyword = request.keyword
        print(keyword)
        if keyword in images:
            selected_image = random.choice(images[keyword])
            with open('images/'+selected_image, 'rb') as image_file:
                image_data = image_file.read()
                print('Found the image sending the response\n')
                return image_search_pb2.ImageResponse(image_data=image_data)
        else:
            print('Keyword not found')
            # Handle the case when the keyword is not found
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details("Keyword not found")
            return image_search_pb2.ImageResponse()

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    image_search_pb2_grpc.add_ImageSearchServiceServicer_to_server(ImageSearchServicer(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    server.wait_for_termination()
    print('Running Server\n')

if __name__ == '__main__':
    serve()
