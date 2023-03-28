export UDIR= .
export GOC = x86_64-xen-ethos-6g
export GOL = x86_64-xen-ethos-6l
export ETN2GO = etn2go
export ET2G   = et2g
export EG2GO  = eg2go

export GOARCH = amd64
export TARGET_ARCH = x86_64
export GOETHOSINCLUDE=ethos
export GOLINUXINCLUDE=linux
export BUILD=ethos

export ETHOSROOT=client/rootfs
export MINIMALTDROOT=client/minimaltdfs


.PHONY: all install clean
all: syncClient syncService

ethos:
	mkdir ethos
	cp -pr /usr/lib64/go/pkg/ethos_$(GOARCH)/* ethos

myRpc.go: myRpc.t
	$(ETN2GO) . myRpc $^

myRpc.goo.ethos : myRpc.go ethos
	ethosGoPackage  myRpc ethos myRpc.go

syncService: banking-server.go myRpc.goo.ethos
	ethosGo banking-server.go

syncClient: banking-client.go myRpc.goo.ethos
	ethosGo banking-client.go

# install types, service,
install: all
	sudo rm -rf client
	(ethosParams client && cd client && ethosMinimaltdBuilder)
	echo 7 > client/param/sleepTime
	ethosTypeInstall myRpc
	ethosServiceInstall myRpc
	ethosDirCreate $(ETHOSROOT)/services/myRpc   $(ETHOSROOT)/types/spec/myRpc/MyRpc all
	install -D  syncClient syncService                   $(ETHOSROOT)/programs
	ethosStringEncode /programs/syncService    > $(ETHOSROOT)/etc/init/services/syncService

# remove build artifacts
clean:
	rm -rf myRpc/ myRpcIndex/ ethos client
	rm -f myRpc.go
	rm -f syncClient
	rm -f syncService
	rm -f myRpc.goo.ethos
	rm -f banking-client.goo.ethos
	rm -f banking-server.goo.ethos

