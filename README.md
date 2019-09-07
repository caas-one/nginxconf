# nginxconf
[![GoDoc](https://godoc.org/github.com/caas-one/nginxconf?status.svg)](https://godoc.org/github.com/caas-one/nginxconf)
[![Go Report Card](https://goreportcard.com/badge/github.com/caas-one/nginxconf)](https://goreportcard.com/badge/github.com/caas-one/nginxconf)

nginxconf 是一个 灵活的, 快捷的 nginx conf文件解析器, 能够在文件, json, go-struct, template 快速转换, 是nginx conf 在线可视化配置的基石 


### 依赖

- 第一步: 安装 Python 环境, Python2, Python3 均可,自行安装
- 第二步: 安装 Python 库: `pip install cscrossplane`

### 安装

```
go get -u github.com/caas-one/nginxconf
```

:blue_book: 安装时,可能会发生一些意外的问题, 我们把可能的问题已经整理到文档,并[提供了解决办法, 点击查看](https://github.com/caas-one/nginxconf/blob/master/doc/install_q.md) , 如果您在文档中未查找到相关问题,欢迎向我们反馈和提 issue.


### Document

- [中文](https://github.com/caas-one/nginxconf/blob/master/doc/README-cn.md)


### 概要

随着时代的发展, 各企业纷纷提出了更高标准的运维需求: 
- 操作可追溯性
- 自动化运维

nginx 作为用户最多的软件之一, nginx配置的变更可追溯性面对巨大的挑战, 同时,日积月累的变更操作造就了冗长的,晦涩的 nginx.conf, 最终可能造成 变不敢变 的结局,基于现状和更高标准的运维需求, 产生了nginxconf

#### nginxconf 现阶段可能给您带来哪些**收益**呢?

- nginx 配置变更可追溯
- nginx 配置可视化
- 按规范/约定自动拆分 nginx 配置
- nginx 配置自动创建并分发(如: 在kubernetes中, 当watch到 nodeport svc的创建, 自动生成 nginx 配置转发)
- nginx 配置多节点分发

为达到上面的需求和收益, nginxconf 所处的地位是什么呢? 

如果上面的需求和收益是业务诉求, 那么 nginxconf 就是实现这些诉求的**基石** ,一切的诉求都在这个基石之上. 

#### nginxconf 真正做了那些事情呢?

- 解析nginx.conf
- 转换成结构化的模型
- 根据模板生成nginx配置

#### nginxconf 解析和渲染方式

- 第一步: nginxconf 通过 cscrossplane (这是一个python库,[如何安装Python?](https://www.python.org/getit/))把 **nginx-in.conf** 解析成一个中间态json
- 第二步: 接着通过巧妙而又暴力的方式(指令判断的方式)把 **中间态json** 中的数据转换到结构化的 **Go struct** 中, 这样就避免了 nginx.conf 中 if 指令和 location 指令带来的麻烦 
- 第三步:最后根据 **template** 把Go struct中的数据重新渲染成 **nginx-out.conf**

流程如下图所示:

<div align=center><img  src="https://github.com/caas-one/nginxconf/blob/master/images/nginx-parse-method.png"/></div>

#### Go struct 和 nginx.conf 层级对应关系

在 Go struct 和 nginx.conf 映射关系设计中, 完全使用了平行映射的方式(以 nginx.conf 中的指令直接作为 Go struct 中的字段名), 使用这种方式可以让开发者如丝般顺滑的接入 nginxconf ,我们以对比的方式生成了一张对比图,如下所示:

<div align=center><img  src="https://github.com/caas-one/nginxconf/blob/master/images/mapper.jpg"/></div>


### 示例
