package cmd

import (
	"fmt"
	"github.com/mosuka/indigo/constant"
	"github.com/mosuka/indigo/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var indigoSettings = viper.New()

//var indigoSettings *viper.Viper

var RootCmd = &cobra.Command{
	Use:   "indigo",
	Short: "Indigo Command Line Interface",
	Long:  `The Indigo Command Line Interface controlls the Indigo Server.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if versionFlag {
			fmt.Printf("%s\n", version.Version)
			os.Exit(0)
		}

		return nil
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return cmd.Help()
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	fmt.Println("root.init()")

	cobra.OnInitialize(initRootCmd)

	//indigoSettings.SetDefault("output_format", constant.DefaultOutputFormat)
	//indigoSettings.SetDefault("log_output", constant.DefaultLogOutputFile)
	//indigoSettings.SetDefault("log_level", constant.DefaultLogLevel)
	//indigoSettings.SetDefault("grpc_port", constant.DefaultGRPCServerPort)
	//indigoSettings.SetDefault("data_dir", constant.DefaultDataDir)
	//indigoSettings.SetDefault("open_existing_index", constant.DefaultOpenExistingIndex)
	//indigoSettings.SetDefault("rest_port", constant.DefaultRESTServerPort)
	//indigoSettings.SetDefault("base_uri", constant.DefaultBaseURI)
	//indigoSettings.SetDefault("grpc_server", constant.DefaultGRPCServer)
	//
	//indigoSettings.SetConfigName("indigo")
	//indigoSettings.SetConfigType("yaml")
	//indigoSettings.AddConfigPath("/etc/indigo/")
	//indigoSettings.AddConfigPath("./etc/")
	//indigoSettings.AddConfigPath("${HOME}/")
	//indigoSettings.AddConfigPath(".")
	//err := indigoSettings.ReadInConfig()
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//
	//indigoSettings.SetEnvPrefix("indigo")
	//indigoSettings.BindEnv("output_format")
	//indigoSettings.BindEnv("log_output")
	//indigoSettings.BindEnv("log_level")
	//indigoSettings.BindEnv("grpc_port")
	//indigoSettings.BindEnv("data_dir")
	//indigoSettings.BindEnv("open_existing_index")
	//indigoSettings.BindEnv("rest_port")
	//indigoSettings.BindEnv("base_uri")
	//indigoSettings.BindEnv("grpc_server")

	RootCmd.PersistentFlags().BoolVarP(&versionFlag, "version", "v", false, "show version numner")

	//RootCmd.PersistentFlags().StringVarP(&outputFormat, "output-format", "f", indigoSettings.GetString("output_format"), "output format")
	RootCmd.PersistentFlags().StringP("output-format", "f", indigoSettings.GetString("output_format"), "output format")
	indigoSettings.BindPFlag("output_format", RootCmd.Flags().Lookup("output-format"))
}

func initRootCmd() {
	fmt.Println("root.initRootCmd()")

	indigoSettings.SetDefault("output_format", constant.DefaultOutputFormat)
	indigoSettings.SetDefault("log_output", constant.DefaultLogOutputFile)
	indigoSettings.SetDefault("log_level", constant.DefaultLogLevel)
	indigoSettings.SetDefault("grpc_port", constant.DefaultGRPCServerPort)
	indigoSettings.SetDefault("data_dir", constant.DefaultDataDir)
	indigoSettings.SetDefault("open_existing_index", constant.DefaultOpenExistingIndex)
	indigoSettings.SetDefault("rest_port", constant.DefaultRESTServerPort)
	indigoSettings.SetDefault("grpc_server", constant.DefaultGRPCServer)
	indigoSettings.SetDefault("base_uri", constant.DefaultBaseURI)

	fmt.Println(indigoSettings.GetString("grpc_server"))

	indigoSettings.SetConfigName("indigo")
	indigoSettings.SetConfigType("yaml")
	indigoSettings.AddConfigPath("/etc/indigo/")
	indigoSettings.AddConfigPath("./etc/")
	indigoSettings.AddConfigPath("${HOME}/")
	indigoSettings.AddConfigPath(".")
	err := indigoSettings.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(indigoSettings.GetString("grpc_server"))

	indigoSettings.SetEnvPrefix("indigo")
	indigoSettings.BindEnv("output_format")
	indigoSettings.BindEnv("log_output")
	indigoSettings.BindEnv("log_level")
	indigoSettings.BindEnv("grpc_port")
	indigoSettings.BindEnv("data_dir")
	indigoSettings.BindEnv("open_existing_index")
	indigoSettings.BindEnv("rest_port")
	indigoSettings.BindEnv("grpc_server")
	indigoSettings.BindEnv("base_uri")

	fmt.Println(indigoSettings.GetString("grpc_server"))
}
