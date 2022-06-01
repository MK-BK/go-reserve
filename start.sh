#!/bin/bash

build_web() {

}

build_server() {

}

start() {
    docker run -p 6666:3306 -it -d --name=go_reserver -e MYSQL_ROOT_PASSWORD=123456 mysql:5.7.19

    build_web()
    build_server()
}

start