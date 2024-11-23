SWAGGER_UI_IMAGE = swaggerapi/swagger-ui
SWAGGER_UI_PORT = 8200
OPENAPI_FILE = openapi.yaml

.PHONY: swagger-ui

swagger-ui:
	@echo "Запуск Swagger UI на http://localhost:$(SWAGGER_UI_PORT)"
	@docker run --rm -p $(SWAGGER_UI_PORT):8080 \
		-e SWAGGER_JSON=/tmp/$(OPENAPI_FILE) \
		-v $(CURDIR)/$(OPENAPI_FILE):/tmp/$(OPENAPI_FILE) \
		$(SWAGGER_UI_IMAGE)
