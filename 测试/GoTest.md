go test 命令会运行文件夹下所有的单元测试函数，其中的log，print函数的输出内容会打印到控制台上
    F:\GoVM\src\go.buppt.cn\测试> go test
        我是回调，x：1
        我是回调，x：2
        PASS
        ok      go.buppt.cn/测试        0.256s
go test fisrt_test.go 指明文件后，只会输出是否通过及运行时长
    F:\GoVM\src\go.buppt.cn\测试> go test first_test.go
        ok      command-line-arguments  0.275s