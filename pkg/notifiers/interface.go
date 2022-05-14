package notifiers

// Notifier interface
type Notifier interface {
	Notify(string) error
	Name() string
}
