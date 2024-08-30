build:
	go mod tidy
	go build -o ${PWD}/bin/passwdgen cmd/main.go

install: build
	sudo cp ${PWD}/bin/passwdgen /usr/local/bin/
	chmod +x /usr/local/bin/passwdgen
