# we take the image of Python in the version 3.8.6
FROM python:3.8.6

# we update pip
RUN pip install --upgrade pip

# we set the directory to use inside container
WORKDIR /api

# we copy the files of our api inside actual folder and send inside container folder
COPY . /api/

# we install all necessary dependencies
RUN pip install pymongo grpcio grpcio-tools

# we open the port 4000
EXPOSE 4000

# we run the command inside the container to run the api
CMD [ "python", "app.py" ]