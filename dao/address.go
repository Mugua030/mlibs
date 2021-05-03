package dao

import (
	"database/sql"
	"github.com/pkg/errors"
)

type Address struct {
	AddrId int64
	Country string
	CountryCode int32
	HomeAddr string
}

var (
	AddrId int64
	Country string
	CountryCode int32
	HomeAddr string
)

func GetDefaultAddr(uid int64) (*Address, error) {
	err := Db.QueryRow("select addrid,country,country_code,home_addr from address where uid= ? limit 1", uid).Scan(&AddrId, &Country, &CountryCode, &HomeAddr)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.Wrap(err, "dao: db query addr fail")
	}

	addr := &Address{
		AddrId: AddrId,
		Country: Country,
		CountryCode: CountryCode,
		HomeAddr: HomeAddr,
	}
	return addr, nil
}
