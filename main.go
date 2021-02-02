package main

import (
	"github.com/giantliao/beatles-client-lib/app/cmd"
	"github.com/giantliao/beatles-client-lib/config"
	"github.com/giantliao/beatles-client-lib/webmain"
	"github.com/giantliao/beatles-win-client/settings"

	"net/http"
)

func main()  {

	if _, err := http.Get("http://127.0.0.1:50211"); err == nil {
		webmain.OpenBrowser("http://127.0.0.1:50211")
		return
	}

	cmd.InitCfg()
	cfg := config.GetCBtlc()
	cfg.Save()

	webmain.StartWEBService(&settings.WinProxy{})
}
