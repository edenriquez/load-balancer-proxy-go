FROM golang:1.12

WORKDIR $GOPATH/src/github.com/edenriquez/load-balancer-proxy-go

# Copy the content of your repository into the image
COPY . ./
RUN go get github.com/ziutek/mymysql 
RUN go get github.com/ziutek/mymysql/native
RUN go get github.com/joho/godotenv