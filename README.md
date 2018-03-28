# s3-example
A simple Golang frontend running in a Docker container to manipulate S3 using the AWS CLI.

This container assumes /etc/s3-aws-creds will be a bash file that will export these environment variables:
AWS_ACCESS_KEY_ID
AWS_SECRET_ACCESS_KEY
AWS_DEFAULT_REGION

# Running the Service
```
docker-compose up -d --build
```
OR
```
docker build -t capnchainsaw/s3-example .
docker run -v /local/path/creds:/etc/s3-aws-creds -it capnchainsaw/s3-example
```

# Running Commands Manually

First execute a bash process in the running service container.

```
docker exec -it s3example_frontend_1 /bin/bash
```

## Listing a directory

TODO

## Uploading a File

TODO

## Downloading a File

TODO
