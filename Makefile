include include.mk

${PROJECT_NAME}: main.go decrypt.go token.go ${PROJECT_NAME}.darwin_amd64
	go get
	go build

${PROJECT_NAME}${PACKAGE_TYPE}-${PACKAGE_VERSION}.darwin_amd64: main.go decrypt.go token.go
	wget https://github.com/jamesandariese/gccc/releases/download/v0.2/gccc
	chmod +x gccc
	(eval `./gccc darwin/amd64 darwingocc` ; go get ; go build -o ${PROJECT_NAME}${PACKAGE_TYPE}-${PACKAGE_VERSION}.darwin_amd64)

${PROJECT_NAME}.1.gz: README.md
	rvm install 2.2.2
	rvm 2.2.2 do gem install ronn
	grep -vE '^\[!\[Build Status\]' README.md | rvm 2.2.2 do ronn |gzip -9> ${PROJECT_NAME}.1.gz
	grep -vE '^\[!\[Build Status\]' README.md | rvm 2.2.2 do ronn -m | cat
