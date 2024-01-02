package repository

import (
	"context"
	"database/sql"
	"log"
	"sensors/domain/sensors"
	"sensors/domain/sensors/model"

	"github.com/jmoiron/sqlx"
)

type Repo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) sensors.RepoInterface {
	return &Repo{
		db: db,
	}
}

func (r *Repo) Insert(payload model.SensorData) error {
	var (
		args []interface{}
	)

	args = append(args, payload.ID1, payload.ID2, payload.Type, payload.Value, payload.Timestamp)

	query := `INSERT INTO sensors (id_1, id_2, type, value, timestamp, status) VALUES (?,?,?,?,?,1)`

	_, err := r.db.Exec(query, args...)
	if err != nil {
		log.Println(123, err)
		return err
	}

	return nil
}

func (r *Repo) Update(payload model.SensorData, filter model.FilterData) error {
	var (
		args []interface{}
	)

	query := `
		UPDATE sensors
		SET
			id_1 = ?,
			id_2 = ?,
			type = ?,
			value = ?,
			timestamp = ?
		WHERE status = 1
	`
	args = append(args, payload.ID1, payload.ID2, payload.Type, payload.Value, payload.Timestamp)

	if filter.From != nil && filter.To != nil {
		query += ` AND timestamp >= ? AND timestamp <= ?`
		args = append(args, filter.From, filter.To)
	}

	if filter.ID1 != nil && filter.ID2 != nil {
		query += ` AND id_1 = ? AND id_2 = ?`
		args = append(args, filter.ID1, filter.ID2)
	}

	p, err := r.db.Exec(query, args...)
	if err != nil {
		log.Println(err)
		return err
	}

	num, err := p.RowsAffected()
	if err != nil {
		log.Println(err)
		return err
	}

	if num == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *Repo) Delete(filter model.FilterData) error {
	var (
		args []interface{}
	)

	query := `UPDATE sensors SET status = 0 WHERE status = 1`

	if filter.From != nil && filter.To != nil {
		query += ` AND timestamp >= ? AND timestamp <= ?`
		args = append(args, filter.From, filter.To)
	}

	if filter.ID1 != nil && filter.ID2 != nil {
		query += ` AND id_1 = ? AND id_2 = ?`
		args = append(args, filter.ID1, filter.ID2)
	}

	p, err := r.db.Exec(query, args...)
	if err != nil {
		log.Println(err)
		return err
	}

	num, err := p.RowsAffected()
	if err != nil {
		log.Println(err)
		return err
	}

	if num == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *Repo) GetAll(params *model.FilterData) (sensors []model.SensorData, err error) {
	var (
		// datas []model.SensorData
		args []interface{}
	)
	query := `SELECT id_1, id_2, type, value, timestamp FROM sensors WHERE status = 1`

	if params.From != nil && params.To != nil {
		query += ` AND timestamp >= ? AND timestamp <= ?`
		args = append(args, params.From, params.To)
	}

	if params.ID1 != nil && params.ID2 != nil {
		query += ` AND id_1 = ? AND id_2 = ?`
		args = append(args, params.ID1, params.ID2)
	}

	rows, err := r.db.QueryContext(context.Background(), query, args...)
	if err != nil {
		log.Println(err.Error())
		return sensors, err
	}

	for rows.Next() {
		var data model.SensorData

		err = rows.Scan(
			&data.ID1, &data.ID2, &data.Type, &data.Value, &data.Timestamp,
		)
		if err != nil {
			log.Println(err.Error())
			return sensors, err
		}

		sensors = append(sensors, data)
	}

	if len(sensors) == 0 {
		return sensors, sql.ErrNoRows
	}

	return sensors, nil
}
