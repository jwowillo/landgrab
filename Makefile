all:
	cd app/cli && go build -o run_landgrab_cli
	cd app/cli && mv run_landgrab_cli $$GOPATH/bin
