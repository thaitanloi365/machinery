package machinery_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/thaitanloi365/machinery/v1/config"
	"github.com/thaitanloi365/machinery/v2"

	backend "github.com/thaitanloi365/machinery/v1/backends/eager"
	broker "github.com/thaitanloi365/machinery/v1/brokers/eager"
	lock "github.com/thaitanloi365/machinery/v1/locks/eager"
)

func TestRegisterTasks(t *testing.T) {
	t.Parallel()

	server := getTestServer(t)
	err := server.RegisterTasks(map[string]interface{}{
		"test_task": func() error { return nil },
	})
	assert.NoError(t, err)

	_, err = server.GetRegisteredTask("test_task")
	assert.NoError(t, err, "test_task is not registered but it should be")
}

func TestRegisterTask(t *testing.T) {
	t.Parallel()

	server := getTestServer(t)
	err := server.RegisterTask("test_task", func() error { return nil })
	assert.NoError(t, err)

	_, err = server.GetRegisteredTask("test_task")
	assert.NoError(t, err, "test_task is not registered but it should be")
}

func TestGetRegisteredTask(t *testing.T) {
	t.Parallel()

	server := getTestServer(t)
	_, err := server.GetRegisteredTask("test_task")
	assert.Error(t, err, "test_task is registered but it should not be")
}

func TestGetRegisteredTaskNames(t *testing.T) {
	t.Parallel()

	server := getTestServer(t)

	taskName := "test_task"
	err := server.RegisterTask(taskName, func() error { return nil })
	assert.NoError(t, err)

	taskNames := server.GetRegisteredTaskNames()
	assert.Equal(t, 1, len(taskNames))
	assert.Equal(t, taskName, taskNames[0])
}

func TestNewWorker(t *testing.T) {
	t.Parallel()

	server := getTestServer(t)

	server.NewWorker("test_worker", 1)
	assert.NoError(t, nil)
}

func TestNewCustomQueueWorker(t *testing.T) {
	t.Parallel()

	server := getTestServer(t)

	server.NewCustomQueueWorker("test_customqueueworker", 1, "test_queue")
	assert.NoError(t, nil)
}

func getTestServer(t *testing.T) *machinery.Server {
	return machinery.NewServer(&config.Config{}, broker.New(), backend.New(), lock.New())
}
