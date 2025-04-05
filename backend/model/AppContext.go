package model

import "backend/database"

type AppContext struct{
	DB *database.DB
}