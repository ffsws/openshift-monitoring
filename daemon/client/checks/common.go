package checks

import (
	"net"
	"log"
	"net/http"
	"crypto/tls"
	"strings"
)

const (
	daemonDNSEndpoint = "daemon.ose-mon-a.endpoints.cluster.local"
	daemonDNSServiceA = "daemon.ose-mon-a.svc.cluster.local"
	daemonDNSServiceB = "daemon.ose-mon-b.svc.cluster.local"
	daemonDNSServiceC = "daemon.ose-mon-c.svc.cluster.local"
	daemonDNSPod = "daemon"
	kubernetesIP = "172.30.0.1"
)

func getIpsForName(n string) []net.IP {
	ips, err := net.LookupIP(n)
	if (err != nil) {
		log.Println("failed to lookup ip for name ", n)
		return nil
	}
	return ips
}

func checkHttp(toCall string) bool {
	if (strings.HasPrefix(toCall, "https")) {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		resp, err := client.Get(toCall)
		if (err != nil) {
			log.Println("error in http check: ", err)
		} else {
			resp.Body.Close()
		}
		return err == nil
	} else {
		resp, err := http.Get(toCall)
		if (err != nil) {
			log.Println("error in http check: ", err)
		} else {
			resp.Body.Close()
		}
		return err == nil
	}
}

func getEndpoint(slow bool) string {
	if (slow) {
		return "slow"
	} else {
		return "fast"
	}
}