package gopool

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

// ------------------------------------------------------------------
//                             Dependencies
// ------------------------------------------------------------------

type taskMock struct {
	counter int
}

func (t *taskMock) Do() {
	t.counter += 1
}

func newTaskMock() Job {
	return &taskMock{
		counter: 0,
	}
}

// ------------------------------------------------------------------
//                             Test Config
// ------------------------------------------------------------------

type WorkerPoolTestSuite struct {
	suite.Suite
	maxWorker  uint
	capacity   uint
	workerPool WorkerPool
	task       Job
}

func (suite *WorkerPoolTestSuite) SetupSuite() {
	suite.maxWorker = 2
	suite.capacity = 1000
	suite.workerPool = NewWorkerPool(suite.maxWorker, suite.capacity)
}

func (suite *WorkerPoolTestSuite) SetupTest() {
	suite.task = newTaskMock()
}

// ------------------------------------------------------------------
//                             NewWorkerPool
// ------------------------------------------------------------------

func (suite *WorkerPoolTestSuite) TestNewWorkerPool_Success() {
	require := suite.Require()

	wp := NewWorkerPool(suite.maxWorker, suite.capacity)

	// checkers
	require.NotNil(wp.(*workerPool))
}

// ------------------------------------------------------------------
//                             AddTask
// ------------------------------------------------------------------

func (suite *WorkerPoolTestSuite) TestAddTask_Success() {
	require := suite.Require()

	suite.workerPool.AddTask(suite.task)
	suite.workerPool.Shutdown()

	// checkers
	require.Equal(suite.task.(*taskMock).counter, 1)
}

// ------------------------------------------------------------------
//                           Run All Tests
// ------------------------------------------------------------------

func TestWorkerPool(t *testing.T) {
	suite.Run(t, new(WorkerPoolTestSuite))
}
