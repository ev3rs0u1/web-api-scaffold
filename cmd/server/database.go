package server

import (
	"web-api-scaffold/internal/app/model"
	"web-api-scaffold/internal/pkg/database"
)

func (s *Server) SetDatabase() *Server {
	var err error
	if s.database, err = database.NewConnection(); err != nil {
		s.logger.Fatal().Err(err).Msg("获取数据库连接时出错")
	}

	err = s.database.
		Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&model.User{},
			&model.File{},
			&model.Object{},
			&model.Chunk{},
			&model.Thumbnail{},
		)

	if err != nil {
		s.logger.Fatal().Err(err).Msg("Database auto migrate failed")
	}

	return s
}
