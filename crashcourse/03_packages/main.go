// I SPENT HOURS TRYING TO FIGURE OUT PACKAGES AND HAD I JUST SAT DOWN AND READ THE DOCS / WATCHED THIS VIDEO I WOULD HAVE GOTTEN IT DONE SO FAST
// FROM NOW ON CHECK THE DOCS FIRST AND A CRASH COURSE VIDEO CLOSE AFTER IF ITS A BASIC THING
package main

import (
	"fmt"
	"math"

	"github.com/samlazrak/crashcourse/03_packages/strutil"
	// there seems to be an issue with the package structure because of the other main.go
)

func main() {
	// basic math functions
	fmt.Println(math.Floor(2.7)) // Floor rounds number down
	fmt.Println(math.Ceil(2.7))  // Ceil rounds number up
	fmt.Println(math.Sqrt(16))   // // Square root

	// creating custom packages - create subdirectory & have all .go files within have package dirname
	// you cannot have two different kinds of packages in a dir
	fmt.Println(strutil.Reverse("Hello World!"))
}
