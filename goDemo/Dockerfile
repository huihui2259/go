FROM centos:latest
COPY main .
COPY init.sh .
EXPOSE 8089

# 目前继承的centos镜像缺少很多命令，这里构造镜像的时候直接安装net-tools
RUN cd /etc/yum.repos.d/
RUN sed -i 's/mirrorlist/#mirrorlist/g' /etc/yum.repos.d/CentOS-*
RUN sed -i 's|#baseurl=http://mirror.centos.org|baseurl=http://vault.centos.org|g' /etc/yum.repos.d/CentOS-*
RUN yum -y install net-tools

CMD ["./main"]