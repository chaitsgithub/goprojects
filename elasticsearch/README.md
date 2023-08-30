https://developer.okta.com/blog/2021/04/23/elasticsearch-go-developers-guide

Docker Commands:
Pull Docker Image: docker pull docker.elastic.co/elasticsearch/elasticsearch:7.5.2
Create a New Volume so data is not lost: docker volume create elasticsearch

docker rm -f elasticsearch
docker run -d --name elasticsearch -p 9200:9200 -e discovery.type=single-node \
    docker.elastic.co/elasticsearch/elasticsearch:7.5.2
docker ps

Data:
http://stapi.co/api/v1/rest/spacecraft/search?pageNumber=0&pageSize=100&pretty

** Use go get github.com/elastic/go-elasticsearch/v7 istead of v8 as elasticsearch is v8