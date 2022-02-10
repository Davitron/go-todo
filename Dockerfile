FROM 1.17.3-alpine
WORKDIR /app
ARG PORT
COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go install github.com/cespare/reflex@latest

EXPOSE $PORT

CMD reflex -g '*.go' go run main.go --start-service
