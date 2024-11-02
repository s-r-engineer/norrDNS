package main

import (
	libraryNordvpn "github.com/s-r-engineer/library/nordvpn"
	"strings"

	"github.com/miekg/dns"
	libraryErrors "github.com/s-r-engineer/library/errors"
	libraryStrings "github.com/s-r-engineer/library/strings"
)

func handleDNSRequest(w dns.ResponseWriter, r *dns.Msg) {
	msg := dns.Msg{}
	msg.SetReply(r)
	q := queryHistory{}
	if len(r.Question) > 0 {
		domain := r.Question[0].Name
		q.QueriedAddress = domain
		splitted := strings.Split(domain, ".")
		countryID := countrieExist(splitted[0])
		if compareSlices(splitted[1:], FQDNBase) && countryID != -1 {
			hostname, _, _, _, err := libraryNordvpn.FetchServerData(countryID)
			libraryErrors.Errorer(err)
			a := &dns.CNAME{
				Hdr:    dns.RR_Header{Name: domain, Rrtype: dns.TypeCNAME, Class: dns.ClassINET, Ttl: 10},
				Target: hostname + ".",
			}
			q.ResolvedAddress = hostname
			msg.Answer = append(msg.Answer, a)
		} else {
			a := &dns.TXT{
				Hdr: dns.RR_Header{Name: domain, Rrtype: dns.TypeTXT, Class: dns.ClassINET, Ttl: 66666},
				Txt: []string{libraryStrings.RandString(6)},
			}
			q.ResolvedAddress = "fuck you"
			msg.Answer = append(msg.Answer, a)
		}
	}
	add(q)
	libraryErrors.Errorer(w.WriteMsg(&msg))
}
