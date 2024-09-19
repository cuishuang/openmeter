package main

import (
	"fmt"

	"github.com/openmeterio/openmeter/.dagger/internal/dagger"
)

func (m *Openmeter) Etoe(
	// +optional
	test string,
) *dagger.Container {
	image := m.Build().ContainerImage("").
		WithExposedPort(10000).
		WithMountedFile("/etc/openmeter/config.yaml", m.Source.File("e2e/config.yaml")).
		WithServiceBinding("kafka", dag.Kafka(dagger.KafkaOpts{Version: kafkaVersion}).SingleNode().Service()).
		WithServiceBinding("clickhouse", clickhouse())

	postgres := dag.Postgres(dagger.PostgresOpts{
		Version: postgresVersion,
	})

	api := image.
		WithExposedPort(8080).
		WithServiceBinding("postgres", postgres.Service()).
		WithEnvVariable("POSTGRES_HOST", "postgres").
		WithExec([]string{"openmeter", "--config", "/etc/openmeter/config.yaml"}).
		AsService()

	sinkWorker := image.
		WithServiceBinding("redis", redis()).
		WithServiceBinding("api", api). // Make sure api is up before starting sink worker
		WithExec([]string{"openmeter-sink-worker", "--config", "/etc/openmeter/config.yaml"}).
		AsService()

	args := []string{"go", "test", "-tags", "musl", "-count=1", "-v"}

	if test != "" {
		args = append(args, "-run", fmt.Sprintf("Test%s", test))
	}

	args = append(args, "./e2e/...")

	return goModuleCross("").
		WithModuleCache(cacheVolume("go-mod-e2e")).
		WithBuildCache(cacheVolume("go-build-e2e")).
		WithSource(m.Source).
		WithEnvVariable("OPENMETER_ADDRESS", "http://api:8080").
		WithEnvVariable("TEST_WAIT_ON_START", "true").
		WithServiceBinding("api", api).
		WithServiceBinding("sink-worker", sinkWorker).
		Exec(args)
}

func clickhouse() *dagger.Service {
	return dag.Container().
		From(fmt.Sprintf("clickhouse/clickhouse-server:%s-alpine", clickhouseVersion)).
		WithEnvVariable("CLICKHOUSE_USER", "default").
		WithEnvVariable("CLICKHOUSE_PASSWORD", "default").
		WithEnvVariable("CLICKHOUSE_DB", "openmeter").
		WithExposedPort(9000).
		WithExposedPort(9009).
		WithExposedPort(8123).
		AsService()
}

func redis() *dagger.Service {
	return dag.Container().
		From(fmt.Sprintf("redis:%s-alpine", redisVersion)).
		WithExposedPort(6379).
		AsService()
}

const (
	SvixJWTSingingSecret = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE3MjI5NzYyNzMsImV4cCI6MjAzODMzNjI3MywibmJmIjoxNzIyOTc2MjczLCJpc3MiOiJzdml4LXNlcnZlciIsInN1YiI6Im9yZ18yM3JiOFlkR3FNVDBxSXpwZ0d3ZFhmSGlyTXUifQ.PomP6JWRI62W5N4GtNdJm2h635Q5F54eij0J3BU-_Ds"
)
