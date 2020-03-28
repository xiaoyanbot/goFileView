goFileView 基于go语言的文件预览
===========================

本仓库为对 https://github.com/leeli73/goFileView.git 的folk分支版本。

GoFileView是受kkFileView( *https://gitee.com/kekingcn/file-online-preview.git*
)启发并基于其网站前端开发的。目前goFileView处于最原始的起步状态，相对简陋，相信随着不断完善成为一套强壮的系统。本人代码风格相对较”狂”，欢迎大家一起来提出建议和完善Go
File View。

特别要感谢kkfileview的开源，让我可以使用它的前端页面直接开发。调用方式也在很大程度上参考了kkfileview。

仅支持Linux系统，url暂时不支持中文等问题。  已测试Ubuntu适用。

目前已经完成
============

Word、Excel、PPT转码为PDF

对Word,Excel,PPT和PDF的在线预览

TODO
====

内置File Server

本地路径指定，省去下载步骤

部署编译
========

准备
----

安装Libreoffice

确保Libreoffice在path目录下

编译
----

下载源码

设置 GOPATH 为当前的目录

cd goFileView

go build main.go


