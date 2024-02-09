package cmd

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/google/generative-ai-go/genai"
	"github.com/koki-develop/kansai/internal/config"
	"github.com/koki-develop/kansai/internal/kansai"
	"github.com/koki-develop/kansai/internal/util"
	"github.com/spf13/cobra"
)

var (
	flagConfigure bool   // --configure
	flagAPIKey    string // --api-key, -k
)

var rootCmd = &cobra.Command{
	Use:  "kansai",
	Long: "kansai is a CLI tool for converting text to Kansai dialect.",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		if flagConfigure {
			key := flagAPIKey
			if key == "" {
				key = os.Getenv("KANSAI_API_KEY")
			}
			if key == "" {
				k, err := util.ReadPassword(os.Stdout, "Enter your Gemini API key")
				if err != nil {
					return err
				}
				key = k
			}

			if err := config.SaveAPIKey(key); err != nil {
				return err
			}

			return nil
		}

		apiKey, err := config.LoadAPIKey()
		if err != nil {
			return err
		}

		client, err := kansai.New(ctx, apiKey)
		if err != nil {
			return err
		}
		defer client.Close()

		ipt := new(bytes.Buffer)
		if _, err := io.Copy(ipt, os.Stdin); err != nil {
			return err
		}

		s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
		s.Start()
		err = client.Convert(ctx, ipt.String(), func(p genai.Part) error {
			if s.Active() {
				s.Stop()
			}
			fmt.Print(p)
			return nil
		})
		if err != nil {
			return err
		}

		fmt.Println()
		return nil
	},
}

func init() {
	rootCmd.Flags().BoolVar(&flagConfigure, "configure", false, "configure API key")
	rootCmd.Flags().StringVarP(&flagAPIKey, "api-key", "k", "", "API Key for the Gemini API")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
