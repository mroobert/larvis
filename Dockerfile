FROM golang:1.20.2 AS build
WORKDIR /app
COPY . .
WORKDIR /app/cmd
ENV CGO_ENABLED=0
RUN go build -o poker

FROM alpine:3.17.3
RUN addgroup -S appuser && adduser -S appuser -G appuser
WORKDIR /app
COPY --from=build /app/cmd .
RUN chown -R appuser:appuser /app
USER appuser
ENTRYPOINT [ "./poker" ]