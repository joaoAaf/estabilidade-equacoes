FROM golang:alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o ./api
FROM gcr.io/distroless/static-debian12
COPY --from=build /app/api /usr/local/bin/api
EXPOSE 8080
USER nonroot:nonroot
CMD [ "api" ]