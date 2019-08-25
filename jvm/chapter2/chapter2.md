1 本章内容和代码的功能是在命令行中指定类路径和类名时，我们的程序应该能正确找到该类，写完之后用jre路径和java.lang.Object测试，成功表明我
们的程序的确能通过命令行中的路径和名字，加载.class文件中内容




2  code complete, 准备install一下，结果main.go出问题了
    说main方法里面parseCmd, printUsage, Cmd 三个未定义
    main方法代码压根没变，chapter1里的main.go也install不了了
    同样的错误：
            PS F:\GoVM\src\go.buppt.cn\jvm\chapter1> go install main.go
            # command-line-arguments
            .\main.go:6:9: undefined: parseCmd
            .\main.go:11:3: undefined: printUsage
            .\main.go:17:20: undefined: Cmd

        install 方法错了，下面这么写就ok了
                PS F:\GoVM\src\go.buppt.cn\jvm\chapter2> go install go.buppt.cn/jvm/chapter2
        试了很多次，试对了
                PS F:\GoVM\src\go.buppt.cn\jvm\chapter2> go install jvm/chapter2
                can't load package: package jvm/chapter2: cannot find package "jvm/chapter2" in any of:
                        c:\go\src\jvm\chapter2 (from $GOROOT)
                        F:\GoVM\src\jvm\chapter2 (from $GOPATH)
                PS F:\GoVM\src\go.buppt.cn\jvm\chapter2> go install main
                can't load package: package main: cannot find package "main" in any of:
                        c:\go\src\main (from $GOROOT)
                        F:\GoVM\src\main (from $GOPATH)
                PS F:\GoVM\src\go.buppt.cn\jvm\chapter2> go install go.buppt.cn/jvm/chapter
                can't load package: package go.buppt.cn/jvm/chapter: cannot find package "go.buppt.cn/jvm/chapter" in any
                of:
                        c:\go\src\go.buppt.cn\jvm\chapter (from $GOROOT)
                        F:\GoVM\src\go.buppt.cn\jvm\chapter (from $GOPATH)
                PS F:\GoVM\src\go.buppt.cn\jvm\chapter2> go install go.buppt.cn/jvm/chapter2  -----OK