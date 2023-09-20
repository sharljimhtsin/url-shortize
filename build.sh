  301  CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o xzx_im
  304  CGO_ENABLED=1 CGO_CFLAGS="-D_LARGEFILE64_SOURCE" CC=musl-gcc go build --ldflags '-linkmode=external -extldflags=-static' -o xzx_im_static
