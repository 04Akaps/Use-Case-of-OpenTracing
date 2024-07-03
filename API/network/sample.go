package network

type sample struct {
	n *Network
}

func newSampleRouter(n *Network) {
	s := &sample{n: n}

	// sample_base
	n.Router(GET, "/send", s.send)
	n.Router(GET, "/send-with-tag", s.sendWithTag)
	n.Router(GET, "/send-with-child", s.sendWithChild)

	// other_host
	n.Router(GET, "/receive-from-other-host", s.receiveSpanRouter)
	n.Router(GET, "/send-other-host", s.sendWithOtherHost)
	n.Router(GET, "/receive-two-from-other-host", s.receiveTwoSpanRouter)

	// panic_host
	n.Router(GET, "/send-for-panic", s.sendForPanic)
	n.Router(GET, "/receive-for-error", s.receiveForError)

	// baggage
	n.Router(GET, "/send-for-baggage", s.sendForBaggage)
	n.Router(GET, "/receive-for-baggage", s.receiveForBaggage)
}
