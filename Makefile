all:
	GOOS=windows GOARCH=amd64 go build  -o bin/beetle.exe
	GOOS=windows GOARCH=386 go build -o bin/beetle32.exe