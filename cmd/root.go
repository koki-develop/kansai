package cmd

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/koki-develop/kansai/internal/config"
	"github.com/spf13/cobra"
	"golang.org/x/term"
	"google.golang.org/api/option"
)

var (
	flagConfigure bool
	flagAPIKey    string
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
				fmt.Print("Enter your Gemini API key: ")
				k, err := term.ReadPassword(int(os.Stdin.Fd()))
				if err != nil {
					return err
				}
				key = string(k)
				fmt.Println()
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

		client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
		if err != nil {
			panic(err)
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

		model := client.GenerativeModel("gemini-pro")

		resp, err := model.GenerateContent(ctx, genai.Text(prompt))
		if err != nil {
			panic(err)
		}

		b := new(strings.Builder)
		for _, c := range resp.Candidates {
			for _, p := range c.Content.Parts {
				fmt.Fprint(b, p)
			}
		}
		out := strings.TrimSpace(strings.TrimPrefix(b.String(), "$$$"))

		fmt.Println(out)
		return nil
	},
}

func init() {
	rootCmd.Flags().BoolVar(&flagConfigure, "configure", false, "configure API key")
	rootCmd.Flags().StringVarP(&flagAPIKey, "key", "k", "", "API Key for the Gemini API")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
