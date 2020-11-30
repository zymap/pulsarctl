#!/usr/bin/env bash
set -e

readonly PROJECT_ROOT=`cd $(dirname $0)/..; pwd`
readonly IMAGE_NAME=pulsarctl-release

docker build -t ${IMAGE_NAME} \
             -f ${PROJECT_ROOT}/scripts/release/Dockerfile ${PROJECT_ROOT}

docker run -v ${PROJECT_ROOT}:/pulsarctl ${IMAGE_NAME} bash -c "/pulsarctl/scripts/release/entrypoint.sh"
