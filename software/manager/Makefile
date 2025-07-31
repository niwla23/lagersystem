visualize:
	go run -mod=mod ariga.io/entviz ./ent/schema 

generate:
	go generate ./ent

dev:
	OPERATOR_BASE_URL="http://192.168.178.64:5000" TYPESENSE_URI="http://localhost:8108" STORAGE_PATH=./data/images DB_URI=file:./data/db.sqlite?_fk=1 go run main.go

