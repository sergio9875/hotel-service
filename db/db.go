package db

import (
	"database/sql"
	"fmt"
)

const (
	DbName        = "hotel-reservation"
	TestDbName    = "hotel-reservation-test"
	DbUri         = "mongodb://localhost:27017"
	MySqlUser     = "root"
	MySqlPassword = "password"
	MySqlHost     = "localhost"
	MySqlDbName   = "test_Db"
)

type Store struct {
	User    UserStore
	Hotel   HotelStore
	Room    RoomStore
	Booking BookingStore
}

func ConnectToMySql(user string, password string, host string, dbname string) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, host, dbname)
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		err := db.Close()
		if err != nil {
			return nil, err
		}
		return nil, err
	}
	return db, nil
}
