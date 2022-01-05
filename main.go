package main

// else if conditional statements go
// import (
// 	"fmt"
// 	"os"
// )

// const (
// 	usage   = "Usage: <username> <password>"
// 	errUser = "Access Denied for %q.\n"
// 	errPwd  = "Invalid password for %q.\n"
// 	succMsg = "Access granted for %q.\n"
// 	user    = "jack"
// 	pass    = "secret"
// )

// func main() {
// 	args := os.Args
// 	if len(args) != 3 {
// 		fmt.Println(usage)
// 		return
// 	}
// 	u, p := args[1], args[2]
// 	if u != user {
// 		fmt.Printf(errUser, u)
// 		return
// 	} else if p != pass {
// 		fmt.Printf(errPwd, u)
// 		return
// 	} else {
// 		fmt.Printf(succMsg, u)
// 	}

// }

// Advanced sting formatting
// import "fmt"

// func main() {
// 	var (
// 		planet   = "Earth"
// 		distance = 98
// 		orbital  = 87.23
// 		hasLife  = false
// 	)

// 	fmt.Printf("Planet: %v\n", planet)
// 	fmt.Printf("Distance: %v millions km\n", distance)
// 	fmt.Printf("Orbital: %v distance in km\n", orbital)
// 	fmt.Printf("Has %v life: %v\n", planet, hasLife)
// }

// IOTA examples
// import "fmt"

// func main() {
// 	const (
// 		monday = iota
// 		tuesday
// 		wednesday
// 		thursday
// 		friday
// 		saturday
// 		sunday
// 	)

// 	fmt.Println(monday, tuesday, wednesday, thursday, friday, saturday, sunday)

// 	const (
// 		EST = -(5 + iota)
// 		_
// 		MST
// 		PST
// 	)

// 	fmt.Println(EST, MST, PST)
// }

// String manipulation
// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {
// 	msg := os.Args[1]
// 	l := len(msg)

// 	s := msg + strings.Repeat("!", l)
// 	s = strings.ToUpper(s)
// 	fmt.Println(s)
// }

// Find the string length using utf8 package
// import (
// 	"fmt"
// 	"unicode/utf8"
// )

// func main() {
// 	name := "ållo"
// 	fmt.Println(utf8.RuneCountInString(name))
// }

// String literals
// import "fmt"

// func main() {
// 	var s string
// 	s = "How are you?"
// 	s = `how are you?`
// 	fmt.Println(s)
// 	s = "<html>\n\t<body>\"Hello\"</body>\n</html>"
// 	fmt.Println(s)
// 	s = `
// <html>
// 	<body>
// 		"Hello"
// 	</body>
// </html>`
// 	fmt.Println(s)
// 	fmt.Println("c:\\my\\dir\\file")
// 	fmt.Println(`c:\my\dir\file`)
// }

// number of arguments
// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	fmt.Printf("%#v\n", os.Args)
// 	fmt.Println("Path:", os.Args[0])
// 	fmt.Println("1st argument:", os.Args[1])
// 	fmt.Println("2nd argument:", os.Args[2])
// 	fmt.Println("3rd argument:", os.Args[3])

// 	fmt.Println(
// 		"Number of items inside os.args",
// 		len(os.Args),
// 	)
// }

// Type conversion
// func main() {
// 	speed := 100
// 	force := 2.5
// 	speed = int(float64(speed) * force)

// 	fmt.Println(speed)
// }
