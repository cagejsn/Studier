#!/bin/bash

yarn --cwd client build
yarn --cwd client container

docker push cagejsn/client.studier
echo built client

CURRENT_WORKING_DIRECTORY=`pwd`
return_to_start_directory() {
cd $CURRENT_WORKING_DIRECTORY
}

cd problem-service
docker build . -t cagejsn/problem-service.studier || return_to_start_directory
docker push cagejsn/problem-service.studier || return_to_start_directory
cd $CURRENT_WORKING_DIRECTORY

cd user-interaction-service
docker build . -t cagejsn/user-interaction-service.studier || return_to_start_directory
docker push cagejsn/user-interaction-service.studier || return_to_start_directory
cd $CURRENT_WORKING_DIRECTORY