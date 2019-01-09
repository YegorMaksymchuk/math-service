#!/usr/bin/env bash
docker build -f Dockerfile.multi-stage -t math-service:latest .
docker run -d -it -p6000:6000 --name=math-service math-service:latest