#!/usr/bin/env bash

set -e

readonly VERSION=${VERSION}
readonly PULSARCTL_HOME=${PULSARCTL_HOME:-"/pulsarctl"}
pushd ${PULSARCTL_HOME}
readonly DIR_PREFIX="pulsarctl"

arch=(amd64 386)
os=(darwin linux windows)
pushd ${PULSARCTL_HOME}
for a in ${arch[@]}; do
    for o in ${os[@]}; do
        if [[ ${o} == "darwin" && ${a} == "386" ]]; then
            continue
        fi
        echo "Building package with os ${o} and arch ${a}"
        dir=${DIR_PREFIX}-${a}-${o}
        if [[ -d ${dir} ]]; then
            rm -rf ${dir}
        fi
        if [[ -f ${dir}.tar.gz ]]; then
            rm ${dir}.tar.gz
        fi
        mkdir -p ${dir}
        CGO_ENABLED=0 GOOS=${o} GOARCH=${a} go build -o ${dir}/pulsarctl \
            -ldflags "-X github.com/streamnative/pulsarctl/pkg/pulsar.ReleaseVersion=Pulsarctl-Go-${VERSION}" .
        cp -r plugins ${dir}
        tar -czf ${dir}.tar.gz ${dir}
        rm -rf ${dir}
    done
done
if [[ ! -d "/pulsarctl/release" ]]; then
    mkdir -p /pulsarctl/release
fi
mv *.tar.gz /pulsarctl/release
popd
