run:
	go build -o build/md2json main.go
	./build/md2json convert data/test.md
