# 项目名字

## 项目结构

├─assets			图片资源
├─config			配置文件
├─docs			  文档合集
├─global			全局变量
├─internal		  内部模块
│  ├─dao			数据访问层 Mysql 数据操作
│  ├─middleware	 HTTP中间件
│  ├─model		  模型层 存放model对象
│  ├─routers		路由相关逻辑处理
│  └─service		项目核心业务逻辑
├─pkg			   项目相关模块(包)
├─scripts		   各类构建,安装,分析脚本
├─storage		   项目生成临时文件
└─third_party	   第三方资源工具 UI库

## 数据库公共字段

```sql
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
```

## 外部库
借助第三方开源库 viper，在项目根目录下执行以下安装命令：

```$ go get -u github.com/spf13/viper@v1.4.0```

Viper 是适用于 Go 应用程序的完整配置解决方案，是目前 Go 语言中比较流行的文件配置解决方案，它支持处理各种不同类型的配置需求和配置格式。

##配置文件
config.yaml
Server:   服务配置,设置gin的运行模式,HTTP监听的端口,读写的最大持续时间
App:      应用配置,设置默认每页数量,所允许的最大页面数,log文件存放的位置
Database: 数据库配置,设置连接参数 

##日志文件
WithLevel：设置日志等级。
WithFields：设置日志公共字段。
WithContext：设置日志上下文属性。
WithCaller：设置当前某一层调用栈的信息（程序计数器、文件信息、行号）。
WithCallersFrames：设置当前的整个调用栈信息。
##模块开发
Model：指定运行 DB 操作的模型实例，默认解析该结构体的名字为表名，格式为大写驼峰转小写下划线驼峰。若情况特殊，也可以编写该结构体的 TableName 方法用于指定其对应返回的表名。
Where：设置筛选条件，接受 map，struct 或 string 作为条件。
Offset：偏移量，用于指定开始返回记录之前要跳过的记录数。
Limit：限制检索的记录数。
Find：查找符合筛选条件的记录。
Updates：更新所选字段。
Delete：删除数据。
Count：统计行为，用于统计模型的记录数。