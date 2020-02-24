package cmd

import (
	"log"
	"os"

	"github.com/MasatoTokuse/webpush/webpush/dbaccess"
	"github.com/MasatoTokuse/webpush/webpush/message"
	"github.com/MasatoTokuse/webpush/webpush/server"
	"github.com/natefinch/lumberjack"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	port       string
	dbServer   string
	dbPort     string
	dbSchema   string
	dbLogin    string
	dbPassword string
	logpath    string
)

func NewCmdRoot() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "webpush",
		Short: "webpush",
		Run: func(cmd *cobra.Command, args []string) {
			var err error
			_ = err
			conargs := getConnectArgs()
			message := message.NewMessage("title from golang", "bodyyyy from golang")
			err = message.Push(conargs)
		},
	}

	flags := cmd.PersistentFlags()
	flags.StringVar(&port, "port", ":3000", "Listen port")
	flags.StringVar(&dbServer, "db_server", "localhost", "db server")
	flags.StringVar(&dbPort, "db_port", "33306", "db port")
	flags.StringVar(&dbSchema, "db_schema", "webpush", "db schema")
	flags.StringVar(&dbLogin, "db_login", "webpush", "db login")
	flags.StringVar(&dbPassword, "db_password", "webpush", "db password")
	flags.StringVar(&logpath, "log", "-", "log file path")

	viper.SetEnvPrefix("PUSH")
	viper.AutomaticEnv()

	if viper.IsSet("log") {
		flags.Set("log", viper.GetString("log"))
	}
	if logpath != "-" {

		log.SetOutput(&lumberjack.Logger{
			Filename:   logpath,
			MaxSize:    500, // megabytes
			MaxBackups: 10,
			MaxAge:     1,    //days
			Compress:   true, // disabled by default
		})
		log.Println("log:" + logpath)

	} else {
		log.Println("log:(stdout)")
	}

	if viper.IsSet("port") {
		flags.Set("port", viper.GetString("port"))
	}
	log.Println("port:" + port)

	if viper.IsSet("db_server") {
		flags.Set("db_server", viper.GetString("db_server"))
	}
	log.Println("db_server:" + dbServer)

	if viper.IsSet("db_port") {
		flags.Set("db_port", viper.GetString("db_port"))
	}
	log.Println("db_port:" + dbPort)

	if viper.IsSet("db_schema") {
		flags.Set("db_schema", viper.GetString("db_schema"))
	}
	log.Println("db_schema:" + dbSchema)

	if viper.IsSet("db_login") {
		flags.Set("db_login", viper.GetString("db_login"))
	}
	log.Println("db_login:" + dbLogin)

	if viper.IsSet("db_password") {
		flags.Set("db_password", viper.GetString("db_password"))
	}
	log.Println("db_password:" + dbPassword)

	return cmd
}

func Execute() {
	server := server.NewServer()
	cmd := NewCmdRoot()
	cmd.AddCommand(NewCmdAuth(server))

	if err := cmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func getConnectArgs() *dbaccess.ConnectArgs {
	var conarg dbaccess.ConnectArgs

	conarg.Address = dbServer
	conarg.Port = dbPort
	conarg.DBName = dbSchema
	conarg.User = dbLogin
	conarg.Password = dbPassword

	return &conarg
}
