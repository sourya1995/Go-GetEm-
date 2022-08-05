package main
import(
	"io"
	"os"
	"fmt"
	"bufio"
	"flag"
)

func cat(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n')
		if err == io.EOF{
			break
		}
		fmt.Fprintf(os.Stdout, "%s", buf)
	}
	return
}

func main() {
	flag.Parse() //PARSE CLI FLAGS
	if flag.NArg() == 0 { //if there are no flags
		cat(bufio.NewReader(os.Stdin)) //new buffered reader to take input from KB, call cat()
	}
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i)) //open file with name given in the ith argument

		if err != nil{
			fmt.Fprintf(os.Stderr, "%s: error reading from %s: %s\n", os.Args[0], flag.Arg(i), err.Error())
			continue //continue with next flag
		}
		cat(bufio.NewReader(f))
	}
}