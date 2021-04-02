start: ## Starts up the container
	docker-compose up

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

infra: ## Starts the tools shell
	docker-compose -f docker-compose-infra.yml run --rm shell

decrypt: ## Decrypts files in the .config directory. Note: sops must be downloaded
	sops -d .config/nonprod.enc.yaml > .config/nonprod.dec.yaml
	sops -d .config/prod.enc.yaml > .config/prod.dec.yaml

encrypt: ## Encrypts files in the .config directory. Note: sops must be downloaded
	sops -e --config .config/.sops.yaml --encrypted-suffix enc .config/nonprod.dec.yaml > .config/nonprod.enc.yaml
	sops -e --config .config/.sops.yaml --encrypted-suffix enc .config/prod.dec.yaml > .config/prod.enc.yaml