package reminder

type Reminder interface {
	Run()
	GetCRONExpression() string
}
