package models

import "database/sql"


type Recipe struct {
    ID int
    Title string
    Description sql.NullString
}

