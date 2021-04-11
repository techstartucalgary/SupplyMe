package main

import "database/sql"

type Asset struct {
	Org       string
	TID       string
	Content   []byte
	Timestamp int
	Signature string
}

type AssetModel struct {
	DB *sql.DB
}

type Unit struct {
	UID           []byte
	NFT           []byte
	CreationStamp int
	Timestamp     int
	Signature     []byte
}

func (m AssetModel) All() ([]Asset, error) {
	rows, err := DB.Query("SELECT * FROM assets")
	if err != nil {
		return nil, err
	}

	var tks []Asset

	for rows.Next() {
		var a Asset

		err := rows.Scan(&a.Org, &a.TID, &a.Content, &a.Timestamp, &a.Signature)
		if err != nil {
			return nil, err
		}

		tks = append(tks, a)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return tks, nil
}
