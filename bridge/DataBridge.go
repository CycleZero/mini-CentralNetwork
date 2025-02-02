package bridge

type Command struct {
	TargetService string
	id            string
	Command       string
}

var InChan = make(chan Command, 100)
