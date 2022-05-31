#!/bin/bash

# start msyql container with docker
start() {
    docker run -p 6666:3306 -it -d --name=go_reserver -e MYSQL_ROOT_PASSWORD=123456 mysql:5.7.19
}

# start
start