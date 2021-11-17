package models

import (
	"context"
	"station-service/db"
)

type Station struct {
	ID        int64   `json:"station_id" db:"station_id"`
	Name      string  `json:"name" db:"name"`
	Latitude  float32 `json:"lat" db:"lat"`
	Longitude float32 `json:"lng" db:"lng"`
	Provider  string  `json:"provider" db:"provider"`
}

type CreateStationParam struct {
	Name      string
	Latitude  float32
	Longitude float32
	Provider  string
}

type UpdateStationParam struct {
	Name      string
	Latitude  float32
	Longitude float32
	Provider  string
}

type ListStationsParam struct {
	Offset int32
	Limit  int32
}

func (s Station) GetByID(ctx context.Context, id int64) (station Station, err error) {
	db := db.GetDB()

	const query = `SELECT * FROM "stations" WHERE "station_id" = $1`
	err = db.GetContext(ctx, &station, query, id)

	return
}

func (s Station) GetAll(ctx context.Context, arg ListStationsParam) (stations []Station, err error) {
	db := db.GetDB()

	const query = `SELECT * FROM "stations" OFFSET $1 LIMIT $2`
	stations = []Station{}
	err = db.SelectContext(ctx, &stations, query, arg.Offset, arg.Limit)

	return
}

func (s Station) Create(ctx context.Context, arg CreateStationParam) (Station, error) {
	db := db.GetDB()

	const query = `
	INSERT INTO "stations"("name", "lat", "lng", "provider") 
	VALUES ($1, $2, $3, $4)
	RETURNING "station_id", "name", "lat", "lng", "provider"
	`
	row := db.QueryRowContext(ctx, query, arg.Name, arg.Latitude, arg.Longitude, arg.Provider)

	var station Station

	err := row.Scan(
		&station.ID,
		&station.Name,
		&station.Latitude,
		&station.Longitude,
		&station.Provider,
	)

	return station, err
}

func (s Station) Update(ctx context.Context, arg UpdateStationParam, id int64) (Station, error) {
	db := db.GetDB()

	const query = `
	UPDATE "stations"
	SET "name" = $2,
		"lat" = $3,
		"lng" = $4,
		"provider" = $5
	WHERE "station_id" = $1
	RETURNING "station_id", "name", "lat", "lng", "provider"
	`
	row := db.QueryRowContext(ctx, query, id, arg.Name, arg.Latitude, arg.Longitude, arg.Provider)

	var station Station
	err := row.Scan(
		&station.ID,
		&station.Name,
		&station.Latitude,
		&station.Longitude,
		&station.Provider,
	)

	return station, err
}

func (s Station) Delete(ctx context.Context, id int64) error {
	db := db.GetDB()

	const query = `
	DELETE FROM stations
	WHERE "station_id" = $1
	`
	_, err := db.ExecContext(ctx, query, id)

	return err
}
