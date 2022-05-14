package thing

// MessageTemplate is the thing message template.
const MessageTemplate = `
Greets !

You have something to do : {{ .Thing }}
`

// TemplateData is the data to be executed in the template.
type TemplateData struct {
	Thing string
}
