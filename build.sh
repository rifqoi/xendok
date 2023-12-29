#!/bin/bash
#
VERSION="v0.0.1"
go build -ldflags="-X github.com/rifqoi/xendok-service/internal/logger.appVersion=${VERSION}" -o main .
