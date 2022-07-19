package cmd

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/fiskil/cdr"
	"github.com/fiskil/cdr/storage"
	"github.com/jrapoport/chestnut/encryptor/crypto"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

func init() {
	rootCmd.AddCommand(accessCmd)
	accessCmd.AddCommand(fetchCmd)
	accessCmd.AddCommand(setCmd)
}

var accessCmd = &cobra.Command{
	Use:   "access",
	Short: "Stores and retrieves access tokens for data holders",
}

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetches an access token",
	Run: func(cmd *cobra.Command, args []string) {
		dbName := "default"
		tokenName := "default"

		if len(args) == 2 {
			dbName = args[0]
			tokenName = args[1]
		}

		if len(args) == 1 {
			tokenName = args[0]
		}

		secret := promptSecret("Encryption password for database: " + dbName)

		store, err := storage.New(dbName, crypto.TextSecret(secret))
		if err != nil {
			fmt.Println(err)
			return
		}

		access, err := store.AccessToken(cmd.Context(), tokenName)
		if err != nil {
			resErr, ok := err.(*cdr.ErrNon2xxResponse)
			if ok {
				b, err := ioutil.ReadAll(resErr.Response)
				fmt.Println(string(b), err)
			}
			fmt.Println(err)
			return
		}

		fmt.Println(access)
	},
}

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets a new access token",
	Run: func(cmd *cobra.Command, args []string) {
		dbName := "default"
		tokenName := "default"

		if len(args) == 2 {
			dbName = args[0]
			tokenName = args[1]
		}

		if len(args) == 1 {
			tokenName = args[0]
		}

		secret := promptSecret("Encryption password for database: " + dbName)

		store, err := storage.New(dbName, crypto.TextSecret(secret))
		if err != nil {
			fmt.Println(err)
			return
		}

		refreshToken := promptPlain("Refresh token provided by data holder")
		arrangmentID := promptPlain("CDR arrangement id provided by data holder")
		clientID := promptPlain("Your client id provided by the data holder")
		tokenEndpoint := promptPlain("Data holders token endpoint")

		a := storage.Arrangement{
			ID:               tokenName,
			RefreshToken:     refreshToken,
			CDRArrangementID: arrangmentID,
			ClientID:         clientID,
			TokenEndpoint:    tokenEndpoint,
		}

		if err := store.NewArrangement(a); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("ok")
	},
}

// promptSecret prompts user for an input that is not echo-ed on terminal.
func promptSecret(question string) string {
	fmt.Printf(question + "\n> ")

	raw, err := terminal.MakeRaw(0)
	if err != nil {
		panic(err)
	}
	defer terminal.Restore(0, raw)

	var (
		prompt string
		answer string
	)

	term := terminal.NewTerminal(os.Stdin, prompt)
	for {
		char, err := term.ReadPassword(prompt)
		if err != nil {
			panic(err)
		}
		answer += char

		if char == "" || char == answer {
			return answer
		}
	}
}

// promptPlain prompts user for an input that is echo-ed on terminal.
func promptPlain(question string) string {
	fmt.Printf(question + "\n> ")

	reader := bufio.NewReader(os.Stdin)
	for {
		answer, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		answer = strings.TrimSuffix(answer, "\n")
		return answer
	}
}
