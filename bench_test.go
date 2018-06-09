package main

import (
	"database/sql"
	"sync"
	"testing"

	"github.com/jackc/pgx"

	_ "github.com/jackc/pgx"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/picobank/instruments-sqlboiler/models"
)

var (
	db                *sql.DB
	conn              *pgx.ConnPool
	setupOnce         sync.Once
	randInstrumentIDs []int
)

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

		rows, _ := conn.Query("select instrument_id from instrument ORDER BY random()")
		for rows.Next() {
			var id int
			rows.Scan(&id)
			randInstrumentIDs = append(randInstrumentIDs, id)
		}
	})

}

// BenchmarkFindInstrument_SQLBoiler tests a select of a single row by ID
func BenchmarkFindInstrument_SQLBoiler(b *testing.B) {
	setup(b)
	//boil.DebugMode = true
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		id := randInstrumentIDs[i%len(randInstrumentIDs)]
		instrument, err := models.FindInstrument(db, id)
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
		row := stmt.QueryRow(randInstrumentIDs[j%len(randInstrumentIDs)])
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
	_, err := conn.Prepare("select_instrument", "select * from instrument where instrument_id = $1")
	if err != nil {
		b.Fatal("Failed prepare query", "err", err)
	}
	var i models.Instrument

	for j := 0; j < b.N; j++ {
		row := conn.QueryRow("select_instrument", randInstrumentIDs[j%len(randInstrumentIDs)])
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
