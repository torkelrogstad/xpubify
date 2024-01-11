name := "xpubify"

test: 
    go test .

build: 
    go build -o {{ name }} .

install: 
    go install .
