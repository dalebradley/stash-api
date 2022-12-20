FROM golang:1.16-alpine as base

FROM base as dev

#RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /app/stash-api

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

# Install the air binary so we get live code-reloading when we save files
RUN go get github.com/cosmtrek/air


COPY . .

# Build the Go app
RUN go build -o ./out/stash-api .


# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["air","./out/stash-api"]