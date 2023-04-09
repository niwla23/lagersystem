visualize:
	go run -mod=mod ariga.io/entviz ./ent/schema 

generate:
	go generate ./ent

dev:
	STORAGE_PATH=./data/images DB_URI=file:///home/alwin/Dokumente/projects/lagersystem/software/manager/data/db.sqlite?_fk=1 go run main.go