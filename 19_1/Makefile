up:
	docker compose up --build -d

down:
	docker stop --time=0 $$(docker ps -aq)
	docker system prune -f
	docker volume prune -f

install_report:
	go get golang.org/x/sync/errgroup
	go get google.golang.org/grpc
	go get google.golang.org/grpc/codes
	go get google.golang.org/grpc/status
	go get google.golang.org/protobuf/reflect/protoreflect
	go get google.golang.org/protobuf/runtime/protoimpl

run_report:
	make install_report
	tail -f /dev/null
