# Use an official Python runtime as a parent image
FROM python:3.8-slim

# Set the working directory in the container
WORKDIR /app

# Copy the generated protobuf files into the container
COPY image_search_pb2.py image_search_pb2_grpc.py /app/
COPY templates/ ./templates/ 

# Copy the Flask app code and any dependencies
COPY app.py requirements.txt /app/

# Install any needed packages specified in requirements.txt
RUN pip install --trusted-host pypi.python.org -r requirements.txt

# Specify the command to run the Flask app
CMD ["python", "app.py"]
