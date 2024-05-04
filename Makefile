# Manually define main variables

ifndef APP_PORT
export APP_PORT = 8080
endif

ifndef APP_HOST
export APP_HOST = 127.0.0.1
endif

# parse additional args for commands

ROOT_FOLDER = .

args = `arg="$(filter-out $@,$(MAKECMDGOALS))" && echo $${arg:-${1}}`

run:  ##@Application Run application server
	go run $(ROOT_FOLDER)/cmd/golang_graphs --rootPath $(ROOT_FOLDER)

run_test:  ##@Application Run application server
	TESTING="testing" go run $(ROOT_FOLDER)/cmd/golang_graphs --rootPath $(ROOT_FOLDER)

create_swagger:
	swag init -g cmd/golang_graphs/main.go

# Это просто пример того как запускать бенчмарки, если запустить из makefile, то он не отработает
bench:
	go test -bench=BenchmarkSimplest -benchmem -benchtime=1x

docker_run:
	set -ex
	sudo docker-compose -f $(ROOT_FOLDER)/deploy/docker-compose.yaml -p deploy up --build -d graphs_db
	sudo docker-compose -f $(ROOT_FOLDER)/deploy/docker-compose.yaml -p deploy up --build -d node-exporter
	sudo docker-compose -f $(ROOT_FOLDER)/deploy/docker-compose.yaml -p deploy up --build -d prometheus
	sudo docker-compose -f $(ROOT_FOLDER)/deploy/docker-compose.yaml -p deploy up --build -d grafana
	sudo docker-compose -f $(ROOT_FOLDER)/deploy/docker-compose.yaml -p deploy up --build -d graphs_back
### Писать в формате FunctionName (CamelCase и первая буква заглавная)
### Например CreateUser
#create_file_handler_common:
#	python3 $(ROOT_FOLDER)/internal/handler/common/create_file.py $(call args) $(ROOT_FOLDER)/internal/handler/common/
