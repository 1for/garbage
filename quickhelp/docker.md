docker run -t -i ubuntu:14.04 /bin/bash
docker run ubuntu:14.04 /bin/echo 'Hello world'

#后台执行
docker run -d ubuntu:14.04 /bin/sh -c "while true; do echo hello world; sleep 1; done"

#查看日志
docker container logs [container ID or NAMES]

#停止容器
docker container stop [container ID or NAMES]

#进入容器
docker exec -i [container ID or NAMES] /bin/bash

#运行 nginx 容器
docker run -p 8093:80 --name mynginx  -v $PWD/conf/nginx.conf:/etc/nginx/nginx.conf -v $PWD/www:/opt/nginx/www -v $PWD/log:/opt/nginx/log  -d nginx

-p 8093:80：将容器的80端口映射到主机的8093端口
-name mynginx：将容器命名为mynginx
-v $PWD/conf/nginx.conf:/etc/nginx/nginx.conf：将主机中当前目录下的nginx.conf挂载到容器的/etc/nginx/nginx.conf
-v $PWD/www:/opt/nginx/www：将主机中当前目录下的www挂载到容器的/opt/nginx/www，参考nginx.conf的root配置
-v $PWD/log:/opt/nginx/log：将主机中当前目录下的log挂载到容器的/opt/nginx/log，参考nginx.conf的log配置

#对比容器内容变更
docker diff [container ID or NAMES]

#提交容器变更内容
docker commit [container ID or NAMES] nginx:ly-0.0.1

#查看镜像的历史
docker history [IMAGE]
