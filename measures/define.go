package measures

import (
	"github.com/dtkkki/allergy/textshandler"
)

const (
	//DFAMeasure ...
	DFAMeasure = "DFA"
)

//SensitiveProcess deine standard handle process
type SensitiveProcess interface {
	SensitiveWordInit()
	Process(text textshandler.DirtyTextInterface)
}
