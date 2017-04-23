// Package main starts the server from command line.
package main

import (
	"fmt"

	"github.com/alinex/go-webserver/cmd"
	"github.com/fatih/color"
)

// output welcome message
func init() {
	// color definitions
	cyan := color.New(color.FgCyan)
	yellow := color.New(color.FgYellow, color.Bold)
	// create logo
	cyan.Print("                            ")
	yellow.Println(" __   ____     __")
	cyan.Print("             ######  #####  ")
	yellow.Print("|  | |    \\   |  |  ")
	cyan.Println(" ########### #####       #####")
	cyan.Print("            ######## #####  ")
	yellow.Print("|  | |     \\  |  |  ")
	cyan.Println("############  #####     #####")
	cyan.Print("           ######### #####  ")
	yellow.Print("|  | |  |\\  \\ |  |  ")
	cyan.Println("#####          #####   #####")
	cyan.Print("          ########## #####  ")
	yellow.Print("|  | |  | \\  \\|  |  ")
	cyan.Println("#####           ##### #####")
	cyan.Print("         ##### ##### #####  ")
	yellow.Print("|  | |  |_ \\     |  ")
	cyan.Println("############     #########")
	cyan.Print("        #####  ##### #####  ")
	yellow.Print("|  | |    \\ \\    |  ")
	cyan.Println("############     #########")
	cyan.Print("       #####   ##### #####  ")
	yellow.Print("|__| |_____\\ \\___|  ")
	cyan.Println("#####           ##### #####")
	cyan.Println("      #####    ##### #####                      #####          #####   #####")
	cyan.Println("     ##### ######### ########################## ############  #####     #####")
	cyan.Println("    ##### ##########  ########################   ########### #####       #####")
	cyan.Println("    ___________________________________________________________________________")
	fmt.Println()
	yellow.Println("                               W E B - S E R V E R")
	cyan.Println("    ___________________________________________________________________________")
	fmt.Println()
}

func main() {
	cmd.Execute()
}
