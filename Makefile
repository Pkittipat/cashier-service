setup:
	mkdir cmd pkg
	mkdir cmd/server
	touch cmd/server/{server,handler,inventory,payment,tree,validate}.go
	mkdir pkg/{inventory,payment,server}
	touch pkg/inventory/{inventory,tree}.go
	touch pkg/payment/{payment,validate}.go
	touch pkg/server/{server,handler,middleware}.go
