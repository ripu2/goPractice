package main /*The purpose to add package in the top is to keep the code organized, just to keep the modules under the required package
But why main ?, it's a naming convention that show that this is the entry point to the app. What happens genrally is when some one is running our app, they genrally do not enter file name so there must be some entry point to initialsie our app right, so all the files undeer main packages are considered as entery points to the moule */

import "fmt" //It's a package, a part of go standard package, use for I/O operatoin pretty similar to scanF and printF in C lanf

func main() { /* when the file gets executed, compiler will search for this function to start executing the file but if we have
	two files with main package, other file won't have main function */
	fmt.Print("Hello world")
}
