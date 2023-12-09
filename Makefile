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

create_swagger:
	swag init -g cmd/golang_graphs/main.go


## Писать в формате FunctionName (CamelCase и первая буква заглавная)
## Например CreateUser
create_file_handler_common:
	python3 $(ROOT_FOLDER)/internal/handler/common/create_file.py $(call args) $(ROOT_FOLDER)/internal/handler/common/
