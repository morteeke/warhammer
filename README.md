# Warhammer

[![Go Reference](https://pkg.go.dev/badge/github.com/brittonhayes/warhammer.svg)](https://pkg.go.dev/github.com/brittonhayes/warhammer)
[![Go Report Card](https://goreportcard.com/badge/github.com/brittonhayes/warhammer)](https://goreportcard.com/report/github.com/brittonhayes/warhammer)

> A Warhammer REST API built with Go and Sqlite.

## ⚡ Quick Start

Get started with docker-compose to seed the database and start the API server.

```sh
docker-compose up --build --force-recreate --remove-orphans --detach
```

or with [task](https://taskfile.dev/)

```sh
task up
```

## 📖 Documentation

API documentation is available in YAML format within the repository. The OpenAPI spec is used to generate the transport logic thanks to [goapi-gen](https://github.com/discord-gophers/goapi-gen).

- [View the OpenAPI Spec](https://github.com/brittonhayes/warhammer/blob/main/api/openapi.yaml)

## 🗄️ Adding or Editing Data

All application data used to seed the database is editable in the [fixtures/](https://github.com/brittonhayes/warhammer/blob/main/fixtures/) directory. The API is built to be self-seeding and read-only, so if there is a need to add more data, simply add it to the fixtures files, rebuild and redeploy the API.

- [View the data](https://github.com/brittonhayes/warhammer/blob/main/fixtures/fixtures.yaml)

## 📈 Monitoring

Application observability is instrumented with OpenTelemetry and available in Grafana, Prometheus, and Jaeger. All application services are behind a traefik reverse proxy.

When running with docker-compose, the following services are available:

- API is available at [http://warhammer.localhost](http://warhammer.localhost)
- Documentation is available at [http://warhammer.localhost/docs](http://warhammer.localhost/docs)
- Traefik is available at [http://localhost:8080](http://localhost:8080)
- Prometheus is available at [http://prometheus.localhost](http://prometheus.localhost)
- Grafana is available at [http://grafana.localhost](http://grafana.localhost)
- Jaeger is available at [http://jaeger.localhost](http://jaeger.localhost)
