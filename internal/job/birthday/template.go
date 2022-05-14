package birthday

// MessageTemplate is the birthday message template.
const MessageTemplate = `
Greets !

This is the birthay of {{ .contact }} !

{{ with age }}
{{ . }} years old ! :)
{{ end }}

Bye !
`

// TemplateData is the data to be executed in the template.
type TemplateData struct {
	contact string
	age     int
}
