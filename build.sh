env GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o largeFileGenerator_darvin .
env GOOS=linux GOARCH=386 go build -ldflags="-s -w" -o largeFileGenerator_x86 .
env GOOS=windows GOARCH=386 go build -ldflags="-s -w" -o largeFileGenerator_win .
