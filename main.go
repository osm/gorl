package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "gorl",
		Short: "Reolink CLI",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			v := viper.New()
			v.AutomaticEnv()
			v.SetEnvPrefix("GORL")

			cmd.Flags().VisitAll(func(f *pflag.Flag) {
				if strings.Contains(f.Name, "-") {
					envVarSuffix := strings.ToUpper(strings.ReplaceAll(f.Name, "-", "_"))
					v.BindEnv(f.Name, fmt.Sprintf("%s_%s", "GORL", envVarSuffix))
				}

				if !f.Changed && v.IsSet(f.Name) {
					val := v.Get(f.Name)
					cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
				}
			})

			return nil
		},
	}

	dlCmd = &cobra.Command{
		Use:   "dl",
		Short: "Download recording",
		Run: func(cmd *cobra.Command, args []string) {
			address, username, password := getRequiredArgs(cmd)
			out := getOutputWriter(cmd)
			file, _ := cmd.Flags().GetString("file")
			dl(address, username, password, file, out)
		},
	}

	lsCmd = &cobra.Command{
		Use:   "ls",
		Short: "List recordings",
		Run: func(cmd *cobra.Command, args []string) {
			address, username, password := getRequiredArgs(cmd)
			startDate, _ := cmd.Flags().GetString("date")
			endDate, _ := cmd.Flags().GetString("end-date")
			ls(address, username, password, startDate, endDate)
		},
	}

	snapCmd = &cobra.Command{
		Use:   "snap",
		Short: "Snap a picture",
		Run: func(cmd *cobra.Command, args []string) {
			address, username, password := getRequiredArgs(cmd)
			out := getOutputWriter(cmd)
			snap(address, username, password, out)
		},
	}

	streamCmd = &cobra.Command{
		Use:   "stream",
		Short: "Start a camera stream",
		Run: func(cmd *cobra.Command, args []string) {
			address, username, password := getRequiredArgs(cmd)
			out := getOutputWriter(cmd)
			stream(address, username, password, out)
		},
	}
)

func main() {
	rootCmd.PersistentFlags().StringP("address", "a", "", "Address (required)")
	rootCmd.PersistentFlags().StringP("username", "u", "", "Username (required)")
	rootCmd.PersistentFlags().StringP("password", "p", "", "Password (required)")
	rootCmd.MarkPersistentFlagRequired("address")
	rootCmd.MarkPersistentFlagRequired("password")
	rootCmd.MarkPersistentFlagRequired("username")

	dlCmd.Flags().StringP("file", "f", "", "File (required)")
	dlCmd.Flags().StringP("output-file", "o", "-", "Output file")
	dlCmd.MarkFlagRequired("file")
	rootCmd.AddCommand(dlCmd)

	lsCmd.Flags().StringP("date", "d", "", "List recordings from -d <date>")
	lsCmd.Flags().StringP("end-date", "e", "", "Combined with -d to list recordings until -e <date>")
	rootCmd.AddCommand(lsCmd)

	snapCmd.Flags().StringP("output-file", "o", "-", "Output file")
	rootCmd.AddCommand(snapCmd)

	streamCmd.Flags().StringP("output-file", "o", "-", "Output file")
	rootCmd.AddCommand(streamCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func getRequiredArgs(cmd *cobra.Command) (string, string, string) {
	address, _ := cmd.Flags().GetString("address")
	password, _ := cmd.Flags().GetString("password")
	username, _ := cmd.Flags().GetString("username")
	return address, username, password
}

func getOutputWriter(cmd *cobra.Command) io.Writer {
	out := os.Stdout

	outputFile, _ := cmd.Flags().GetString("output-file")
	if outputFile != "-" {
		var err error
		if out, err = os.Create(outputFile); err != nil {
			exitf("%v\n", err)
		}
	}

	return out
}
