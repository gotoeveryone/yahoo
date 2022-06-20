package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	version       = "development"
	IsShowVersion = flag.Bool("v", false, "show version")
	length        = flag.Int("l", 3, "length of the output")
)

func main() {
	flag.Parse()

	if *IsShowVersion || (len(os.Args) > 0 && os.Args[1] == "version") {
		fmt.Println(version)
		os.Exit(0)
	}

	if len(os.Args) > 0 && os.Args[1] == "help" {
		fmt.Println(`Usage: yahoo [subcommand] [options]
  subcommand
    version: show version
  options
    -l: output length
    -v: show version`)
		os.Exit(0)
	}

	v := []string{}
	for i := 0; i < *length; i++ {
		v = append(v, "                        |  |")
	}

	fmt.Printf(`
  ＊         +        巛 ヽ
                       〒  !      +        。          +        。          ＊         。
             +          |  |
%s
        ＊         +   /  /      イヤッッホォォォオオォオウ！
                ∧＿∧  /  /
              （´∀｀ /  /  +        。          +        。      ＊         。
              ,-       ｆ
             / ｭﾍ      | ＊         +        。          +      。  +
            〈＿｝）   |
                 /     \ +        。          +        +          ＊
               ./  ,ﾍ   |
    ｶﾞﾀﾝ  ||| j   /  |  | |||
  ――――――――――――
`, strings.Join(v, "\n"))
}
