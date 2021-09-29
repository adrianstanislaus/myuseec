package main

import (
	_middlewares "myuseek/app/middlewares"
	"myuseek/app/routes"
	_userUsecase "myuseek/business/users"
	_userController "myuseek/controller/users"
	_mysqlDriver "myuseek/drivers/mysql"
	"time"

	_userRepository "myuseek/drivers/databases/users"

	_artistUsecase "myuseek/business/artists"
	_artistController "myuseek/controller/artists"
	_artistRepository "myuseek/drivers/databases/artists"

	_songUsecase "myuseek/business/songs"
	_songController "myuseek/controller/songs"
	_songRepository "myuseek/drivers/databases/songs"

	_playlistUsecase "myuseek/business/playlists"
	_playlistController "myuseek/controller/playlists"
	_playlistRepository "myuseek/drivers/databases/playlists"

	"log"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`app/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func DbMigrate(db *gorm.DB) {
	db.AutoMigrate(&_userRepository.Users{}, &_artistRepository.Artist{}, &_songRepository.Song{}, &_playlistRepository.Playlist{})
}

func main() {
	// init koneksi databse
	configDB := _mysqlDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}

	configJWT := _middlewares.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	Conn := configDB.InitialDB()
	DbMigrate(Conn)

	e := echo.New()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	userRepository := _userRepository.NewMysqlUserRepository(Conn)
	userUseCase := _userUsecase.NewUserUsecase(configJWT, userRepository, timeoutContext)
	userController := _userController.NewUserController(userUseCase)
	artistRepository := _artistRepository.NewMysqlArtistRepository(Conn)
	artistUseCase := _artistUsecase.NewArtistUsecase(artistRepository, timeoutContext)
	artistController := _artistController.NewArtistController(artistUseCase)
	songRepository := _songRepository.NewMysqlSongRepository(Conn)
	songUseCase := _songUsecase.NewSongUsecase(songRepository, timeoutContext)
	songController := _songController.NewSongController(songUseCase)
	playlistRepository := _playlistRepository.NewMysqlPlaylistRepository(Conn)
	playlistUseCase := _playlistUsecase.NewPlaylistUsecase(playlistRepository, timeoutContext)
	playlistController := _playlistController.NewPlaylistController(playlistUseCase)

	routesInit := routes.ControllerList{
		JwtConfig:          configJWT.Init(),
		UserController:     *userController,
		ArtistController:   *artistController,
		SongController:     *songController,
		PlaylistController: *playlistController,
	}

	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))

}
