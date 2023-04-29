package cmd

import (
	"github.com/fatih/color"
)

func example() string {
	var example string
	example += color.BlueString(" get the palm tree of a website:") + "\n"
	example += color.CyanString("  webpalm -u https://google.com -l 1 --live") + "\n\n"

	example += color.BlueString(" dumping emails and comments in 2 level palming : (regexes are separated by comma)") + "\n"
	example += color.CyanString("  webpalm -u https://google.com -l 2 --regexes emails=\"([a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+)\",comments=\"\\<\\!--.*?-->\"") + "\n\n"

	example += color.BlueString(" export the result to file:") + "\n"
	example += color.CyanString("  json: webpalm -u https://google.com -l 2 --regexes  comments=\"\\<\\!--.*?-->\" -o result.json") + "\n"
	example += color.CyanString("  xml : webpalm -u https://google.com -l 2 --regexes  comments=\"\\<\\!--.*?-->\" -o result.xml") + "\n"
	example += color.CyanString("  txt : webpalm -u https://google.com -l 2 --regexes  comments=\"\\<\\!--.*?-->\" -o result.txt") + "\n"
	return example
}

func regexestable() string {
	return color.RedString(`
 Regex Examples:

  +-----------------+------------------------------------------------------+
  |    Name         |                         Pattern                      |
  +-----------------+------------------------------------------------------+
  |    Email        |    ([a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+)  |
  |    Comment      |                \<\!--.*?-->                          |
  |    10 Numbers   |                \b\d{1,10}\b                          |
  |    Token        |          \b[A-Za-z]+\w*(?=\s|$)                      |
  |    password     |                 password.*$                          |
  +-----------------+------------------------------------------------------+
`)
}

func options(url string, level int, liveMode bool, exportFile string, regexMap map[string]string, statusResponses []int) string {
	var options string
	//wrap it into big square
	options += color.RedString("┌")
	options += color.RedString("[")
	options += color.MagentaString(url)
	options += color.RedString("]\n")
	options += color.RedString("│")
	options += color.BlueString("Level: ") + color.CyanString("%d", level) + "\n"
	options += color.RedString("│")
	options += color.BlueString("Live Mode: ") + color.CyanString("%t", liveMode) + "\n"
	options += color.RedString("│")
	options += color.BlueString("Export File: ") + color.CyanString(exportFile) + "\n"
	options += color.RedString("│")
	options += color.BlueString("Regexes: ") + "\n"
	for k, v := range regexMap {
		options += color.RedString("│")
		options += color.CyanString("  %s: %s\n", k, v)
	}
	options += color.RedString("│")
	options += color.BlueString("Excluded Status: ") + color.CyanString("%v", statusResponses) + "\n"
	options += color.RedString("└")
	return options
}

func usage() string {
	return `webpalm`
}

func long() string {
	return color.HiBlueString(`webpalm is a command-line tool that extracts palm tree struct and body 
data pages using regular expressions.`)
}
func banner() string {
	version := color.MagentaString("v0.0.1")
	author := color.MagentaString("github.com/XORbit01")
	discord := color.MagentaString("discord.gg/g9y7D3xCab")

	banner := `

      ////    //////
   ////  /// //    //
  //  ///////////   /
  /  //    ##   /// /  WEBPALM
 // //     ##     / /
  / /     ###    //
    /     ###    /
         ####
         ####
        ####
~~~~~~~~~~~~~~~~~~~~~
`
	//color // with hiGreen and # and hiCyan
	bannerColor := ""
	for _, c := range banner {
		if c == '#' {
			bannerColor += color.YellowString(string(c))
		} else if c == '~' {
			bannerColor += color.YellowString(string(c))
		} else if c == '/' {
			bannerColor += color.GreenString(string(c))
		} else {
			bannerColor += string(c)
		}
	}

	bannerColor += color.HiBlueString("webpalm ") + version + "\n"
	bannerColor += color.HiBlueString("author: ") + author + "\n"
	bannerColor += color.HiBlueString("discord: ") + discord + "\n"
	bannerColor += "\n"
	return bannerColor
}
