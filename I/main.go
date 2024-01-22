package I

type Reader interface {
	Read() string
}

type Writer interface {
	Write(content string)
}

type Printer interface {
	Print() string
}

type TextDocument struct {
	content string
}

func (d *TextDocument) Read() string {
	return d.content
}

func (d *TextDocument) Write(content string) {
	d.content = content
}

func (d *TextDocument) Print() string {
	return "Printing: " + d.content
}

// Implement Reader, Writer, and Printer for TextDocument

type ReadOnlyDocument struct {
	content string
}

func (d *ReadOnlyDocument) Read() string {
	return d.content
}

// Implement only Reader for ReadOnlyDocument
