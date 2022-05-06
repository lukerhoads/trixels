package main

import (
	"time"
	"errors"
	"gorm.io/driver/mysql"
  	"gorm.io/gorm"
)

type Store struct {
	*gorm.DB
}

func NewStore(dsn string) (*Store, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Store{
		DB: db,
	}, nil
}

func (s *Store) CreateCommit(x uint16, y uint16, color string, address string) error {
	commit := Commit { 
		CreatedAt: time.Now(),
		X: x,
		Y: y,
		Color: color,
		Address: address,
	}

	result := s.DB.Create(&commit);
	if (result.Error != nil) {
		return result.Error
	}

	return nil
}

func (s *Store) GetDayCommits() []Commit {
	var commits []Commit
	year, month, day := time.Now().Date()
    today := time.Date(year, month, day, 0, 0, 0, 0, time.Now().Location())
	s.DB.Where("created_at >", today).Find(&commits)
	return commits
}