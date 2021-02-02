package settings

import (
	"golang.org/x/sys/windows/registry"
	"log"
)

type WinProxy struct{}

func (proxy *WinProxy) SetProxy(mode int) {
	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Internet Settings`, registry.ALL_ACCESS)
	if err != nil {
		log.Println("can't open registry", err.Error())
		return
	}
	defer k.Close()

	proxy.ClearProxy()

	if mode == 0 {
		err := k.SetStringValue("AutoConfigURL", "http://127.0.0.1:50211/web/gfw.js")
		if err != nil {
			log.Println("can't set gfw.js on registry", err.Error())
		}

	} else {
		if err := k.SetDWordValue("ProxyEnable", 1); err != nil {
			log.Println("can't set ProxyEnable to 1", err.Error())
		}

		if err := k.SetStringValue("ProxyServer", "localhost:40020"); err != nil {
			log.Println("can't set ProxyServer", err.Error())
		}
	}
}

func (proxy *WinProxy) ClearProxy() {
	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Internet Settings`, registry.ALL_ACCESS)
	if err != nil {
		log.Println("can't open registry", err.Error())
		return
	}
	defer k.Close()

	if err := k.DeleteValue("AutoConfigURL"); err != nil {
		//fmt.Print("can't find key AutoConfigURL", err.Error())
		//do nothing
	}

	if err := k.SetDWordValue("ProxyEnable", 0); err != nil {
		log.Print("can't set ProxyEnable to 0", err.Error())
	}
}

