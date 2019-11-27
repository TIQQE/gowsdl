package soap

// Envelope interface
type Envelope interface {
	SetHeaders([]interface{})
	SetContent(interface{})
}

var _ Envelope = (*SOAPEnvelope)(nil)
