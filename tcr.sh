#!/bin/bash
go test ./... && git commit -am working || git reset --hard