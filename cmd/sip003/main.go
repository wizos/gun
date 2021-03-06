package main

import (
	"github.com/Qv2ray/gun/pkg/impl"
	"log"
)

func main() {
	log.Println("gun is running in SIP003 mode.")
	arguments, err := GetSIP003Arguments()
	if err != nil {
		log.Fatalf("failed to parse sip003 arguments: %v", err)
	}

	options, err := ParsePluginOptions(arguments.Options)
	if err != nil {
		log.Fatalf("failed to parse plugin options: %v", err)
	}

	switch options["mode"] {
	case "client":
		impl.GunServiceClientImpl{
			RemoteAddr: arguments.RemoteAddr,
			LocalAddr:  arguments.LocalAddr,
			ServerName: options["sni"],
		}.Run()
	case "server":
		impl.GunServiceServerImpl{
			RemoteAddr: arguments.RemoteAddr,
			LocalAddr:  arguments.LocalAddr,
			CertPath:   options["cert"],
			KeyPath:    options["key"],
		}.Run()
	default:
		log.Fatalf("unknown run mode")
	}
}
