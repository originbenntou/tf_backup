PWD=$(shell pwd)
TS=$(shell date +%s)


# Docker build \
プライベートリポジトリDownloadのためアクセス権限を持つGithubアクセストークンが必要 \
事前に $ export GITHUB_ACCESS_TOKEN=発行トークン で環境変数にセット \
とは言えプッシュをトリガーにCloudBuildされるのでbuildコマンドを使うことはない...

## e.g. $ make build_all GITHUB_ACCESS_TOKEN=${GITHUB_ACCESS_TOKEN}
build_all:
	@docker build ./account --build-arg GITHUB_ACCESS_TOKEN=${GITHUB_ACCESS_TOKEN}
	@docker build ./trend --build-arg GITHUB_ACCESS_TOKEN=${GITHUB_ACCESS_TOKEN}
	@docker build ./gateway --build-arg GITHUB_ACCESS_TOKEN=${GITHUB_ACCESS_TOKEN}

## e.g. $ make build_one GITHUB_ACCESS_TOKEN=${GITHUB_ACCESS_TOKEN} SVC=account
build_one:
	@docker build ./${SVC} --build-arg GITHUB_ACCESS_TOKEN=${GITHUB_ACCESS_TOKEN}

# Generate Proto

## e.g. $ make gen_go_proto SVC=account
gen_go_proto:
	protoc --proto_path=${GOPATH}/src --proto_path=. --go_out=plugins=grpc:${GOPATH}/src --govalidators_out=${GOPATH}/src ./proto/${SVC}/*.proto

## e.g. $ make gen_py_proto SVC=trend
gen_py_proto:
	python -m grpc_tools.protoc -I ./proto/${SVC}/ --python_out=./proto/${SVC}/python --grpc_python_out=./proto/${SVC}/python ./proto/${SVC}/${SVC}.proto

# Migrate DB \
ローカル開発専用のコマンド群 \
migrate_up,migrate_downはmysqlコンテナが起動しているときに実行

## e.g. $ make migrate_init TITLE=alter_user_table
migrate_init:
	docker run -v ${PWD}/mysql/migrations:/migrations \
		migrate/migrate create -ext sql -dir migrations -seq ${TITLE}

## e.g. $ make migrate_up DB=account VERSION=000001
migrate_up:
	docker run -v ${PWD}/mysql/migrations:/migrations \
		--network tf_backend_default migrate/migrate \
		-path=/migrations/ \
		-database "mysql://2929:2929@tcp(mysql:3306)/${DB}?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true" up ${VERSION}

## e.g. $ make migrate_down DB=account VERSION=000001
migrate_down:
	docker run -v ${PWD}/mysql/migrations:/migrations \
		--network tf_backend_default migrate/migrate \
		-path=/migrations/ \
		-database "mysql://2929:2929@tcp(mysql:3306)/${DB}?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true" down ${VERSION}
