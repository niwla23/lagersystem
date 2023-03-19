visualize:
	go run -mod=mod ariga.io/entviz ./ent/schema 

generate:
	go generate ./ent

dev:
	STORAGE_PATH=/tmp DB_URI=file:///tmp/db.sqlite?_fk=1 go run main.go