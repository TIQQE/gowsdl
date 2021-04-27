package soap

// Envelope interface
type Envelope interface {
	SetHeader(*SOAPHeader)
	SetContent(interface{})
}

var _ Envelope = (*SOAPEnvelope)(nil)

// SetHeader func
func (se *SOAPEnvelope) SetHeader(header *SOAPHeader) {
	se.Header = header
}

// SetContent func
func (se *SOAPEnvelope) SetContent(content interface{}) {
	se.Body.Content = content
}
