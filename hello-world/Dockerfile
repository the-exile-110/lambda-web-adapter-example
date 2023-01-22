FROM arm64v8/golang:1.19.1-bullseye as build-env

WORKDIR /go/src/app

COPY . .

RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o /go/bin/app

FROM gcr.io/distroless/static:latest-arm64

COPY --from=build-env /go/bin/app /
COPY --from=public.ecr.aws/awsguru/aws-lambda-adapter:0.6.0 /lambda-adapter /opt/extensions/lambda-adapter

ENV PORT=3000
EXPOSE 3000

CMD ["/app"]