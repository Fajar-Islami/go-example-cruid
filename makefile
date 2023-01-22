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