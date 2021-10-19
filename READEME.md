

构建镜像

docker build . -t mach86devops/k8s_work:v2.1.1

启动镜像

docker run -p 8080:8080 mach86devops/k8s_work:v2.1.1 --

发布镜像

docker login

可以使用中国国内的镜像网站

docker push 

推送



$ wget https://www.kernel.org/pub/linux/utils/util-linux/v2.24/util-linux-2.24.tar.gz  
$ tar -xzvf util-linux-2.24.tar.gz  
$ cd util-linux-2.24/  
$ ./configure --without-ncurses  
$ make nsenter  
$ sudo cp nsenter /usr/local/bin 


PID=$(docker inspect --format {{.State.Pid}} <container_name_or_ID>)

sudo nsenter --target $PID --mount --uts --ipc --net --pid 