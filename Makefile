test: clean TestGetOnlyAppCfg TestGetFullCfg TestGetSpecifiedEnvCfg

TestGetOnlyAppCfg: clean
	cp -rf fixtures/configs1 configs
	go test -gocheck.f TestGetOnlyAppCfg
	rm -rf configs

TestGetFullCfg: clean
	cp -rf fixtures/configs2 configs
	go test -gocheck.f TestGetFullCfg
	rm -rf configs

TestGetSpecifiedEnvCfg: clean
	cp -rf fixtures/configs2 configs
	go test -gocheck.f TestGetSpecifiedEnvCfg -parkour.env prod
	rm -rf configs

clean:
	rm -rf configs
	go build
