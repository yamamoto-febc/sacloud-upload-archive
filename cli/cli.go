package cli

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/webguerilla/ftps"
	API "github.com/yamamoto-febc/libsacloud/api"
	"io"
	"io/ioutil"
	"os"
)

// Exit codes are int values that represent an exit code for a particular error.
const (
	ExitCodeOK    int = 0
	ExitCodeError int = 1 + iota
)

// CLI is the command line object
type CLI struct {
	// outStream and errStream are the stdout and stderr
	// to write message from the CLI.
	OutStream, ErrStream io.Writer
}

type context struct {
	token     string
	secret    string
	zone      string
	inputFile *os.File
	imageName string
}

// Run invokes the CLI with the given arguments.
func (cli *CLI) Run(args []string) int {
	var (
		token  string
		secret string
		zone   string
		file   string
	)

	// Define option flag parse
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.ErrStream)

	flags.StringVar(&token, "token", "", "AccessToken / $SAKURACLOUD_ACCESS_TOKEN")
	flags.StringVar(&token, "t", "", "AccessToken(Short) / $SAKURACLOUD_ACCESS_TOKEN")

	flags.StringVar(&secret, "secret", "", "AccessTokenSecret / $SAKURACLOUD_ACCESS_TOKEN_SECRET")
	flags.StringVar(&secret, "s", "", "AccessTokenSecret(Short) / $SAKURACLOUD_ACCESS_TOKEN_SECRET")

	flags.StringVar(&zone, "zone", "is1a", "Target zone")
	flags.StringVar(&zone, "z", "is1a", "Target zone(Short)")

	flags.StringVar(&file, "file", "", "ISO image file")
	flags.StringVar(&file, "f", "", "ISO image file(Short)")

	token = os.Getenv("SAKURACLOUD_ACCESS_TOKEN")
	secret = os.Getenv("SAKURACLOUD_ACCESS_TOKEN_SECRET")

	// Parse commandline flag
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	//validate options
	if token == "" {
		fmt.Fprintf(cli.ErrStream, "%s is required\n", "token")
		return ExitCodeError
	}

	if secret == "" {
		fmt.Fprintf(cli.ErrStream, "%s is required\n", "secret")
		return ExitCodeError
	}

	if zone == "" {
		fmt.Fprintf(cli.ErrStream, "%s is required\n", "zone")
		return ExitCodeError
	}

	//validate args
	if flags.NArg() != 1 {
		fmt.Fprintf(cli.ErrStream, "missing args.please set [ImageName]\n")
		return ExitCodeError
	}

	imageName := flags.Arg(0)

	// exists file or STDIN?
	var inputFile *os.File
	if file == "" {

		fi, err := os.Stdin.Stat()
		if err != nil {
			panic(err)
		}
		if fi.Size() == 0 && fi.Mode()&os.ModeNamedPipe == 0 {
			fmt.Fprintf(cli.ErrStream, "missing file. please use redirect or pipe or '-file' option\n")
			return ExitCodeError
		}

		inputFile = os.Stdin

	} else {
		_, err := os.Stat(file)
		if err != nil {
			fmt.Fprintf(cli.ErrStream, "file[%s] is not found \n", file)
			return ExitCodeError
		}
		inputFile, _ = os.Open(file)
	}
	defer inputFile.Close()

	return cli.UploadImageToSakura(&context{
		token:     token,
		secret:    secret,
		zone:      zone,
		inputFile: inputFile,
		imageName: imageName,
	})
}

func (cli *CLI) UploadImageToSakura(params *context) int {

	// APIクライアント作成
	api := API.NewClient(params.token, params.secret, params.zone)

	//ISO格納領域作成
	newImage := api.Archive.New()
	newImage.Name = params.imageName
	newImage.SizeMB = 20480
	//Create and OpenFTP
	image, err := api.Archive.Create(newImage)

	if err != nil {
		fmt.Fprintf(cli.ErrStream, "Create Archive failed. %#v\n", err)
		return ExitCodeError
	}

	ftp, err := api.Archive.OpenFTP(image.ID, true)

	if err != nil {
		fmt.Fprintf(cli.ErrStream, "Create Archive failed. %#v\n", err)
		return ExitCodeError
	}

	//upload
	ftpsClient := &ftps.FTPS{}
	ftpsClient.TLSConfig.InsecureSkipVerify = true

	err = ftpsClient.Connect(ftp.HostName, 21)
	if err != nil {
		fmt.Fprintf(cli.ErrStream, "Connect FTP failed. %#v\n", err)
		return ExitCodeError
	}

	err = ftpsClient.Login(ftp.User, ftp.Password)
	if err != nil {
		fmt.Fprintf(cli.ErrStream, "Auth FTP failed. %#v\n", err)
		return ExitCodeError
	}

	reader := bufio.NewReader(params.inputFile)
	fileBytes, _ := ioutil.ReadAll(reader)

	fmt.Fprintln(cli.OutStream, "FTP information:")
	fmt.Fprintln(cli.OutStream, "  user: "+ftp.User)
	fmt.Fprintln(cli.OutStream, "  pass: "+ftp.Password)
	fmt.Fprintln(cli.OutStream, "  host: "+ftp.HostName)

	fmt.Fprintln(cli.OutStream, "uploading...")
	err = ftpsClient.StoreFile("sacloud_upload_archive.data", fileBytes)
	if err != nil {
		fmt.Fprintf(cli.ErrStream, "Storefile FTP failed. %#v\n", err)
		return ExitCodeError
	}

	fmt.Fprintln(cli.OutStream, "done.")
	ftpsClient.Quit()

	// close image FTP after upload
	_, err = api.Archive.CloseFTP(image.ID)
	if err != nil {
		fmt.Fprintf(cli.ErrStream, "Close FTP failed. %#v\n", err)
		return ExitCodeError
	}

	return ExitCodeOK

}
