# Demo Tasks

<br/>

## Task 0: Install a ubuntu 16.04 server 64-bit

> **Steps:**

>   - Install virtualbox on laptop from laptop
>   - Download iso image from [http://releases.ubuntu.com/16.04/ubuntu-16.04.6-server-amd64.iso](http://releases.ubuntu.com/16.04/ubuntu-16.04.6-server-amd64.iso "ubuntu-16.04-server-amd64")
>   - Create VM with NAT network, 4G MEM, and 2 CPU using the image above
>   - Set port forward from host machine to the VM
>     - 2222 -> 22
>     - 8080 -> 80
>     - 8081 -> 8081
>     - 8082 -> 8082
>     - 31080 -> 31080
>     - 31081 -> 31081
>   - Start VM and login to
>   - Add current user to `sudoers` and grant the ability for the user running `sudo` without password
```
>sudo su -
>chmod u+w /etc/sudoers
```
>     add `account_name ALL=(ALL:ALL) ALL` under `root ALL=(ALL:ALL) ALL` in `/etc/sudoers`
>     change `%sudo ALL=(ALL:ALL) ALL` to `%sudo ALL=(ALL:ALL) NOPASSWD:ALL`
```
>chmod u-w /etc/sudoers
```

<br/>

## Task 1: Update system

> **Steps:**

> - Set proxy
```
>export http_proxy=proxy_url
>export https_proxy=proxy_url
>echo 'Acquire::http::Proxy "proxy_url";' > /etc/apt/apt.conf.d/100proxy.conf
>echo 'Acquire::https::Proxy "proxy_url";' >> /etc/apt/apt.conf.d/100proxy.conf
```
> - Check current kernel
```
>uname -r
4.4.0-142-generic
```
> - Upgrade the system
```
apt update
apt upgrade -y
reboot
```
> - Check upgraded kernel
```
>uname -a 
4.4.0-170-generic
```

> **Optional Steps**:

> - Take a snapshot of current VM
> - Get latest kernel packages in mainline
```
>wget https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.207/linux-headers-4.4.207-0404207_4.4.207-0404207.201912210540_all.deb
>wget https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.207/linux-headers-4.4.207-0404207-generic_4.4.207-0404207.201912210540_amd64.deb
>wget https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.207/linux-image-unsigned-4.4.207-0404207-generic_4.4.207-0404207.201912210540_amd64.deb
>wget https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.207/linux-modules-4.4.207-0404207-generic_4.4.207-0404207.201912210540_amd64.deb
```
> - Install packages for kernel
```
>dpkg -i -R .
>reboot
```
> - Check upgraded kernel
```
>uname -r
4.4.207-0404207-generic
```

<br/>

## Task 2: Install gitlab-ce version in the host

> **Steps:**

> - Install gitlab-ce
```
>curl -sS https://packages.gitlab.com/install/repositories/gitlab/gitlab-ce/script.deb.sh | bash
>EXTERNAL_URL="http://127.0.0.1" apt-get install gitlab-ce -y
```
> - Login gitlab
>   - open `http://127.0.0.1:8080` in host machine browser
>   - set password
>   - log in as `root` user

> - Disable monitoring process to reduce mem usage 
```
vi /etc/gitlab/gitlab.rb
```
>   - `monitoring_role['enable'] = false`
>   - `prometheus['enable'] = false`
>   - `node_exporter['enable'] = false`
>   - `grafana['enable'] = false`
> - Make the changes effect
```
>gitlab-ctl reconfigure
```
> - Verify gitlab is running
>   - open `http://127.0.0.1:8080` in host machine browser and log in

<br/>

## Task 3: Create a demo group/project in gitlab

> **Steps:**

> - Create group `demo` and project `go-web-hello-world` under the group `demo` in gitlab via GUI
> - Install docker and pull golang image	
```
>curl -fsSL https://get.docker.com -o get-docker.sh
>sudo sh get-docker.sh
>docker version
Client: Docker Engine - Community
 Version:           19.03.5
 API version:       1.40
 Go version:        go1.12.12
 Git commit:        633a0ea838
 Built:             Wed Nov 13 07:50:12 2019
 OS/Arch:           linux/amd64
 Experimental:      false

Server: Docker Engine - Community
 Engine:
  Version:          19.03.5
  API version:      1.40 (minimum version 1.12)
  Go version:       go1.12.12
  Git commit:       633a0ea838
  Built:            Wed Nov 13 07:48:43 2019
  OS/Arch:          linux/amd64
  Experimental:     false
 containerd:
  Version:          1.2.10
  GitCommit:        b34a5c8af56e510852c35414db4c1f4fa6172339
 runc:
  Version:          1.0.0-rc8+dev
  GitCommit:        3e425f80a8c931f88e6d94a8c831b9d5aa481657
 docker-init:
  Version:          0.18.0
  GitCommit:        fec3683
>docker pull golang:1.13.5
```
> - git installed fater Tasks 1
> - Create git repo locally
```
>mkdir /root/golang
>cd golang
>git init
```
> - Create a go APP
>   - edit a `go-web-hello-world.go` with following codes
```
package main

import (
    "fmt"
    "net/http"
    "flag"
    "strconv"
    "log"
)

func main() {
    var p int
    flag.IntVar(&p, "port", 80, "Need a postive number less than 65535")
    flag.Parse()

    if p <= 0 || p >= 65535 {
        p = 80
    }
    ps := strconv.Itoa(p)

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Go Web Hello World!")
    })
    log.Print("Listening at :" + ps)
    log.Fatal(http.ListenAndServe(":" + ps, nil))
}
```


> - Build the go APP
```
>cd /root/golang
>docker run --rm -v /root/golang:/go/src golang:1.13.5 go build -o src/go-web-hello-world src/go-web-hello-world.go
```
> - Check-in code to repo
```
>git config --global user.name "Your Name"
>git config --global user.email "Your Email Addr"
>git add main.go
>git add go-web-hello-world
>git commit -m "hello world"
>git remote add origin http://127.0.0.1/demo/go-web-hello-world.git
>git push origin master
```
> - Verify code in repo
>   - open `http://127.0.0.1:8080/demo/go-web-hello-world` in host machine browser


<br/>

## Tasks 4: Build the app and expose ($ go run) the service to 8081 port

> **Steps:**

> - Build and run the APP
```
>cd /root/golang
>docker run --rm -v /root/golang:/go/src golang:1.13.5 go build -o src/go-web-hello-world src/go-web-hello-world.go
./go-web-hello-world --port 8081
```
>   - open `http://127.0.0.1:8081` in host machine browser, got `Go Web Hello World!`


<br/>


## Tasks 5: Install docker

> **Steps:**

> - Done in Tasks 3


<br/>


## Task 6: Run the APP in container

> **Steps:**

> - Rebuild the go APP with `CGO_ENABLED=0`
>   - check [https://forums.docker.com/t/standard-init-linux-go-195-exec-user-process-caused-no-such-file-or-directory/43777](https://forums.docker.com/t/standard-init-linux-go-195-exec-user-process-caused-no-such-file-or-directory/43777)
```
>/root/golang/
>vi Dockerfile
```
>   - put below in `Dockerfile`
```
# this is to build the image
FROM golang:alpine as build

WORKDIR /go/src/go-web-hello-world
ADD . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

CMD ["./go-web-hello-world"]

# this is to publish the image
FROM scratch AS prod

COPY --from=build /go/src/go-web-hello-world/go-web-hello-world .
CMD ["./go-web-hello-world"]
```
> - run docker build
```
>docker build -t go-web-hello-world:v0.1 .
>docker start -p 8082:80 go-web-hello-world:v0.1
```
>   - open `http://127.0.0.1:8081` in host machine browser, got `Go Web Hello World!

> - Check in the Dockerfile
```
>cd /root/go/src/go-web-hello-world
>git add Dockerfile
>git commit -m "dockerfile"
>git push
```

<br/>


## Task 7: Push image to dockerhub

> **Steps:**
```
>docker tag go-web-hello-world docker.io/slououou/go-web-hello-world:v0.1
>docker login`
>docker push
```
> - Check pushed images on the website


<br/>


## Task 8: Document the procedure in a MarkDown file
**Skip**

<br/>
