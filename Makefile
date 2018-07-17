BIN_DIR=_output/bin
RELEASE_VER=0.1

init:
	mkdir -p ${BIN_DIR}

generate-code:
	go build -o ${BIN_DIR}/deepcopy-gen ./cmd/deepcopy-gen/
	${BIN_DIR}/deepcopy-gen -i github.com/xieydd/kubenetes-crd/pkg/apis/v1alpha/ --go-header-file="C:/Users/MyPC/go/src/github.com/xieydd/kubenetes-CRD/hack/boilerplate/boilerplate.go.txt" -v=4 --logtostderr -O zz_generated.deepcopy

clean:
	rm -rf _output/
	rm -f kube-arbitrator

vender:
	go run C:/Users/MyPC/go/src/github.com/kardianos/govendor/main.go update +e

