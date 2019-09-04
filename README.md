# nginx-go
Quick and reliable way to convert NGINX configurations into JSON and back.


#### 用户目标:

1. 给定一个nginx conf文件 ==> go 结构体
2. 单个域名的配置生成 本质是多个server和upstream
3. 按照域名拆分成单个文件, 易于管理
4. 