#!/bin/bash

# set -xeu
set -x

_TAG='go-1.8'

docker build --no-cache .

docker tag $(docker images -q | head -1) ${_TAG}
sh ./docker-run.sh ${_TAG}
