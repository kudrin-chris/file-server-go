package store

import (
	"database/sql"
	"fileServer/internal/app/model"
)

type RecorderRepository struct {
	store *Store
}

func (r *RecorderRepository) CreateRecorder(rec *model.Record) (*model.Record, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO record (N, Mqtt, Invid, Msg_id, Text, Context, Class, Level, Area, Block, Type_, Bit_, Invert_bit) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING id",
		rec.N,
		rec.Mqtt,
		rec.Invid,
		rec.Msg_id,
		rec.Text,
		rec.Context,
		rec.Class,
		rec.Level,
		rec.Area,
		rec.Block,
		rec.Type_,
		rec.Bit_,
		rec.Invert_bit,
	).Scan(&rec.Id); err != nil {
		return nil, err
	}
	return rec, nil
}
func (r *RecorderRepository) CreateDevice(dev *model.Device) (*model.Device, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO device (Unig_guid) VALUES ($1) RETURNING id",
		dev.Unig_guid,
	).Scan(&dev.Device_ID); err != nil {
		return nil, err
	}
	return dev, nil
}

func (r *RecorderRepository) CreateFileName(dev *model.File_name) (*model.File_name, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO file_name (Name, Flag_error) VALUES ($1, $2) RETURNING id",
		dev.Name,
		dev.Flag_error,
	).Scan(&dev.File_ID); err != nil {
		return nil, err
	}
	return dev, nil
}

func (r *RecorderRepository) FindByUnitQuid(unig_guid string) (*model.Record, error) {

	u := &model.Record{}
	if err := r.store.db.QueryRow(
		"SELECT recorder.N, recorder.Mqtt, recorder.Invid, recorder.Msq_id, recorder.Text, recorder.Context, recorder.Class, recorder.Level, recorder.Area, recorder.Addr, recorder.Block, recorder.Type_, recorder.Bit_ , recorder.Invert_bit FROM device join recorder on device.device_ID = recorder.device_ID  WHERE device.unig_guid = $1",
		unig_guid,
	).Scan(
		&u.N,
		&u.Mqtt,
		&u.Invid,
		&u.Msg_id,
		&u.Text,
		&u.Context,
		&u.Class,
		&u.Level,
		&u.Area,
		&u.Addr,
		&u.Block,
		&u.Type_,
		&u.Bit_,
		&u.Invert_bit,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return u, nil
}
