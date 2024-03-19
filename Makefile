// https://github.com/open-telemetry/opentelemetry-demo/blob/main/Makefile
	
.PHONY: start
start:
	docker compose up --force-recreate --remove-orphans --detach
	@echo ""
	@echo "Online Boutique is running."
	@echo "Go to http://localhost:8080 for the demo UI."
	@echo "Go to http://localhost:8080/jaeger/ui for the Jaeger UI."
	@echo "Go to http://localhost:8080/grafana/ for the Grafana UI."
	@echo "Go to http://localhost:8080/loadgen/ for the Load Generator UI."
	@echo "Go to http://localhost:8080/feature/ for the Feature Flag UI."

.PHONY: stop
stop:
	docker compose --profile tests --profile odd down --remove-orphans --volumes
	@echo ""
	@echo "Online Boutique is stopped."

# Use to restart a single service component
# Example: make restart service=frontend
.PHONY: restart
restart:
# work with `service` or `SERVICE` as input
ifdef SERVICE
	service := $(SERVICE)
endif

ifdef service
	docker compose stop $(service)
	docker compose rm --force $(service)
	docker compose create $(service)
	docker compose start $(service)
else
	@echo "Please provide a service name using `service=[service name]` or `SERVICE=[service name]`"
endif
