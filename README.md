<h1 align="center">S3 Proxy</h1>

<p align="center">
  <a href="http://godoc.org/github.com/oxyno-zeta/s3-proxy" rel="noopener noreferer" target="_blank"><img src="https://img.shields.io/badge/godoc-reference-blue.svg" alt="Go Doc" /></a>
  <a href="https://github.com/oxyno-zeta/s3-proxy/actions?query=workflow%3AGo" rel="noopener noreferer" target="_blank"><img src="https://github.com/oxyno-zeta/s3-proxy/workflows/Go/badge.svg?branch=master" alt="Go Github Actions" /></a>
  <a href="https://goreportcard.com/report/github.com/oxyno-zeta/s3-proxy" rel="noopener noreferer" target="_blank"><img src="https://goreportcard.com/badge/github.com/oxyno-zeta/s3-proxy" alt="Go Report Card" /></a>
  <a href="https://golangci.com" rel="noopener noreferer" target="_blank"><img src="https://golangci.com/badges/github.com/oxyno-zeta/s3-proxy.svg" alt="GolangCI" /></a>
  <a href="https://coveralls.io/github/oxyno-zeta/s3-proxy?branch=master" rel="noopener noreferer" target="_blank"><img src="https://coveralls.io/repos/github/oxyno-zeta/s3-proxy/badge.svg?branch=master" alt="Coverage Status" /></a>
  <a href="https://hub.docker.com/r/oxynozeta/s3-proxy" rel="noopener noreferer" target="_blank"><img src="https://img.shields.io/docker/pulls/oxynozeta/s3-proxy.svg" alt="Docker Pulls" /></a>
  <a href="https://github.com/oxyno-zeta/s3-proxy/blob/master/LICENSE" rel="noopener noreferer" target="_blank"><img src="https://img.shields.io/github/license/oxyno-zeta/s3-proxy" alt="GitHub license" /></a>
  <a href="https://github.com/oxyno-zeta/s3-proxy/releases" rel="noopener noreferer" target="_blank"><img src="https://img.shields.io/github/v/release/oxyno-zeta/s3-proxy" alt="GitHub release (latest by date)" /></a>
</p>

- [Why ?](#why)
- [Features](#features)
- [Configuration](#configuration)
- [Templates](#templates)
- [API](#api)
  - [GET](#get)
  - [PUT](#put)
  - [DELETE](#delete)
- [AWS IAM Policy](#aws-iam-policy)
- [Grafana Dashboard](#grafana-dashboard)
- [Deployment](#deployment)
  - [Kubernetes - Helm](#kubernetes---helm)
  - [Docker](#docker)
- [TODO](#todo)
- [Inspired by](#inspired-by)
- [Thanks](#thanks)
- [Author](#author)
- [License](#license)

## Why ?

First of all, yes, this is another S3 proxy written in Golang.

I've created this project because I couldn't find any other that allow to proxy multiple S3 buckets or to have custom templates with OpenID Connect authentication and also to get, upload and delete files.

## Features

- Multi S3 bucket proxy
- Index document (display index document instead of listing when found)
- Custom templates
- AWS S3 Login from files or environment variables
- Custom S3 endpoints supported
- Basic Authentication support
- Multiple Basic Authentication support
- OpenID Connect Authentication support
- Multiple OpenID Connect Provider support
- Bucket mount point configuration with hostname and multiple path support
- Authentication by path and http method on each bucket
- Prometheus metrics
- Allow to publish files on S3 bucket
- Allow to delete files on S3 bucket

## Configuration

See here: [Configuration](./docs/configuration.md)

## Templates

See here: [Templates](./docs/templates.md)

## API

### GET

This kind of requests will allow to get files or directory listing.

If path ends with a slash, the backend will consider this as a directory and will perform a directory listing or will display index document.
Example: `GET /dir1/`

If path doesn't end with a slash, the backend will consider this as a file request. Example: `GET /file.pdf`

### PUT

This kind of requests will allow to send file in directory.

The PUT request path must be a directory and must be a multipart form with a key named `file` with a file inside.
Example: `PUT --form file:@file.pdf /dir1/`

### DELETE

This kind of requests will allow to delete files (**only**).

The DELETE request path must contain the file name. Example: `DELETE /dir1/dir2/file.pdf`.

## AWS IAM Policy

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        // Needed for GET API/Action
        "s3:ListBucket",
        "s3:GetObject",
        // Needed for PUT API/Action
        "s3:PutObject",
        "s3:HeadObject",
        // Needed for DELETE API/Action
        "s3:DeleteObject"
      ],
      "Resource": ["arn:aws:s3:::<bucket-name>", "arn:aws:s3:::<bucket-name>/*"]
    }
  ]
}
```

## Grafana Dashboard

This project exports Prometheus metrics. Here is an example of Prometheus dashboard that you can import as JSON file: [dashboard](docs/s3-proxy-dashboard.json).

## Deployment

### Kubernetes - Helm

A helm chart have been created to deploy this in a Kubernetes cluster.

You can find it here: [https://github.com/oxyno-zeta/helm-charts/tree/master/stable/s3-proxy](https://github.com/oxyno-zeta/helm-charts/tree/master/stable/s3-proxy)

### Docker

First, write the configuration file in a config folder. That one will be mounted.

Run this command:

```shell
docker run -d --name s3-proxy -p 8080:8080 -p 9090:9090 -v $PWD/config:/config oxynozeta/s3-proxy
```

## TODO

- Support more authentication and authorization systems
- JSON response
- Add tests

## Inspired by

- [pottava/aws-s3-proxy](https://github.com/pottava/aws-s3-proxy)

## Thanks

- My wife BH to support me doing this

## Author

- Oxyno-zeta (Havrileck Alexandre)

## License

Apache 2.0 (See in LICENSE)
