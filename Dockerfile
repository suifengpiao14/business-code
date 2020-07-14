# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.14-alpine AS build
RUN apk add --no-cache gcc musl-dev
RUN wget https://github.com/upx/upx/releases/download/v3.96/upx-3.96-amd64_linux.tar.xz && \
tar -Jxf upx*.tar.xz && \
cp upx*/upx /usr/bin
# Copy the local package files to the container's workspace.
WORKDIR /app
ADD . /app 
# 先get 这样多次编译的时候能使用缓存
RUN go install ...

#编译
RUN cd /app && \ 
go build -trimpath -ldflags "-s -w" -o ./businessCode .   && \
go build -trimpath -ldflags "-s -w" -o ./static/http ./static  


RUN  upx ./businessCode ./static/http

FROM alpine
RUN apk update && \
   apk add ca-certificates && \
      update-ca-certificates && \
         rm -rf /var/cache/apk/*

	 WORKDIR /app
	 COPY --from=build /app/businessCode  /app/
	 COPY --from=build /app/static/http /app/static/ 
	 COPY --from=build /app/static/dist  /app/static/dist/
	 COPY --from=build /app/entrypoint.sh  /app/

	 ENTRYPOINT /app/entrypoint.sh
	 EXPOSE 8095 8096
