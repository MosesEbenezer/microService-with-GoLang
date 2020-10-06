FROM golang:latest
# line above specifies the image that we want to build our own image off of

# then we start setting up that image
RUN mkdir /app
# first make the work directory in above line - where our code will live

ADD . /app
# we then copy all our project contant "." to that newly created app directory

WORKDIR /app
# we then specify the work directory as app - like above

RUN go build -o main .
# then we run the build command passing in the current directory
# this builds our binary excutable

CMD ["/app/main"]
# we finally check that binary executable with the command above

# from this, we can then create our image from this in the command line
# with the command - $ docker build -t image_name .    ... the "." specifies the directory you want to build it 

# and then also run it locally within a container to be viewed on the localhost with port mapping
# - $ docker run -it -p 8080:80 image_name