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



### ROADMAP

### Document

- [中文](https://github.com/caas-one/nginxconf/blob/master/doc/README-cn.md)


### 概要

#### nginxconf 的解析方式

nginxconf 解析nginx.conf的方式如下图所示:

<div align=center><img  src="https://github.com/caas-one/nginxconf/blob/master/images/nginx-parse-method.png"/></div>

#### nginx conf 和 go struct 对应关系


```bash
|---------------------------------------|
|                                       |
|                Global                 |
|                                       |
|           |----------------|          |
|           |     event      |          |
|           |----------------|          |
|                                       |
|       |-----------------------|       |
|       |         http          |       |
|       |   |----------------|  |       |
|       |   |    Server      |  |       |
|       |   |----------------|  |       |
|       |   |    Location    |  |       |
|       |   |       ……       |  |       |
|       |   |       if       |  |       |
|       |   |       ……       |  |       |
|       |   |    Location    |  |       |
|       |   |----------------|  |       |
|       |                       |       |
|       |   |----------------|  |       |
|       |   |     Upstream   |  |       |
|       |   |----------------|  |       |
|       |   |       ……       |  |       |
|       |   |----------------|  |       |
|       |-----------------------|       |
|                                       |
|       |-----------------------|       |
|       |         mail          |       |
|       |   |----------------|  |       |
|       |   |    Server      |  |       |
|       |   |----------------|  |       |
|       |-----------------------|       |
|                                       |
|---------------------------------------|
```

###  caas-one

<div align=center><img width="150" height="150" src="https://github.com/caas-one/nginxconf/blob/master/images/caas-one.jpeg"/></div>


### 示例