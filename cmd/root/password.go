package root

import (
	"fmt"

	pwd "github.com/qinya0/go-tools/pkg/pwd"
	"github.com/qinya0/go-tools/utils/password"
	"github.com/spf13/cobra"
)

var (
	defaultFile = ".go-tools.pwd"

	// pwd flag
	fPwd       string
	fFile      string

	// generate flag
	gUpperStr bool
	gInt      bool
	gSpecial  bool
	gPointer  bool
	gNum      int

	// show flag
	swName  string

	// show all flag: nil

	// save flag
	sName string
	sPwd  string
	sMsg  string
)

var (
	pwdCmd = &cobra.Command{
		Use: "pwd",
		Short: "a tools for password",
	}
	pwdGenerateCmd = &cobra.Command{
		Use: "generate",
		Short: "generate a pwd",
		Run: func(cmd *cobra.Command, args []string) {
			pwd := password.GeneratePassword(gUpperStr, gInt, gSpecial, gPointer, gNum)
			if pwd == "" {
				fmt.Println("[pwd] unknown error")
				return
			}
			fmt.Println("[pwd] pwd:", pwd)
		},
	}
	pwdShowCmd = &cobra.Command{
		Use: "show",
		Short: "show a pwd by name",
		Run: showFunc,
	}
	pwdShowAllCmd = &cobra.Command{
		Use: "showall",
		Short: "show all pwds",
		Run: showAllFunc,
	}
	pwdSaveCmd = &cobra.Command{
		Use: "save",
		Short: "save a pwd",
		Run: saveFunc,
	}
)

var (
	pwdHelp = "[pwd] must set pwd flag!\nsuch as:\n\tgo-tools pwd --admin xxx [show|showall|save] [flag]"
)

func init() {
	pwdCmd.AddCommand(pwdGenerateCmd)
	pwdCmd.AddCommand(pwdShowCmd)
	pwdCmd.AddCommand(pwdShowAllCmd)
	pwdCmd.AddCommand(pwdSaveCmd)

	// InitFlag
	InitFlag()
}

func InitFlag() {
	// pwd
	pwdCmd.PersistentFlags().StringVar(&fPwd, "admin", "", "password for manager")
	pwdCmd.PersistentFlags().StringVar(&fFile, "f", "", "file for save pwd")

	// pwdGenerateCmd
	pwdGenerateCmd.PersistentFlags().BoolVar(&gInt, "i", true, "generate password with int")
	pwdGenerateCmd.PersistentFlags().BoolVar(&gUpperStr, "u", false, "generate password with upper string")
	pwdGenerateCmd.PersistentFlags().BoolVar(&gSpecial, "s", false, "generate password with special character")
	pwdGenerateCmd.PersistentFlags().BoolVar(&gPointer, "p", false, "generate password with pointer(_-:;,.) character")
	pwdGenerateCmd.PersistentFlags().IntVar(&gNum, "n", 10, "the length of generate password")

	// show
	pwdShowCmd.PersistentFlags().StringVar(&swName, "n", "", "name for show")

	// save
	pwdSaveCmd.PersistentFlags().StringVar(&sName, "n", "", "name for save")
	pwdSaveCmd.PersistentFlags().StringVar(&sPwd, "p", "", "password for save")
	pwdSaveCmd.PersistentFlags().StringVar(&sMsg, "m", "", "msg for save")
}

func showAllFunc(cmd *cobra.Command, args []string) {
	// get all pwd from file
	if fPwd == "" {
		fmt.Println(pwdHelp)
		return
	}
	if fFile == "" {
		fFile = defaultFile
	}

	pwds, err := pwd.NewManager(fPwd, fFile).GetAll()
	if err != nil {
		fmt.Printf("[pwd] GetAll err:%s ", err.Error())
		return
	}
	for _, p := range pwds {
		fmt.Printf("[pwd] %s\n", p.String())
	}
}
func showFunc(cmd *cobra.Command, args []string) {
	// get all pwd from file
	if fPwd == "" {
		fmt.Println(pwdHelp)
		return
	}
	if fFile == "" {
		fFile = defaultFile
	}
	if swName == "" {
		fmt.Printf("[pwd] param can't be empty")
		return
	}
	pwds, err := pwd.NewManager(fPwd, fFile).GetOne(swName)
	if err != nil {
		fmt.Printf("[pwd] SaveOne err:%s ", err.Error())
		return
	}
	for _, p := range pwds {
		fmt.Printf("[pwd] %s\n", p.String())
	}

}

func saveFunc(cmd *cobra.Command, args []string) {
	// get all pwd from file
	if fPwd == "" {
		fmt.Println(pwdHelp)
		return
	}
	if fFile == "" {
		fFile = defaultFile
	}
	// get all data from file
	if sName == "" || sPwd == "" {
		fmt.Printf("[pwd] param can't be empty")
		return
	}
	err := pwd.NewManager(fPwd, fFile).SaveOne(sName, sPwd, sMsg)
	if err != nil {
		fmt.Printf("[pwd] SaveOne err:%s ", err.Error())
		return
	}
	fmt.Printf("[pwd] SaveOne Pwd(%s) Successfully\n", sName)
}