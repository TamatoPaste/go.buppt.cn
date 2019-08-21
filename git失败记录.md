1  不会用vscode的git，只能在文件夹中使用git bash
    今天git commit显示Changes not staged for commit
    如下：
    $ git commit
        On branch master
        Your branch is up to date with 'origin/master'.
            Changes not staged for commit:
                modified:   bin/gocode.exe
                modified:   src/github.com/go-delve/delve (modified content)
                modified:   src/github.com/karrick/godirwalk (modified content)
                modified:   src/github.com/rogpeppe/godef (modified content)
                modified:   src/github.com/uudashr/gopkgs (modified content)
                modified:   "src/go.buppt.cn/\346\217\222\344\273\266\351\205\215\347\275\256\345\244\261\350\264\245\350\256\260\345\275\225.md"
                modified:   "src/go.buppt.cn/\346\236\201\345\256\242\346\227\266\351\227\264/2\345\217\230\351\207\217/main.go"
                modified:   "src/go.buppt.cn/\346\236\201\345\256\242\346\227\266\351\227\264/3\345\270\270\351\207\217/main.go"
                modified:   "src/go.buppt.cn/\346\265\213\350\257\225/first_test.go"
    意思就是文件有修改但是没有准备好commit，没有准备好commit的意思是，这个时候你也可以放弃这些修改，怎么才能准备好呢？
    git add . (注意add和点之间有空格)
    这时候就可以commit了
    git commit
    弹出一个文件，上面显示了修改，让你填message，随便填个啥，保存关闭文件
    commit完成
    再  git push
    OK，完工



2 准备在公司的电脑上把环境配置好，代码从GitHub上面clone下来，用vscode打开，插件需要重新安装，src/github.com里面代码都没有，好像git仓库里面有git仓库，内部仓库的代码不会下载，现在决定只在github上托管src目录下代码，不把github.com和golang.org文件夹中的内容托管。
把项目的markdown文件放一起了，出现错误refusing to merge unrelated histories

3 19-08-21,准备上传今日修改
    git push   ------>   fatal:No configured push destination.
    将url中的repo添加
    git remote add origin 'github.com/TamatoPaste/go.buppt.cn'
    再git push --------> fatal: The current branch master has no upstream           branch.
            To push the current branch and set the remote as upstream, use
                git push --set-upstream origin master
    原因是没有将本地的分支与远程仓库的分支进行关联

     git pull
        fatal: 'github.com/TamatoPaste/go.buppt.cn' does not appear to be a git repository
        fatal: Could not read from remote repository.

        Please make sure you have the correct access rights
        and the repository exists.
    百度弄不明白，关闭gitbash，重开，输入 git push
    TM的OK了，淦！
