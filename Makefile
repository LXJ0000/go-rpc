DIR = $(shell pwd)/app

CONFIG_PATH = $(shell pwd)
IDL_PATH = $(shell pwd)/idl

SERVICES := user
service = $(word 1, $@)

node = 0

BIN = $(shell pwd)/bin



.PHONY: proto
proto:
	@for file in $(IDL_PATH)/*.proto; do \
		protoc -I $(IDL_PATH) $$file --go-grpc_out=$(IDL_PATH)/pb --go_out=$(IDL_PATH)/pb; \
	done

.PHONY: $(SERVICES)
$(SERVICES):
	go build -o $(BIN)/$(service) $(DIR)/$(service)/cmd
	$(BIN)/$(service) -config $(CONFIG_PATH) -srvnum=$(node)