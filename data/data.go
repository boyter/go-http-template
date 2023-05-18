package data

import (
	"database/sql"
)

type DataModel struct {
	DB *sql.DB
}

func (m *DataModel) GetById(id int64) (*Data, error) {
	stmt := `
		SELECT	id,
				name,
            	value
		FROM	data
		WHERE	id = ?
		LIMIT   1
`

	row := m.DB.QueryRow(stmt, id)
	data := &Data{}

	err := row.Scan(
		&data.Id,
		&data.Name,
		&data.Value,
	)

	if err == sql.ErrNoRows {
		return nil, ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	return data, nil
}

func (m *DataModel) GetByName(name string) (*Data, error) {
	stmt := `
		SELECT	id,
				name,
            	value
		FROM	data
		WHERE	name = ?
		LIMIT   1
`

	row := m.DB.QueryRow(stmt, name)
	data := &Data{}

	err := row.Scan(
		&data.Id,
		&data.Name,
		&data.Value,
	)

	if err == sql.ErrNoRows {
		return nil, ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	return data, nil
}

func (m *DataModel) Update(data Data) (*Data, error) {
	stmt := `
		UPDATE data SET
		value = ?
		WHERE id=?
		LIMIT 1
	`

	_, err := m.DB.Exec(stmt,
		data.Value,
		data.Id)

	if err != nil {
		return nil, err
	}

	acc, err := m.GetById(data.Id)
	if err != nil {
		return nil, err
	}

	return acc, nil
}

func (m *DataModel) Insert(data Data) error {
	stmt := `
		insert into data values (null, ?, ?);
	`

	_, err := m.DB.Exec(stmt,
		data.Name,
		data.Value)

	if err != nil {
		return err
	}

	return nil
}
