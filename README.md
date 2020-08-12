# warpin-messaging

A simple RESTful api for sending messages

## Prerequisite

Before you can use the app, you have to build it first. Run the following command in the project root directory. 
```bash
bin/setup
```
It will create a binary file required to run the app. After the built process succeeds, it will run a unit test to make sure all functionality runs as expected. 

## Endpoints
All endpoints are defined in ```spec.yaml``` file and can be imported into ```Postman``` collection. Or you can access the [Documentation](https://app.swaggerhub.com/apis-docs/archisdi/warpin-messaging/1.0).

## Dockerfile
You can get production ready docker image from [dockerhub](https://hub.docker.com/repository/docker/archisdi/warpin-messaging).
```bash
docker pull archisdi/warpin-messaging
```

## Test
To run the unit test, run the following command in the project root directory. 
```bash
bin/test
```