FROM golang
COPY main.go /
CMD go run /main.go
RUN go get "github.com/microcosm-cc/bluemonday"
