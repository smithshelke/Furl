package executors

type Configurable interface {
	Get() (map[string]any, error)
}
