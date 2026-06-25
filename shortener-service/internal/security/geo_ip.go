package security

import (
	"github.com/oschwald/geoip2-golang"
	"net"
)

type GeoIP interface {
	Country(
		ip string,
	) string

	Close() error
}

type MaxMindGeoIP struct {
	db *geoip2.Reader
}

func NewGeoIP(
	path string,
) (
	*MaxMindGeoIP,
	error,
) {

	db, err := geoip2.Open(
		path,
	)

	if err != nil {
		return nil, err
	}

	return &MaxMindGeoIP{
		db: db,
	}, nil
}

func (g *MaxMindGeoIP) Country(
	ip string,
) string {

	parsedIP := net.ParseIP(
		ip,
	)

	if parsedIP == nil {
		return ""
	}

	record, err := g.db.Country(
		parsedIP,
	)

	if err != nil {
		return ""
	}

	return record.Country.Names["en"]
}

func (g *MaxMindGeoIP) Close() error {

	if g.db == nil {
		return nil
	}

	return g.db.Close()
}
