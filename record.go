package ipinfo_io_lite

import "net/netip"

type Record struct {
	ContinentCode string `json:"continent_code" maxminddb:"continent"`
	Continent     string `json:"continent" maxminddb:"continent_name"`
	CountryCode   string `json:"country_code" maxminddb:"country"`
	Country       string `json:"country" maxminddb:"country_name"`
	ASN           string `json:"asn" maxminddb:"asn"`
	AsName        string `json:"as_name" maxminddb:"as_name"`
	AsDomain      string `json:"as_domain" maxminddb:"as_domain"`
}

type DbRecord struct {
	Prefix netip.Prefix `json:"prefix" maxminddb:"prefix"`
	Record
}

type ApiRecord struct {
	IP netip.Addr `json:"ip"`
	Record
}
