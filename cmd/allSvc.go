/*
Copyright © 2023 fahrizalfarid
*/
package cmd

import (
	"fmt"
	"log"

	"time"

	"github.com/fahrizalfarid/user-service-rpc/conf"
	api "github.com/fahrizalfarid/user-service-rpc/src/api/server"
	user "github.com/fahrizalfarid/user-service-rpc/src/user-service/server"
	validator "github.com/fahrizalfarid/user-service-rpc/src/validator-service/server"
	"github.com/spf13/cobra"
)

// allSvcCmd represents the allSvc command
var allSvcCmd = &cobra.Command{
	Use:   "allSvc",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: runAllSrv,
}

func init() {
	rootCmd.AddCommand(allSvcCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// allSvcCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// allSvcCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
func runAllSrv(cmd *cobra.Command, args []string) {
	err := conf.LoadEnv("./.env")
	if err != nil {
		panic(err)
	}

	time.Sleep(1 * time.Second)
	go func() {
		log.Fatal(user.RunUserSrv().Run())
	}()

	time.Sleep(1 * time.Second)
	go func() {
		log.Fatal(api.RunServer().Start(":8080"))
	}()

	time.Sleep(3 * time.Second)
	go func() {
		log.Fatal(validator.RunUserValidatorSrv().Run())
	}()

	time.Sleep(3 * time.Second)

	fmt.Println("all services called")
	select {}
}
