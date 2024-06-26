package raft

// StateMachine is the interface for allowing the host application to save and
// recovery the state machine. This makes it possible to make snapshots
// and compact the log.
type StateMachine interface {
	Save() ([]byte, error)
	Recovery([]byte) error
}

type defaultStateMachine struct {
	smData []byte
}

func NewStateMachine() StateMachine {
	return &defaultStateMachine{}
}

func (sm *defaultStateMachine) Save() ([]byte, error) {
	return sm.smData, nil
}

func (sm *defaultStateMachine) Recovery(b []byte) error {
	sm.smData = b
	return nil
}
