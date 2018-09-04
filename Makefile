BIN_DIR=_output/bin
RELEASE_VER=0.1
HOME=/home/xieyd

init:
	mkdir -p ${BIN_DIR}
	rm -rf pkg/client
	bash ${GOPATH}/src/github.com/xieydd/kubenetes-crd/hack/generate-groups.sh "all" github.com/xieydd/kubenetes-crd/pkg/client  github.com/xieydd/kubenetes-crd/pkg/apis unisound.org:v1alpha --go-header-file=${GOPATH}/src/github.com/xieydd/kubenetes-crd/hack/boilerplate.go.txt

generate-code:
	go build -o ${BIN_DIR}/deepcopy-gen ./cmd/deepcopy-gen/
	${BIN_DIR}/deepcopy-gen -i github.com/xieydd/kubenetes-crd/pkg/apis/v1alpha/ --go-header-file=${GOPATH}/src/github.com/xieydd/kubenetes-crd/hack/boilerplate.go.txt -v=4 --logtostderr -O zz_generated.deepcopy

clean:
	rm -rf _output/
	rm -f kube-arbitrator
	rm -rf vendor/

vender-init:
	go run ${GOPATH}/src/github.com/kardianos/govendor/main.go init
	go run ${GOPATH}/src/github.com/kardianos/govendor/main.go add +e

vender-update:
	go run ${GOPATH}/src/github.com/kardianos/govendor/main.go update +e

crd-test:
	go run ${GOPATH}/src/github.com/xieydd/kubenetes-crd/test/kube-crd.go --kubeconfig=${HOME}/.kube/config

