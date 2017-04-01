package main

import (
	"flag"
	"fmt"
	"os"

	"strings"

	"github.com/dtkkki/allergy/measures"
	"github.com/dtkkki/allergy/textshandler"
)

const (
	version = "0.0.1"
)

var (
	// switch = flag.String("switch","", "The switch to decide on/off")
	dirtWords     = flag.String("dirty_words", "", "dirty words to be handle（输入要处理的污秽评论）")
	measure       = flag.String("measure", "DFA", "The measure you give to filter（选择过滤文本的算法）")
	senstiveWords = flag.String("swords", "", "The words you give to filter（输入指定的敏感词）")
	senstiveFile  = flag.String("sfile", "", "The file path you give to filter（输入指定的敏感词文件路径）")
	replace       = flag.Bool("replace", false, "Optional decide to replace（决定是否代替敏感词）")
	replaceWords  = flag.String("words", "*", "The words to be replaced（代替敏感词的词汇）")
)

func main() {
	flag.Parse()

	if *senstiveWords == "" &&
		*senstiveFile == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	fmt.Println("hello")

	if *dirtWords == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	switch *measure {
	case measures.DFAMeasure:
		m := measures.NewDFAFilter(*senstiveFile, strings.Split(*senstiveWords, ","))
		//TODO 包设计错误 正确语义： DFA 采用 正则表达式 预处理 污秽文本
		// 现在语义错误 DFA 要过滤的污秽文本 采用 正则表达式预处理
		ph := textshandler.NewRextHandler("[^a-zA-Z0-9\u4E00-\u9FA5]", "")
		t := textshandler.NewDirtyText(*dirtWords, *replace, *replaceWords, ph)
		run(m, t)
	default:
		flag.PrintDefaults()
		os.Exit(0)
	}

	fmt.Println("world")
}

func run(m measures.SensitiveProcess, t textshandler.DirtyTextInterface) {
	m.SensitiveWordInit()
	m.Process(t)
	fmt.Println("--filterText--", "todo....")
}
