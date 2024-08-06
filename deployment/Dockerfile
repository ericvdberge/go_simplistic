FROM golang:alpine as build
WORKDIR /app

COPY . .
RUN go mod download
WORKDIR /app/cmd
RUN go build -o ./dist

FROM scratch as final
EXPOSE 8080
COPY .env .
COPY --from=build ./app/cmd/dist /app
COPY --from=build ./app/internal/web internal/web
CMD ["/app"]




