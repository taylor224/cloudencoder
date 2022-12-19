<div align="center">
    <h1><code>OpenEncoder</code></h1>
    <p><strong>Open Source Cloud Encoder for FFmpeg</strong></p>
    <p>NVidia CUDA GPU 를 사용할수 있도록 개조된 OpenEncoder 입니다.
    </p>
    <p>⚠️ NVidia CUDA 를 사용하지 않는 경우 의도한 대로 작동이 되지 않을수 있습니다.</p>
    <p>
        <a href="https://travis-ci.org/alfg/openencoder">
          <img src="https://travis-ci.org/alfg/openencoder.svg?branch=master" alt="Build Status" />
        </a>
        <a href="https://godoc.org/github.com/alfg/openencoder">
          <img src="https://godoc.org/github.com/alfg/openencoder?status.svg" alt="GoDoc" />
        </a>
        <a href="https://goreportcard.com/report/github.com/alfg/openencoder">
          <img src="https://goreportcard.com/badge/github.com/alfg/openencoder" alt="Go Report Card" />
        </a>
        <a href="https://hub.docker.com/r/alfg/openencoder/builds">
          <img src="https://img.shields.io/docker/automated/alfg/openencoder.svg" alt="Docker Automated build" />
        </a>
        <a href="https://hub.docker.com/r/alfg/openencoder">
          <img src="https://img.shields.io/docker/pulls/alfg/openencoder.svg" alt="Docker Pulls" />
        </a>
    </p>
</div>

## cloud_encoder.py 사용방법
<p>OpenEncoder 를 ffmpeg binary 처럼 사용 가능하게 해주는 Python Script</p>
<p>본 스크립트는 인코딩을 요청하는 웹서버 등에서 사용하는 스크립트 입니다.</p>
<p>사전에 상단의 설정 값들을 변경해주시기 바랍니다.</p>
<p>alias ffmpeg = "python3 cloud_encoder.py"</p>
<p>하셔서 사용하시면 됩니다.</p>
<p>단 parameter 는 input 파일 path와 output 파일 path만 받아서 OpenEncoder 에서 미리 처리된 방식으로 인코딩 하니 주의하시기 바랍니다.</p>
<p>⚠️ 현재 인코딩 된 파일을 다운로드 하는 부분이 단일 워커에서만 사용 가능하도록 구현되어 있습니다. 차후 여러 워커에서 사용할수 있도록 여러 워커에서 파일을 체크해서 파일이 있는 워커에서 다운받을수 있도록 구현할 예정입니다.</p>

## Features
* HTTP API for submitting jobs to a redis-backed FFmpeg worker
* Local save and FTP and S3 storage (AWS, Digital Ocean Spaces and Custom S3 Providers supported)
* Web Dashboard UI for managing encode jobs, workers, users and settings
* Machines UI/API for scaling cloud worker instances in a VPC
* Database stored FFmpeg encoding presets
* User accounts and roles


## Preview
![Screenshot](screenshot.png)    


## Development

#### Requirements
* Docker
* Go 1.11+
* NodeJS 8+ (For web dashboard)
* FFmpeg
* Postgres
* S3 API Credentials & Bucket (AWS or Digital Ocean)
* Digital Ocean API Key (only required for Machines API)

Docker is optional, but highly recommended for this setup. This guide assumes you are using Docker.


#### Setup
* Start Redis and Postgres in Docker:
```
docker-compose up -d redis
docker-compose up -d db
```

When the database container runs for the first time, it will create a persistent volume as `/var/lib/postgresql/data`. It will also run the scripts in `scripts/` to create the database, schema, settings, presets, and an admin user.

* Build & start API server:
```
go build -v && ./openencoder server
```

* Start the worker:
```
./openencoder worker
```

* Start Web Dashboard for development:
```
cd static && npm run serve
```

* Open `http://localhost:8081/dashboard` in the browser and login with `admin/password`.


See [Quick-Setup-Guide](https://github.com/alfg/openencoder/wiki/Quick-Setup-Guide-%5Bfor-development%5D) for full development setup guide.

## API
See: [API.md](/API.md)


## Scaling
You can scale workers by adding more machines via the Web UI or API.

Currently only `Digital Ocean` is supported. More providers are planned.

See: [API.md](/API.md) for Machines API documentation.


## Documentation
See: [wiki](https://github.com/alfg/openencoder/wiki) for more documentation.


## Roadmap
See: [Development Project](https://github.com/alfg/openencoder/projects/1) for current development tasks and status.


## License
MIT
