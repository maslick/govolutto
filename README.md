# =govolutto=
Golang version of money transfer REST API ([see link](https://github.com/maslick/revolutto))

[![Build Status](https://travis-ci.org/maslick/govolutto.svg?branch=master)](https://travis-ci.org/maslick/govolutto)
[![image size](https://img.shields.io/badge/image%20size-4MB-blue.svg)](https://cloud.docker.com/u/maslick/repository/docker/maslick/govolutto)
[![Maintainability](https://api.codeclimate.com/v1/badges/e189c55d25e618f34704/maintainability)](https://codeclimate.com/github/maslick/govolutto/maintainability)
[![codecov](https://codecov.io/gh/maslick/govolutto/branch/master/graph/badge.svg)](https://codecov.io/gh/maslick/govolutto)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)

## Features
* Written in Go :heart:
* Lightweight executable: ~4.5MB
* Production ready: Dockerfile + k8s yaml

## Installation
```zsh
go test -v test/*
go build -ldflags="-s -w" -o govolutto
go build -ldflags="-s -w" -o govolutto.zip && upx govolutto.zip
```

## API
* Get balance: ``GET v1/{username}/balance``
* Transfer money: ``POST v1/transfer``
* Health endpoint: ``GET /health``

```json
{
  "from": "scrooge",
  "to": "daisy",
  "amount": 100.0
}
```

## Usage
```zsh
./govolutto
./govolutto.zip

http :8080/v1/daisy/balance | jq
{
  "balance": "100",
  "username": "daisy"
}

http :8080/v1/scrooge/balance | jq
{
  "balance": "10000",
  "username": "scrooge"
}

http POST :8080/v1/transfer <<< '{"from": "scrooge", "to": "daisy", "amount": 10000.0}' | jq
{
  "amount": "10000",
  "from": "scrooge",
  "success": "true",
  "to": "daisy"
}

http POST :8080/v1/transfer <<< '{"from": "daisy", "to": "scrooge", "amount": 10000.0}' | jq
{
  "amount": "10000",
  "from": "daisy",
  "success": "true",
  "to": "scrooge"
}
```

## Load test
```zsh
echo "POST http://localhost:8080/v1/transfer" | vegeta attack -body payload.json -header="Content-Type: application/json" -rate=500 -duration=5s | tee results.bin | vegeta report
Requests      [total, rate, throughput]  2500, 500.20, 500.18
Duration      [total, attack, wait]      4.998199772s, 4.997979158s, 220.614µs
Latencies     [mean, 50, 95, 99, max]    387.021µs, 212.602µs, 372.705µs, 5.781057ms, 28.793267ms
Bytes In      [total, mean]              157500, 63.00
Bytes Out     [total, mean]              142500, 57.00
Success       [ratio]                    100.00%
Status Codes  [code:count]               200:2500
Error Set:

cat results.bin | vegeta report -type="hist[0,1ms,5ms,10ms,20ms,50ms,100ms,500ms,1000ms]"
Bucket           #     %       Histogram
[0s,     1ms]    2460  98.40%  #########################################################################
[1ms,    5ms]    13    0.52%
[5ms,    10ms]   10    0.40%
[10ms,   20ms]   11    0.44%
[20ms,   50ms]   6     0.24%
[50ms,   100ms]  0     0.00%
[100ms,  500ms]  0     0.00%
[500ms,  1s]     0     0.00%
[1s,     +Inf]   0     0.00%
```

## Docker
* Lightweight Docker image (4.5MB)
* See [Dockerfile](Dockerfile)
```zsh
GOOS=linux go build -ldflags="-s -w" -o build/govolutto.zip && upx build/govolutto.zip
docker build -t maslick/govolutto .
docker run -d -p 8081:8080 maslick/govolutto

http http://`docker-machine ip default`:8081/v1/daisy/balance | jq
{
  "balance": "100",
  "username": "daisy"
}
```

## k8s
```zsh
kubectl apply -f k8s/deployment.yaml
kubectl get all -l project=govolutto
kubectl port-forward govolutto-api-5b58b69647-877qd 8083:8080
http :8083/health
```

## Links
* [upx](https://upx.github.io/)
* [httpie](https://httpie.org/)
* [vegeta](https://github.com/tsenart/vegeta)
* [jq](https://stedolan.github.io/jq/)
