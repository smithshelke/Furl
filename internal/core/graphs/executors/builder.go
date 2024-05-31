package executors

type ExecutorBuilder interface {
	WithConfig(Configurable) *ExecutorBuilder
	WithIO(IO) *ExecutorBuilder
	Build() Executor
}

type ExecutorConfigBuilder interface {
	Build() Configurable
}

type ExecutorIOBuilder interface {
	Build() (IO, error)
}
