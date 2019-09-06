# nginxconf

[![GoDoc](https://godoc.org/github.com/caas-one/nginxconf?status.svg)](https://godoc.org/github.com/caas-one/nginxconf)

[![Go Report Card](https://goreportcard.com/badge/github.com/caas-one/nginxconf)](https://goreportcard.com/badge/github.com/caas-one/nginxconf)

nginxconf 是一个 灵活的, 快捷的 nginx conf文件解析器, 能够在文件, json, go-struct, template 快速转换, 是nginx conf 在线可视化配置的基石 


### 依赖

### 环境

### ROADMAP

### Document

- [中文](https://github.com/caas-one/nginxconf/blob/master/doc/README-cn.md)


### 概要

#### nginxconf 的解析方式

nginxconf 解析nginxconf的方式如下图所示:

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

