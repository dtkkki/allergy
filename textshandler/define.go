package textshandler

import "regexp"
import "fmt"

//DirtyTextInterface ...需要处理的污秽评论
type DirtyTextInterface interface {
	PreHandle()
}

//DirtyText 污秽文本
type DirtyText struct {
	content        string
	replace        bool
	replaceSymbols string
	prev           *DirtyText
	next           *DirtyText
	handler        DirtyTextHandler
}

func NewDirtyText(text string, replace bool, symbols string, handler DirtyTextHandler) *DirtyText {
	return &DirtyText{
		content:        text,
		replace:        replace,
		replaceSymbols: symbols,
		handler:        handler,
	}
}

//PreHandle 预处理函数
//比如： { 强-奸-习-近-平} --> { 强奸习近平}
func (dt *DirtyText) PreHandle() {
	dt.content = dt.handler.Handle(dt.content)
	fmt.Println("--preHandle--\n", dt.content, "\n\n")
}

func (dt *DirtyText) IsReplace() bool {
	return dt.replace
}

func (dt *DirtyText) NextText() *DirtyText {
	return dt.next
}

func (dt *DirtyText) PrevText() *DirtyText {
	return dt.prev
}

//DirtyTextHandler ...对污秽评论的一些处理回调函数
type DirtyTextHandler interface {
	Handle(text string) string
}

//RegextHandler 通过正则表达式处理污秽评论中特殊字段
type RegextHandler struct {
	regext   *regexp.Regexp
	replacer string
}

func NewRextHandler(regext, replace string) *RegextHandler {
	nr, err := regexp.Compile(regext)
	if err != nil {

	}
	return &RegextHandler{
		regext:   nr,
		replacer: replace,
	}
}

//Handle forexample regexp = "[^a-zA-Z0-9\u4E00-\u9FA5]"
func (rh *RegextHandler) Handle(text string) string {
	return rh.regext.ReplaceAllString(text, rh.replacer)
}
