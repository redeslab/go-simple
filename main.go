package main

import (
	"fmt"
	"github.com/redeslab/go-simple/node"
	"github.com/redeslab/go-simple/pbs"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

var param struct {
	debug    bool
	version  bool
	CMDPort  string
	password string
	path     string
}

var rootCmd = &cobra.Command{
	Use:   "Simple",
	Short: "Simple",
	Long:  `usage description`,
	Run:   mainRun,
}

func init() {

	rootCmd.Flags().BoolVarP(&param.version, "version",
		"v", false, "Simple version")

	rootCmd.Flags().StringVarP(&param.password, "password",
		"p", "", "Password to unlock miner")

	rootCmd.Flags().StringVarP(&param.path, "wallet.path",
		"w", "", "wallet path used in this miner")
	rootCmd.Flags().StringVarP(&param.CMDPort, "cmdPort",
		"c", "42776", "Cmd service port")

	rootCmd.Flags().BoolVarP(&param.debug, "debug", "d", false, "true: ropsten, false: mainnet")

	rootCmd.AddCommand(pbs.InitCmd)
	rootCmd.AddCommand(pbs.ConfCmd)
	rootCmd.AddCommand(pbs.ShowCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func mainRun(_ *cobra.Command, _ []string) {
	if param.version {
		fmt.Println("Simple version: ", node.Version)
		return
	}

	node.InitNodeConfig(param.password, param.path)
	node.Inst().StartUp()
	go pbs.StartCmdService(param.CMDPort)
	done := make(chan bool, 1)
	go waitSignal(done)
	<-done
}

func waitSignal(done chan bool) {
	pid := strconv.Itoa(os.Getpid())
	fmt.Printf("\n>>>>>>>>>>miner start at pid(%s)<<<<<<<<<<\n", pid)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	sig := <-sigCh

	node.Inst().Stop()
	fmt.Printf("\n>>>>>>>>>>process finished(%s)<<<<<<<<<<\n", sig)
	done <- true
}
