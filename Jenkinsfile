def branch
def revision
def registryIp
def buildName = "build-govolutto"
def imageName = "maslick/govolutto-unzipped"
def registryName = "docker.io"

pipeline {
    agent {
        kubernetes {
            label "${buildName}-pod"
            defaultContainer 'jnlp'
            yaml """
apiVersion: v1
kind: Pod
metadata:
  labels:
    job: build-service
spec:
  containers:
  - name: go
    image: golang:1.13.4
    command: ["cat"]
    tty: true
  - name: docker
    image: docker:18.09.2
    command: ["cat"]
    tty: true
    volumeMounts:
    - name: docker-sock
      mountPath: /var/run/docker.sock
  volumes:
  - name: docker-sock
    hostPath:
      path: /var/run/docker.sock
"""
        }
    }
    options {
        skipDefaultCheckout true
    }

    stages {
        stage ('checkout') {
            steps {
                script {
                    def repo = checkout scm
                    revision = sh(script: 'git log -1 --format=\'%h.%ad\' --date=format:%Y%m%d-%H%M | cat', returnStdout: true).trim()
                    branch = repo.GIT_BRANCH.take(20).replaceAll('/', '_')
                    if (branch != 'master') {
                        revision += "-${branch}"
                    }
                    sh "echo 'Building revision: ${revision}'"
                }
            }
        }
        stage ('build') {
            steps {
                container('go') {
                    sh "go mod download"
                    sh "go get github.com/google/wire/cmd/wire"
                    sh "wire ./src"
                    sh "go test ./test -coverpkg=./... -coverprofile=coverage.txt -covermode=atomic -race -tags test"
                    sh "CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o build/govolutto.zip"
                }
            }
        }
        stage ('dockerize') {
            steps {
                container('docker') {
                    sh "docker build . -t ${imageName}:${revision} --build-arg REVISION=${revision}"
                }
            }
        }
        stage ('publish artifact') {
            when {
                expression {
                    branch == 'di'
                }
            }
            steps {
                container('docker') {
                    withCredentials([usernamePassword(credentialsId: 'dockerhub_registry', usernameVariable: 'DOCKER_USERNAME', passwordVariable: 'DOCKER_PASSWORD')]){
                        sh "docker tag ${imageName}:${revision} ${registryName}/${imageName}:latest"
                        sh "docker tag ${imageName}:${revision} ${registryName}/${imageName}:${revision}"

                        sh "docker login ${registryName} --username $DOCKER_USERNAME --password $DOCKER_PASSWORD"
                        sh "docker push ${registryName}/${imageName}:latest"
                        sh "docker push ${registryName}/${imageName}:${revision}"
                    }
                }
            }
        }
    }
}
