# 使用官方Nginx镜像作为基础镜像
FROM nginx:1.16.1

# 执行命令修改基础镜像的内容
RUN echo "hello world" > /usr/share/nginx/html/index.html

# 允许外界访问容器的80端口
EXPOSE 80

