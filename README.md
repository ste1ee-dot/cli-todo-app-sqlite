To run use go.
go build
go install

Then you should be able to use "todo" command while in the project directory.

Imports used:
  "encoding/csv"           - for working with csv files
	"fmt"                    - formatted I/O
	"os"                     - for opening files
	"strconv"                - for converting strings and vice versa
	"text/tabwriter"         - for nice tab alined output
	"github.com/spf13/cobra" - makes it easier to create an actual CLI application
