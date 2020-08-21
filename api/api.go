package api

import "github.com/tphume/converse_be/db"

// Our http web service
type Server struct {
	DBClient *db.Client
}
