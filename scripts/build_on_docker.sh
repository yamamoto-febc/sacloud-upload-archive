#!/bin/bash

set -e

DOCKER_IMAGE_NAME="sacloud-upload-archive-build"
DOCKER_CONTAINER_NAME="sacloud-upload-archive-build-container"

if [[ $(docker ps -a | grep $DOCKER_CONTAINER_NAME) != "" ]]; then
  docker rm -f $DOCKER_CONTAINER_NAME 2>/dev/null
fi

docker build -t $DOCKER_IMAGE_NAME .

docker run --name $DOCKER_CONTAINER_NAME \
       -e SAKURACLOUD_ACCESS_TOKEN \
       -e SAKURACLOUD_ACCESS_TOKEN_SECRET \
       -e SAKURACLOUD_ZONE \
       -e SAKURACLOUD_TRACE_MODE \
       $DOCKER_IMAGE_NAME make "$@"
if [[ "$@" == *"build"* ]]; then
  docker cp $DOCKER_CONTAINER_NAME:/go/src/github.com/yamamoto-febc/sacloud-upload-archive/bin ./
fi
docker rm -f $DOCKER_CONTAINER_NAME 2>/dev/null
