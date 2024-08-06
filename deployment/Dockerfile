FROM golang:alpine as build
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN go build -o ./dist

FROM scratch as final
EXPOSE 8080
COPY --from=build /app/dist /app
COPY --from=build /app/static /static
COPY --from=build /app/web /web
CMD ["/app"]




