package cmd

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/spf13/cobra"
	"google.golang.org/api/option"
)

var rootCmd = &cobra.Command{
	Use: "kansai",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		// TODO: from config file or flag or env
		apiKey := os.Getenv("GEMINI_API_KEY")
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
		out := strings.TrimSpace(strings.TrimPrefix(b.String(), "$$---$$"))

		fmt.Println(out)
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
