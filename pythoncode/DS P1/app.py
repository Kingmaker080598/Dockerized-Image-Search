from flask import Flask, render_template, request, send_file, Response
import grpc
import image_search_pb2
import image_search_pb2_grpc
import base64


app = Flask(__name__)

@app.route("/", methods=["GET", "POST"])
def index():
    image_data = None
    error_message = None

    if request.method == "POST":
        keyword = request.form["keyword"]
        image_data, error_message = search_image(keyword)

        # Convert image_data to base64
        image_base64 = base64.b64encode(image_data).decode() if image_data else None

        return render_template("result.html", image_base64=image_base64, error_message=error_message)
    
    return render_template("index.html")


def search_image(keyword):
    try:
        channel = grpc.insecure_channel('server:50051')
        stub = image_search_pb2_grpc.ImageSearchServiceStub(channel)
        request = image_search_pb2.SearchRequest(keyword=keyword)
        response = stub.SearchImage(request)

        if response.image_data:
            print('Image Found')
            return response.image_data, None
        # else:
        #     print("No images found for the keyword.")
        #     return None, "No images found for the keyword."
    except grpc.RpcError as e:
        return None, str(e)


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=80)
