package configs

import "fmt"

const banner = `
   _   _          _   _             
  | \ | |   ___  (_) | |__     ___  
  |  \| |  / _ \ | | | '_ \   / _ \ 
  | |\  | |  __/ | | | |_) | | (_) |
  |_| \_|  \___| |_| |_.__/   \__
`

const version = "v 0.1"

func Banner() {
	fmt.Println(banner)
	fmt.Println("\t\t Author: Swifto")
	fmt.Println("\t\t https://github.com/Jerry-Swift\n")
}
