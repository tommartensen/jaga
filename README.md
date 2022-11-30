# jaga

**IN DEVELOPMENT**

Jaga (from Swedish: att jaga) is a [Strava](https://www.strava.com) application that finds segments at a location and reports the KOM time and pace to beat.

## Development Setup

As the OAuth flow is not implemented yet, please provide your `ACCESS_TOKEN` as an environment variable to authenticate Strava API requests.

## Requirements

```bash
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```
