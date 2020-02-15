GEN := ./gen
SPEC := ./swagger/swagger.yml
APP_NAME := wire-commander

.PHONY: validate
validate:
	@swagger validate $(SPEC)

.PHONY: generate
generate: validate
	@mkdir -p $(GEN)
	@swagger generate server --spec=$(SPEC) --target=$(GEN) --exclude-main --name=$(APP_NAME)

.PHONY: clean
clean:
	@rm -rf $(GEN)

