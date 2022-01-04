proto:
	docker run -u `id -u`:`id -g` --rm -v /Users/aakashgoyal/go/src/github.com/Infoblox-CTO/goQuickOverview/grpcexample/pkg:/go/src/github.com/aakashgyl/goQuickOverview/grpcexample/pkg -e XDG_CACHE_HOME=/tmp/.cache infoblox/atlas-gentool:v20 --go_out="plugins=grpc:." --preprocess_out="." --validate_out="lang=go:." --grpc-gateway_out=logtostderr=true:. github.com/aakashgyl/goQuickOverview/grpcexample/pkg/pb/service.proto

help:
	@echo "this has http and grpc example..."
