# GONEList

<img align="right" width="240" src="https://github.com/cugxuan/gonelist/wiki/img/logo/logo.png">

[![Build Status](https://travis-ci.org/cugxuan/gonelist.svg)](https://travis-ci.org/cugxuan/gonelist)
[![Latest Release](https://img.shields.io/github/release/cugxuan/gonelist.svg)](../../releases)
[![All Releases Download](https://img.shields.io/github/downloads/cugxuan/gonelist/total.svg)](../../releases)

这是一款使用 `golang + vue` 编写的基于 onedrive 的**在线共享网盘**  
效果展示：https://gonelist.cugxuan.cn  
前端项目地址：https://github.com/Sillywa/gonelist-web  
有问题请提 issue，也可以进入 QQ 群交流，群号：1083165608

# 功能特性

- 「跨平台」，兼容 Linux/Windows/MacOS 等多个平台
- 「极速部署」，下载对应程序，修改配置即可前后端一键部署
- 「直链下载」，文件直链下载，下载不消耗服务器流量
- 「自动刷新」，自动刷新缓存，可自定义时间
- 「自定义目录」，支持将 onedrive 的某个目录作为根目录
- 「在线播放」，支持在线播放音频和视频
- 「海量文件」，支持单目录下上千个文件，正常使用
- 「多平台」，支持个人版、教育账号、世纪互联等
- 「README」，支持页面添加 README
- ...

注：支持绝大部分教育账号，部分 **教育账号** 因为需要管理员同意无法使用


# 安装配置

# 实体服务安装教程
如果您的**整个微软账号和 onedrive 网盘**内，**没有隐私内容**，可以按照下面的流程快速配置体验效果，完整的下载安装流程请看 [安装文档](https://github.com/cugxuan/gonelist/wiki/Install)  

## 快速配置体验
下载 [Github Release](https://github.com/cugxuan/GOIndex/releases) 或者 [gonelist-release](https://gonelist.cugxuan.cn/#/gonelist-release) 中对应的包，直接运行即可启动，以 Linux 系统本地启动为例
```
// 下载对应的安装包，也可下载 gonelist-release 中的包
$ wget https://github.com/cugxuan/gonelist/releases/download/v0.3.2/gonelist_linux_amd64.tar.gz
// 如果速度过慢，还可以使用 CDN 链接下载
$ wget http://g.cugxuan.cn/v0.3.2/gonelist_linux_amd64.tar.gz
$ tar -zxf gonelist_linux_amd64.tar.gz && cd gonelist_linux_amd64
$ ./gonelist_linux_amd64
```
打开 http://localhost:8000 按照提示登录后即可。如果是在本地部署，登陆成功会跳转到首页，此时已经完成部署。  
如果是在服务器部署，登陆成功会跳转到 http://localhost:8000/auth?code=xxx ，将当前网址改成 http://yoursite:8000/auth?code=xxx 再回车等待文件加载后，会自动跳转你的网站 http://yoursite:8000 。如果登陆后一直没有反应，可能是因为文件夹数量过多导致，建议设置「子文件夹」选项  
默认情况下读取当前路径的 `config.json` 作为配置文件，或加参数 `--conf=dir1/file.json` 指定配置文件路径

## 实体systemd服务安装

视频教程(包含了Azure应用程序的配置) https://www.bilibili.com/video/BV1PA411t7Jw/

## docker运行
视频教程 https://www.bilibili.com/video/BV1Vz4y1R7EK/

直接使用项目的`docker-compose.yml`去`docker-compose up -d`即可，建议把配置文件放在一个文件夹里，把文件夹挂载进去。否则直挂文件docker挂载的是inode
```
.
├── config
│   └── config.json
└── docker-compose.yml
```
配置文件最好是切到对应的tag那去下载，容器的话配置文件的`dist_path`值得改为`/etc/dist/`

## config.json

如果需要修改配置，在配置文件中填对对应的内容即可
```
{
  //------建议填入自己的 id 和 secret --------
  "client_id": "88966400-cb81-49cb-89c2-6d09f0a3d9e2",
  "client_secret": "/FKad]FPtKNk-=j11aPwEOBSxYUYUU54",
  "redirect_url": "http://localhost:8000/auth",
  // 设置一个自己喜欢的字符串
  "state": "23333",
  "china_cloud": false,
  "server": {
    // 监听的端口
    "port": 8000,
    // 自动刷新的时间单位是分钟，默认 10 分钟，不要超过 1 小时
    "refresh_time": 10,
    // 登陆成功后，跳转的 URL，可不设置，新版已自动跳转
    "site_url": "http://localhost:8000",
    // 自定义 onedrive 的子文件夹
    "folder_sub": "/",
    //静态页面的目录，默认当前路径下的dist目录
    "dist_path": "./dist/",
    // 是否绑定到 0.0.0.0
    "bind_global": true
  }
}
```

# Contributors

- 开发：<a href="https://github.com/cugxuan"><img src="https://avatars1.githubusercontent.com/u/23120372?s=400&v=4" width="30"></a>
<a href="https://github.com/Sillywa/"><img src="https://avatars0.githubusercontent.com/u/22909601?s=400&v=4" width="30"></a>
<a href="https://github.com/zhangguanzhang"><img src="https://avatars3.githubusercontent.com/u/18641678?s=400&v=4" width="30"></a>
- logo 设计：<a href="http://lambertchan.me/"><img src="https://avatars0.githubusercontent.com/u/39192150?s=400&v=4" width="30"></a>