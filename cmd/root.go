package cmd

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"

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

		prompt := fmt.Sprintf(`以下の $$$ 以降の文章を、できる限り元の文章の意味を保ったまま、自然な関西弁に変換してください。
なお、余計な解説や説明は一切不要です。変換後の文章のみを出力してください。
$$$
%s`, ipt.String())

		err = client.GenerateContentStream(ctx, prompt, func(p genai.Part) error {
			fmt.Print(p)
			return nil
		})
		if err != nil {
			return err
		}

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
