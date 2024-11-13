
.PHONY: all run

all:
	rm -rf bin
	mkdir -p bin
	GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o bin/occupation_mac_amd64 occupation.go
	GOOS=windows GOARCH=amd64 go build -ldflags="-H windowsgui -s -w" -o bin/occupation_win_amd64.exe occupation.go
	GOOS=windows GOARCH=arm64 go build -ldflags="-H windowsgui -s -w" -o bin/occupation_win_arm64.exe occupation.go
	cd bin && mv occupation_mac_amd64 楓之谷職業隨機選擇器 && zip -P "Et^v?kf.*zUp6Q>" -r occupation_mac_amd64.zip 楓之谷職業隨機選擇器
	cd bin && mv occupation_win_amd64.exe 楓之谷職業隨機選擇器.exe && zip -P "Et^v?kf.*zUp6Q>" -r occupation_win_amd64.zip 楓之谷職業隨機選擇器.exe
	cd bin && mv occupation_win_arm64.exe 楓之谷職業隨機選擇器.exe && zip -P "Et^v?kf.*zUp6Q>" -r occupation_win_arm64.zip 楓之谷職業隨機選擇器.exe
	cd bin && rm -f 楓之谷職業隨機選擇器*

run:
	go run -ldflags="-s -w" occupation.go
