package executors

type Executor interface {
	Run() error
	Stop() error
}
