# open**encoder**
> Open Source Cloud Encoder

⚠️*Currently a work-in-progress!*

* HTTP API for submitting jobs to an FFmpeg worker
* Redis-backed worker
* S3-based storage
* Web Dashboard UI for managing encode jobs

https://godoc.org/github.com/alfg/enc

[![Build Status](https://travis-ci.org/alfg/openencoder.svg?branch=master)](https://travis-ci.org/alfg/openencoder) 
[![GoDoc](https://godoc.org/github.com/alfg/openencoder?status.svg)](https://godoc.org/github.com/alfg/openencoder)
[![Go Report Card](https://goreportcard.com/badge/github.com/alfg/openencoder)](https://goreportcard.com/report/github.com/alfg/openencoder)

## Develop
#### Requirements
* Docker
* Go 1.11+
* FFmpeg
* Postgres
* AWS S3 Credentials & Bucket

#### Setup
* Start Redis and Postgres in Docker:
```
docker-compose up -d redis
docker-compose up -d db
```

* Create DB and run `scripts/schema.sql` to set up schema.

* Set environment variables in `docker-compose.yml`:
```
AWS_S3_BUCKET=
AWS_S3_REGION=
AWS_ACCESS_KEY=
AWS_SECRET_KEY
...
```

*Environment variables will override defaults set in `config/default.yml`.*

* Build & start API server:
```
go build -v && openencoder.exe server
```

* Build & start worker:
```
go build -v && openencoder.exe worker
```

* Start Web Dashboard for development:
```
cd static && npm run serve
```

## Usage
```bash
curl -X POST \
  http://localhost:8080/api/encode \
  -H 'Content-Type: application/json' \
  -d '{
	"profile": "h264_baseline_360p_600",
	"source": "s3:///src/ToS-1080p.mp4",
	"dest": "s3:///dst/tears-of-steel/"
  }'
```

## API
See: [API.md](/API.md)

## Scaling
TBD

## TODO
* Distributed chunked encoding
* Encoding profiles API/DB
* Encoding status and health-checks
* Machine scaling
* Digital Ocean S3-Compatibility

## License
MIT