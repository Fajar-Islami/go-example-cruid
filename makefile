git-add:
	git add .
	git commit -am '${cmt}'

test:
	echo ${cmt}

entermysql:
	docker exec -it mysql_fiber_gorm_example mysql -u ADMIN -pSECRET rakamin_intern

entermysqlroot:
	docker exec -it mysql_fiber_gorm_example mysql -u root -pSECRET_ROOT rakamin_intern

runenv:
	docker compose up -d

run:
	docker compose up -d
	go run app/main.go

commit:
	git add .
	git commit -am '${cmt}'

struct:
	gomodifytags -file ${file} -struct ${struct} -add-tags ${tags}

stop:
	docker compose stop

down:
	docker compose down -v

logs:
	docker compose logs -f

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o  ./dist/example ./app/main.go

dockerbuild:
	docker build --rm -t example_fiber .
	docker image prune --filter label=stage=dockerbuilder -f

dockerun:
	docker run --name example_fiber  -p 8080:8080 example_fiber 

dockerrm:
	docker rm example_fiber -f
	docker rmi example_fiber

dockeenter:
	docker exec -it example_fiber bash