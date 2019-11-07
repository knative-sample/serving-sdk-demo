#!/bin/bash
#****************************************************************#
# Create Date: 2019-02-02 22:16
#********************************* ******************************#

ROOTDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

NAME="serving-controller"
NAMESPACE="knative-sample"
GIT_COMMIT="$(git rev-parse --verify HEAD)"
GIT_BRANCH=`git branch | grep \* | cut -d ' ' -f2`
TAG="${GIT_BRANCH}_${GIT_COMMIT:0:8}-$(date +%Y''%m''%d''%H''%M''%S)"

docker build -t "${NAME}:${TAG}" -f ${ROOTDIR}/Dockerfile ${ROOTDIR}/../

array=( registry.cn-hangzhou.aliyuncs.com )
for registry in "${array[@]}"
do
    echo "push images to ${registry}/${NAMESPACE}/${NAME}:${TAG}"
    docker tag "${NAME}:${TAG}" "${registry}/${NAMESPACE}/${NAME}:${TAG}"
    docker push "${registry}/${NAMESPACE}/${NAME}:${TAG}"

    docker tag "${NAME}:${TAG}" "${registry}/${NAMESPACE}/${NAME}:latest"
    docker push "${registry}/${NAMESPACE}/${NAME}:latest"
done
