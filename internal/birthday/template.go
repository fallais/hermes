package birthday

// MessageTemplate is the birthday message template
const MessageTemplate = `
Greets !

This is the birthay of {{ contact }} !

{{ with age }}
{{ . }} years old ! :)
{{ end }}

Bye !
`
