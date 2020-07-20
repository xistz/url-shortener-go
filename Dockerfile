FROM golang:1.14-alpine as development

WORKDIR /app

COPY . .