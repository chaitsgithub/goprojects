NGINX Server
docker rm nginx-server --force
docker build -t nginx-server .
docker run --name nginx-server -p 80:80 -d nginx-server

Go Server
docker rm go-webserver --force
docker build -t go-webserver .
docker run --name go-webserver -p 8080:8080 -d go-webserver
