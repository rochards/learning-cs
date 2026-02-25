DBUSER ?= appuser
DBPASS ?= apppass
export DBUSER DBPASS

.PHONY: run stop

run:
	docker compose up -d
	@echo "Waiting for MySQL to be healthy..."
	@for i in $$(seq 1 60); do \
		status=$$(docker inspect --format '{{.State.Health.Status}}' recordings-mysql 2>/dev/null || true); \
		if [ "$$status" = "healthy" ]; then \
			break; \
		fi; \
		if [ $$i -eq 60 ]; then \
			echo "MySQL did not become healthy within 60s (last status: $$status)"; \
			exit 1; \
		fi; \
		sleep 1; \
	done
	go run .

stop:
	docker compose down -v
