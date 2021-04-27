package soap

// Envelope interface
type Envelope interface {
	SetHeaders([]interface{})
	SetContent(interface{})
}

var _ Envelope = (*SOAPEnvelope)(nil)

// SetHeaders func
func (se *SOAPEnvelope) SetHeaders(headers []interface{}) {
	if se.Header == nil {
		se.Header = &SOAPHeader{}
	}
	se.Header.Headers = headers
}

// SetContent func
func (se *SOAPEnvelope) SetContent(content interface{}) {
	se.Body.Content = content
}
