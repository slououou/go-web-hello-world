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

## Task 9: Install a single node Kubernetes cluster using kubeadm

> **Steps:**

> - Disable swap
```
>free -m
```
>   - comment the swap line in `/etc/fstab` and reboot

> - Check prerequisite packages
```
>apt list --installed apt-transport-https curl
```
> - Install kubeadm kubectl kubelet
```
>apt update
>curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
cat <<EOF | sudo tee /etc/apt/sources.list.d/kubernetes.list
deb https://apt.kubernetes.io/ kubernetes-xenial main
EOF
>apt update
>apt install -y kubelet kubeadm kubectl
>apt-mark hold kubelet kubeadm kubectl
```
> - Install single master Kubernetes
>   - edit /lib/systemd/system/docker.service to add env
```
[Service]
Environment="http_proxy=http://proxy_url/" "https_proxy=http://proxy_url/"
```
>   - install kubernetes
```
>kubeadm config images pull v1.17.0
>unset http_proxy
>unset https_proxy
>kubeadm init --pod-network-cidr=192.168.0.0/16
>mkdir -p /root/.kube;cp /etc/kubernetes/admin.conf /root/.kube/config
>kubectl apply -f https://docs.projectcalico.org/v3.8/manifests/calico.yaml
>kubectl taint nodes --all node-role.kubernetes.io/master-
```
> - Check in `admin.conf`
```
>cd /root/golang/
>cp /etc/kubernetes/admin.conf .
>git add admin.conf
>git commit -m "admin.conf"
>git push
```
### Task 10: deploy the hello world container

> - Deploy app in cluster
>   - create demo.yaml for deployment, put below into it
```
---
apiVersion: v1
kind: Namespace
metadata:
  name: demo
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: go-web-hello-world
  namespace: demo
spec:
  replicas: 1
  selector:
    matchLabels:
      app: go-web-hello-world
  template:
    metadata:
      labels:
        app: go-web-hello-world
    spec:
      containers:
      - name: go-web-hello-world
        image: slououou/go-web-hello-world:v0.1
        ports:
        - containerPort: 80
---
apiVersion: v1
kind: Service
metadata:
  name: go-web-hello-world
  namespace: demo
  labels:
    app: go-web-hello-world
spec:
  type: NodePort
  ports:
  - port: 80
    nodePort: 31080
    targetPort: 80
  selector:
    app: go-web-hello-world
```
>   - deploy the app
```
kubectl apply -f demo.yaml
```
>   - open bowser on host and visit 127.0.0.1:31080
>   - get "Go Web Hello World!"

### Task 11: install kubernetes dashboard
> - Issue dashboard tls crt and key
>   - generate dashboard key
```
>apt install openssl
>openssl genrsa -out dashboard.key 2048
```
>   - put below in dashboard csr.conf
```
[ req ]
default_bits = 2048
prompt = no
default_md = sha256
req_extensions = req_ext
distinguished_name = dn

[ dn ]
C = CN
ST = ST
L = CI
O = O
OU = OU
CN = kubernetes-dashboard

[ req_ext ]
subjectAltName = @alt_names

[ alt_names ]
DNS.1 = kubernetes-dashboard
DNS.2 = kubernetes-dashboard.kubernetes-dashboard
DNS.3 = kubernetes-dashboard.kubernetes-dashboard.svc
DNS.4 = kubernetes-dashboard.kubernetes-dashboard.svc.cluster
DNS.5 = kubernetes-dashboard.kubernetes-dashboard.svc.cluster.local
IP.1 = 127.0.0.1
IP.2 = 10.0.2.15

[ v3_ext ]
authorityKeyIdentifier=keyid,issuer:always
basicConstraints=CA:FALSE
keyUsage=keyEncipherment,dataEncipherment
extendedKeyUsage=serverAuth,clientAuth
subjectAltName=@alt_names
```
>   - generate dashboard crt
```
openssl req -new -key dashboard.key -out dashboard.csr -config csr.conf
openssl x509 -req -in dashboard.csr -CA /etc/kubernetes/pki/ca.crt -CAkey /etc/kubernetes/pki/ca.key -CAcreateserial -out dashboard.crt -days 10000 -extensions v3_ext -extfile csr.conf
```
>   - put below in dashboard.yaml
```
apiVersion: v1
kind: ServiceAccount
metadata:
  labels:
    k8s-app: kubernetes-dashboard
  name: kubernetes-dashboard
  namespace: kubernetes-dashboard

---

kind: Service
apiVersion: v1
metadata:
  labels:
    k8s-app: kubernetes-dashboard
  name: kubernetes-dashboard
  namespace: kubernetes-dashboard
spec:
  ports:
    - port: 443
      targetPort: 8443
      nodePort: 31081
  selector:
    k8s-app: kubernetes-dashboard
  type: NodePort
---
apiVersion: v1
kind: Secret
metadata:
  labels:
    k8s-app: kubernetes-dashboard
  name: kubernetes-dashboard-key-holder
  namespace: kubernetes-dashboard
type: Opaque
---
apiVersion: v1
kind: Secret
metadata:
  labels:
    k8s-app: kubernetes-dashboard
  name: kubernetes-dashboard-csrf
  namespace: kubernetes-dashboard
type: Opaque
data:
  csrf: ""
---

kind: ConfigMap
apiVersion: v1
metadata:
  labels:
    k8s-app: kubernetes-dashboard
  name: kubernetes-dashboard-settings
  namespace: kubernetes-dashboard

---

kind: Role
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  labels:
    k8s-app: kubernetes-dashboard
  name: kubernetes-dashboard
  namespace: kubernetes-dashboard
rules:
  # Allow Dashboard to get, update and delete Dashboard exclusive secrets.
  - apiGroups: [""]
    resources: ["secrets"]
    resourceNames: ["kubernetes-dashboard-key-holder", "kubernetes-dashboard-certs", "kubernetes-dashboard-csrf"]
    verbs: ["get", "update", "delete"]
    # Allow Dashboard to get and update 'kubernetes-dashboard-settings' config map.
  - apiGroups: [""]
    resources: ["configmaps"]
    resourceNames: ["kubernetes-dashboard-settings"]
    verbs: ["get", "update"]
    # Allow Dashboard to get metrics.
  - apiGroups: [""]
    resources: ["services"]
    resourceNames: ["heapster", "dashboard-metrics-scraper"]
    verbs: ["proxy"]
  - apiGroups: [""]
    resources: ["services/proxy"]
    resourceNames: ["heapster", "http:heapster:", "https:heapster:", "dashboard-metrics-scraper", "http:dashboard-metrics-scraper"]
    verbs: ["get"]

---

kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  labels:
    k8s-app: kubernetes-dashboard
  name: kubernetes-dashboard
rules:
  # Allow Metrics Scraper to get metrics from the Metrics server
  - apiGroups: ["metrics.k8s.io"]
    resources: ["pods", "nodes"]
    verbs: ["get", "list", "watch"]

---

apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  labels:
    k8s-app: kubernetes-dashboard
  name: kubernetes-dashboard
  namespace: kubernetes-dashboard
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: kubernetes-dashboard
subjects:
  - kind: ServiceAccount
    name: kubernetes-dashboard
    namespace: kubernetes-dashboard

---

apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: kubernetes-dashboard
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: kubernetes-dashboard
subjects:
  - kind: ServiceAccount
    name: kubernetes-dashboard
    namespace: kubernetes-dashboard

---

kind: Deployment
apiVersion: apps/v1
metadata:
  labels:
    k8s-app: kubernetes-dashboard
  name: kubernetes-dashboard
  namespace: kubernetes-dashboard
spec:
  replicas: 1
  revisionHistoryLimit: 10
  selector:
    matchLabels:
      k8s-app: kubernetes-dashboard
  template:
    metadata:
      labels:
        k8s-app: kubernetes-dashboard
    spec:
      containers:
        - name: kubernetes-dashboard
          image: kubernetesui/dashboard:v2.0.0-beta8
          imagePullPolicy: Always
          ports:
            - containerPort: 8443
              protocol: TCP
          args:
            - --tls-cert-file=tls.crt
            - --tls-key-file=tls.key
            - --namespace=kubernetes-dashboard
            # Uncomment the following line to manually specify Kubernetes API server Host
            # If not specified, Dashboard will attempt to auto discover the API server and connect
            # to it. Uncomment only if the default does not work.
            # - --apiserver-host=http://my-address:port
          volumeMounts:
            - name: kubernetes-dashboard-certs
              mountPath: /certs
              # Create on-disk volume to store exec logs
            - mountPath: /tmp
              name: tmp-volume
          livenessProbe:
            httpGet:
              scheme: HTTPS
              path: /
              port: 8443
            initialDelaySeconds: 30
            timeoutSeconds: 30
          securityContext:
            allowPrivilegeEscalation: false
            readOnlyRootFilesystem: true
            runAsUser: 1001
            runAsGroup: 2001
      volumes:
        - name: kubernetes-dashboard-certs
          secret:
            secretName: kubernetes-dashboard-certs
        - name: tmp-volume
          emptyDir: {}
      serviceAccountName: kubernetes-dashboard
      nodeSelector:
        "beta.kubernetes.io/os": linux
      # Comment the following tolerations if Dashboard must not be deployed on master
      tolerations:
        - key: node-role.kubernetes.io/master
          effect: NoSchedule
---
```
>   - deploy kubernetes dashboard
```
kubectl create ns kubernetes-dashboard
kubectl create secret tls kubernetes-dashboard-certs --cert=dashboard.crt --key=dashboard.key -nkubernetes-dashboard
kubectl apply -f dashboard.yaml
```
>   - view dashboard in https://127.0.0.1:31801/ on host

### Task 12: generate token for dashboard login in task 11
> - put below in sa.yaml
```
apiVersion: v1
kind: ServiceAccount
metadata:
  name: admin-user
  namespace: kubernetes-dashboard
```
> - put below in crb.yaml
```
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: admin-user
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
- kind: ServiceAccount
  name: admin-user
  namespace: kubernetes-dashboard
```
> - get token from bash
```
kubectl apply -f sa.yaml
kubectl apply -f crb.yaml
kubectl -nkubernetes-dashboard describe secret $(kubectl describe serviceaccount -nkubernetes-dashboard admin-user | grep Tokens:| sed 's/Tokens:\s*//') |grep token: | awk  '{print $2}'
```
> - use this token to login kubernetes dashboard
