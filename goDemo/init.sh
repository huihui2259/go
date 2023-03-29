#!/bin/bash
# 构建镜像时容器里会缺少很多组件，这里是安装net-tools，这样就可以使用netstat命令
cd /etc/yum.repos.d/

sed -i 's/mirrorlist/#mirrorlist/g' /etc/yum.repos.d/CentOS-*

sed -i 's|#baseurl=http://mirror.centos.org|baseurl=http://vault.centos.org|g' /etc/yum.repos.d/CentOS-*

yum -y install net-tools