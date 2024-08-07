.PHONY: *

#*------------VARS------------

#*______URL______

server = http://127.0.0.1:8080
github = https://github.com/bastean/dsgo

#*______Go______

go-tidy = go mod tidy -e

#*______Node______

npx = npx --no --
npm-ci = npm ci --legacy-peer-deps

release-it = ${npx} release-it -V
release-it-dry = ${npx} release-it -V -d --no-git.requireCleanWorkingDir

#*______Bash______

bash = bash -o pipefail -c

#*______Git______

git-reset-hard = git reset --hard HEAD

#*______Docker______

compose = cd deployments/ && docker compose
compose-env = ${compose} --env-file

#*------------RULES------------

#*______Upgrades______

upgrade-managers:
	#? sudo apt update && sudo apt upgrade -y
	npm upgrade -g

upgrade-go:
	go get -t -u ./...

upgrade-node:
	${npx} ncu --root -ws -u
	rm -f package-lock.json
	npm i --legacy-peer-deps

upgrade-reset:
	${git-reset-hard}
	${npm-ci}

upgrade:
	go run ./scripts/upgrade

#*______Tools______

scan-tools:
	curl -sSfL https://raw.githubusercontent.com/trufflesecurity/trufflehog/main/scripts/install.sh | sudo sh -s -- -b /usr/local/bin v3.63.11
	curl -sSfL https://raw.githubusercontent.com/aquasecurity/trivy/main/contrib/install.sh | sudo sh -s -- -b /usr/local/bin v0.52.2
	go install github.com/google/osv-scanner/cmd/osv-scanner@latest

lint-tools:
	go install honnef.co/go/tools/cmd/staticcheck@latest

dev-tools:
	go install github.com/air-verse/air@latest
	go install github.com/a-h/templ/cmd/templ@latest
	npm i -g prettier

test-tools:
	go run github.com/playwright-community/playwright-go/cmd/playwright@latest install chromium --with-deps
	npm i -g concurrently wait-on

install-tools: scan-tools lint-tools dev-tools test-tools

#*______Dependencies______

install-deps:
	go mod download
	${npm-ci}

copy-deps:
	go run ./scripts/copydeps

#*______Generators______

generate-required:
	go generate ./...
	find . -name "*_templ.go" -type f -delete
	templ generate

#*______Initializations______

init: upgrade-managers install-tools install-deps copy-deps generate-required

init-zero:
	git init
	$(MAKE) init
	${npx} husky install

#*______Linters/Formatters______

lint: generate-required
	go mod tidy
	gofmt -l -s -w .
	${npx} prettier --ignore-unknown --write .
	templ fmt .

lint-check:
	staticcheck ./...
	${npx} prettier --check .

#*______Scanners______

leak-check:
	sudo trufflehog git file://. --only-verified
	trivy repo --scanners secret .

leak-remote-check:
	sudo trufflehog git ${github} --only-verified
	trivy repo --scanners secret ${github}

vuln-check:
	osv-scanner --call-analysis=all -r .
	trivy repo --scanners vuln .

misconfig-check:
	trivy repo --scanners misconfig .

scan-leaks: leak-check leak-remote-check

scan-vulns: vuln-check

scan-misconfigs: misconfig-check

scans: scan-leaks scan-vulns scan-misconfigs

#*______Tests______

test-sut:
	air

test-clean: generate-required
	go clean -testcache
	cd test/ && mkdir -p report

test-codegen:
	${npx} playwright codegen ${server}

test-sync:
	${npx} concurrently -s first -k --names 'SUT,TEST' '$(MAKE) test-sut' '${npx} wait-on -l ${server} && $(TEST_SYNC)'

test-unit: test-clean
	${bash} 'go test -v -cover ./pkg/context/... -run TestUnit.* |& tee test/report/unit.report.log'

test-integration: test-clean
	${bash} 'go test -v -cover ./pkg/context/... -run TestIntegration.* |& tee test/report/integration.report.log'

test-acceptance-sync: 
	${bash} 'TEST_URL="${server}" go test -v -cover ./internal/app/... -run TestAcceptance.* |& tee test/report/acceptance.report.log'

test-acceptance: test-clean
	TEST_SYNC="$(MAKE) test-acceptance-sync" $(MAKE) test-sync

tests-sync:
	${bash} 'TEST_URL="${server}" go test -v -cover ./... |& tee test/report/report.log'

tests: test-clean
	TEST_SYNC="$(MAKE) tests-sync" $(MAKE) test-sync

#*______Releases______

release:
	${release-it}

release-alpha:
	${release-it} --preRelease=alpha
	
release-beta:
	${release-it} --preRelease=beta

release-ci:
	${release-it} --ci --no-git.requireCleanWorkingDir $(OPTIONS)

release-dry:
	${release-it-dry}
 
release-dry-version:
	${release-it-dry} --release-version

release-dry-changelog:
	${release-it-dry} --changelog

#*______Builds______

build: generate-required lint
	rm -rf build/
	go build -ldflags="-s -w" -o build/dsgo ./cmd/dsgo

#*______ENVs______

sync-env-reset:
	${git-reset-hard}

sync-env:
	cd deployments && go run ../scripts/syncenv

#*______Git______

commit:
	${npx} cz

WARNING-git-forget:
	git rm -r --cached .
	git add .

WARNING-git-genesis:
	git clean -e .env* -fdx
	${git-reset-hard}
	$(MAKE) init

#*______Docker______

docker-usage:
	docker system df

docker-it:
	docker exec -it $(ID) bash

compose-dev-down:
	${compose-env} .env.dev down
	docker volume rm -f dsgo-database-mysql-dev

compose-dev: compose-dev-down
	${compose-env} .env.dev up

compose-test-down:
	${compose-env} .env.test down
	docker volume rm -f dsgo-database-mysql-test

compose-test-integration: compose-test-down
	${compose-env} .env.test --env-file .env.test.integration up --exit-code-from dsgo

compose-test-acceptance: compose-test-down
	${compose-env} .env.test --env-file .env.test.acceptance up --exit-code-from dsgo

compose-tests: compose-test-down
	${compose-env} .env.test up --exit-code-from dsgo

compose-prod-down:
	${compose-env} .env.prod down

compose-prod: compose-prod-down
	${compose-env} .env.prod up

demo-down:
	${compose-env} .env.demo down

demo: demo-down
	${compose-env} .env.demo up

compose-down: compose-dev-down compose-test-down compose-prod-down demo-down

WARNING-docker-prune-soft:
	docker system prune
	$(MAKE) compose-down
	$(MAKE) docker-usage

WARNING-docker-prune-hard:
	docker system prune --volumes -a
	$(MAKE) compose-down
	$(MAKE) docker-usage

#*______Fixes______

fix-playwright: upgrade-go tools-test
