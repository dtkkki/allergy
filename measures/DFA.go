package measures

import (
	"fmt"

	"github.com/dtkkki/allergy/textshandler"
	"github.com/dtkkki/allergy/toolkits"
	"github.com/streamrail/concurrent-map"
)

//DFAFilter ...
type DFAFilter struct {
	dfaMap        cmap.ConcurrentMap
	path          string
	targetedwords []string
}

//DFAFilter ...
func NewDFAFilter(path string, ws []string) *DFAFilter {
	s := &DFAFilter{
		dfaMap:        cmap.New(),
		path:          path,
		targetedwords: ws,
	}
	fmt.Println("--initDFA--\n\n")
	return s
}

//SensitiveWordInit ...
func (sw *DFAFilter) SensitiveWordInit() {
	sw.targetedwords = append(sw.targetedwords, toolkits.InitSensitiveWordsFromPath(sw.path)...)
	if 0 == len(sw.targetedwords) {
		return
	}
	var currentMap cmap.ConcurrentMap
	var newWordMap cmap.ConcurrentMap
	for _, word := range sw.targetedwords {
		currentMap = sw.dfaMap
		for index, c := range []rune(word) {
			cstr := string(c)
			cval, ok := currentMap.Get(cstr)
			if ok {
				currentMap = cval.(cmap.ConcurrentMap)
			} else {
				newWordMap = cmap.New()
				newWordMap.Set("isEnd", "0")
				currentMap.Set(cstr, newWordMap)
				currentMap = newWordMap
			}

			if index == len([]rune(word))-1 {
				currentMap.Set("isEnd", "1")
			}
		}

	}
	xxoo, _ := sw.dfaMap.MarshalJSON()
	fmt.Println("--构成敏感词库--\n", string(xxoo), "\n\n")
}

func (sw *DFAFilter) checkSensitiveWord(text textshandler.DirtyTextInterface) {

}

//Process ...
func (sw *DFAFilter) Process(text textshandler.DirtyTextInterface) {
	text.PreHandle()

}
