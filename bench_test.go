package main

import (
	"database/sql"
	"sync"
	"testing"

	"github.com/jackc/pgx"

	_ "github.com/jackc/pgx"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/picobank/sqlboiler/models"
)

var db *sql.DB
var conn *pgx.ConnPool
var setupOnce sync.Once

func setup(b *testing.B) {
	setupOnce.Do(func() {
		var err error
		err = initViper()
		if err != nil {
			log.Error("Error while reading configuration")
			return
		}

		connectionString := createConnectionString()
		db, err = sql.Open("pgx", connectionString)
		if err != nil {
			log.Error("Failed to open stdlib database", "err", err)
			return
		}
		var config pgx.ConnPoolConfig
		config.ConnConfig, err = pgx.ParseDSN(connectionString)
		if err != nil {
			log.Error("Failed parse DSN", "err", err)
			return
		}
		config.TLSConfig = nil
		config.UseFallbackTLS = false

		config.MaxConnections = 10
		conn, err = pgx.NewConnPool(config)
		if err != nil {
			log.Error("Failed to open native connection", "err", err)
			return
		}
	})

}

// BenchmarkFindInstrument_SQLBoiler tests a select of a single row by ID
func BenchmarkFindInstrument_SQLBoiler(b *testing.B) {
	setup(b)
	//boil.DebugMode = true
	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		instrument, err := models.FindInstrument(db, 1)
		if err != nil {
			b.Fatal("Failed to select Instrument", "err", err)
		}
		if instrument == nil {
			b.Fatal("Unable to find instrumnet")
		}
	}
}

// BenchmarkFindInstrument_SQLBoiler tests a select of a single row by ID
func BenchmarkFindInstrument_PGXStdlib(b *testing.B) {
	setup(b)

	b.ResetTimer()
	stmt, err := db.Prepare("select * from instrument where instrument_id = $1")
	if err != nil {
		b.Fatal("Failed prepare query", "err", err)
	}
	defer stmt.Close()
	var i models.Instrument

	for j := 0; j < b.N; j++ {
		row := stmt.QueryRow(1)
		err = row.Scan(
			&i.InstrumentID,
			&i.Symbol,
			&i.Name,
			&i.Description,
			&i.InstrumentClassID,
			&i.CurrencyID,
			&i.FromDate,
			&i.ThruDate,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.CreatedBy,
			&i.UpdatedBy,
		)
		if err != nil {
			b.Fatal("Failed to scan instrument", "err", err)
		}
	}
}

// BenchmarkFindInstrument_SQLBoiler tests a select of a single row by ID
func BenchmarkFindInstrument_PGXNative(b *testing.B) {
	setup(b)

	b.ResetTimer()
	_, err := conn.Prepare("select", "select * from instrument where instrument_id = $1")
	if err != nil {
		b.Fatal("Failed prepare query", "err", err)
	}
	var i models.Instrument

	for j := 0; j < b.N; j++ {
		row := conn.QueryRow("select", 1)
		err = row.Scan(
			&i.InstrumentID,
			&i.Symbol,
			&i.Name,
			&i.Description,
			&i.InstrumentClassID,
			&i.CurrencyID,
			&i.FromDate,
			&i.ThruDate,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.CreatedBy,
			&i.UpdatedBy,
		)
		if err != nil {
			b.Fatal("Failed to scan instrument", "err", err)
		}
	}
}
