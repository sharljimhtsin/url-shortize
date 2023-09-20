# syntax=docker/dockerfile:1

FROM golang:1.21.1-alpine3.18
WORKDIR /app
COPY . .
CMD ["./xzx_im_static"]
EXPOSE 8081
