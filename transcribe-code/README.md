# 抄代码

前有大神抄书, 本人抄代码. 既然看不下去, 直接手动敲一份. 

代码随心情加注释, 代码主要语言控制在3种以内, 使用最新tag为基准. 

测试代码不在抄写范围内, 测试代码自己写, 测试代码主要针对不理解的逻辑进行测试.

抄写不一定涉及完整项目, 但要包含所有依赖模块.

## 部分命令

* 快速创建目录结构

```sh
find . -type d >tmp.txt
xargs mkdir -p <tmp.txt
```

## 工具

### golang

生成依赖图工具 `go get github.com/google/godepq`, 使用`graphviz`生成图片或svg

命令`godepq -from github.com/google/godepq -o dot | dot -Tpng -o test.png`

golang严格限制包循环依赖, 总能找到没有依赖的包.


## 目录

* [goleveldb](https://github.com/syndtr/goleveldb) commit:64b5b1c
