# 错误码管理(business-code)
## 介绍
基于go+vue 简单实现的 程序错误码管理微服务，错误码由模块(2位)、类型(2位)和业务错误码(4位)组成，采用sqlite数据库记录。管理页面进入static 目录启动服务即可。
接口监听端口8095，管理页面监听端口8096 主要用于错误码规范连接，管理错误码
## 部署
1. 解压
2. 执行 bash server.sh start 会启动程序监听2个端口8095（接口地址）、8096（html http服务）
3. nginx代理，error-code.huishoubao.com.cn.conf 是nginx配置案例，可以修改需要的域名
## todo
1. 参数规则校验
2. 列表搜索优化
3. 根据中文名称，自动建议常量表达试
