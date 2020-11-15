# 网页版小说解析服务

## 服务程序说明

#### NovelServer 

将网页版的小说解析成json格式API输出的模式。
可以通过配置文件进行小说网站源的配置。

[NovelServer API文档说明](https://docs.apipost.cn/view/1e517785a3282a91#3196964)


## 技术栈

基础框架：[GoFrame](https://github.com/gogf/gf) 【 [中文文档](https://goframe.org/index) 】

网页解析库：[goquery](https://github.com/PuerkitoBio/goquery)     

## 源码运行

在项目目录下通源码直接行

`go run main.go`

通过编译脚本，进行编译二进制文件运行。编译后的文件在bin目录下，把自动将需要的相关文件复制到该目录下。

`./build.sh`

分别有三个参数可选支持不同的操作系统，linux|windows|mac

**关于curl.sh 运行脚本**

在mac或是linux下通过该脚本可以后台运行服务。通过参数进行启动、重启、停止及显示后台实时信息的操作。

```
应用的启动命令说明：

./curl.sh pid|start|stop|restart|status|tail

```

> start：启动应用
> stop：停止应用
> restart：重新启动应用
> status：查看应用状态
> tail：查看应用运行的动态输出日志信息



## 部署说明

### 一、独立部署

服务器推荐使用*nix服务器系列(包括:Linux, MacOS, *BSD)，以下使用Linux系统为例，介绍如何部署。

将应用服务目录复制到目标位置，里面已经写好了执行的脚本，通过脚本来执行。

```
curl.sh脚本参数：

start|stop|restart|status|tail

```
### 二、代理部署

推荐使用Nginx作为反向代理的前端接入层，有两种配置方式实现动静态请求的拆分。

```
server {
    listen       80;
    server_name  www.abc.com;

    access_log   /var/log/gf-app-access.log;
    error_log    /var/log/gf-app-error.log;

    location ~ .*\.(gif|jpg|jpeg|png|js|css|eot|ttf|woff|svg|otf)$ {
        access_log off;
        expires    1d;
        root       /var/www/gf-app/public;
        try_files  $uri @backend;
    }

    location / {
        try_files $uri @backend;
    }

    location @backend {
        proxy_pass                 http://127.0.0.1:8199;
        proxy_redirect             off;
        proxy_set_header           Host             $host;
        proxy_set_header           X-Real-IP        $remote_addr;
        proxy_set_header           X-Forwarded-For  $proxy_add_x_forwarded_for;
    }
}

```

其中，8199为NovelServer应用Web服务监听端口。这个端口在config.toml文件的server下Address参数中配置。

### 三、supervisor

`supervisor`是用`Python`开发的一套通用的进程管理程序，能将一个普通的命令行进程变为后台`daemon`，并监控进程状态，异常退出时能自动重启。官方网站：http://supervisord.org/ 常见配置如下：

```undefined
[program:NovelServer]
user=root
command=/var/www/NovelServer
stdout_logfile=/var/log/NovelServer-stdout.log
stderr_logfile=/var/log/NovelServer-stderr.log
autostart=true
autorestart=true
```

使用步骤如下：

1. 使用`sudo service supervisor start`启动`supervisor`服务；
2. 创建应用配置文件`/etc/supervisor/conf.d/NovelServer.conf`, 内容如上;
3. 使用`sudo supervisorctl`进入`supervisor`管理终端；
4. 使用`reload`重新读取配置文件并重启当前`supoervisor`管理的所有进程；
5. 也可以使用`update`重新加载配置(默认不重启)，随后使用`start AssessServer启动指定的应用程序；
6. 随后可以使用`status`指令查看当前`supervisor`管理的进程状态；

## 应用配置文件

在config中进行日志、数据库、服务端口及白名单等设置。配置实时生效。

`

## 采用模板配置文件说明


**Rule字段的定义和用法**

通过 css selector的定位进行数据提取。可以下载chrome的扩展程序[SelectorGadget](https://chrome.google.com/webstore/detail/selectorgadget/mhjhnkcfbdhnjickkkdbjoemdmbfginb/related?hl=zh-CN) 进行辅助选取。

**字段说：**

**`Range`：**内容截取范围
**`Type`：**选择的类型，text,src,href,alt
**`Rule`:** 选择的起点，如："div[class='bookname'] h1"

> :nth-child(n) 选择器匹配属于其父元素的第 N 个子元素，不论元素的类型。
>
> n 可以是数字、关键词或公式。
>

**`Filter`:** 过滤器，采用正规选择。

> **例：**
>
> 1、删除空行
>
> ```
>  \s
> ```
>
> 2、删除多余的汉字与空格及空行
> 如：
> ```
> 作者：\n                                    \n                                        异乡说书人\n                                    \n                                
> ```
> 正则：
>
> ```
> [作者：]|[\s]
> ```

