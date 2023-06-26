package cmd

import (
	"github.com/fatih/color"
)

func example() string {
	var example string
	example += color.BlueString(" get the palm tree of a website:") + "\n"
	example += color.CyanString("  webpalm -u https://google.com -l1 --live") + "\n\n"

	example += color.BlueString(" dumping emails from google.com domain pages and comments in 2 level palming : (regexes are separated by comma)") + "\n"
	example += color.CyanString("  webpalm -u https://google.com -l2 -i google.com --regexes emails=\"([a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+)\",comments=\"\\<\\!--.*?-->\"") + "\n\n"

	example += color.BlueString(" export the result to file:") + "\n"
	example += color.CyanString("  json: webpalm -u https://google.com -l2 --regexes  comments=\"\\<\\!--.*?-->\" -o result.json") + "\n"
	example += color.CyanString("  xml : webpalm -u https://google.com -l2 --regexes  comments=\"\\<\\!--.*?-->\" -o result.xml") + "\n"
	example += color.CyanString("  txt : webpalm -u https://google.com -l2 --regexes  comments=\"\\<\\!--.*?-->\" -o result.txt") + "\n"

	return example
}

func regexestable() string {
	table := `
                    |                   Pattern
      --------------+----------------------------------------------------
       Email        |    ([a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+)  
       Comment      |           \<\!--.*?-->                          
       10 Numbers   |            \b\d{1,10}\b                          
       Token        |          	[a-zA-Z0-9]{32}                      
       password     |           \bpassword\b.{0,10}                         
 `
	coloredTable := ""
	for _, c := range table {
		if c == '|' || c == '+' || c == '-' {
			coloredTable += color.YellowString(string(c))
		} else {
			coloredTable += string(c)
		}
	}
	return coloredTable
}

func options(url string, level int, liveMode bool, exportFile string, regexMap map[string]string, statusResponses []int, includedUrls []string, maxConcurrency int) string {
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
	if exportFile == "" {
		options += color.BlueString("Export to: ") + color.CyanString("%s", "nothing") + "\n"
	} else {
		options += color.BlueString("Export to: ") + color.CyanString(exportFile) + "\n"
	}
	options += color.RedString("│")
	options += color.BlueString("Regexes: ") + "\n"
	if len(regexMap) == 0 {
		options += color.RedString("│")
		options += color.CyanString("  %s\n", "nothing")
	}
	for k, v := range regexMap {
		options += color.RedString("│")
		options += color.CyanString("  %s: %s\n", k, v)
	}
	options += color.RedString("│")
	options += color.BlueString("Crawl Only : ") + "\n"
	if len(includedUrls) == 0 {
		options += color.RedString("│")
		options += color.CyanString("  %s\n", "all")
	}
	for _, v := range includedUrls {
		options += color.RedString("│")
		options += color.CyanString("  %s\n", v)
	}
	options += color.RedString("│")
	if len(statusResponses) == 0 {
		options += color.BlueString("Excluded Status: ") + color.CyanString("%s", "nothing") + "\n"
	} else {
		options += color.BlueString("Excluded Status: ") + color.CyanString("%v", statusResponses) + "\n"
	}
	options += color.RedString("│")
	options += color.BlueString("Max Concurrency: ") + color.CyanString("%d", maxConcurrency) + "\n"
	options += color.RedString("└")
	return options
}

func usage() string {
	return `webpalm`
}

func long() string {
	return color.HiBlueString(`webpalm is a command-line tool that generates a
palm tree struct of web urls and dump data from
body pages using regular expressions.`)
}
func banner() string {
	version := color.MagentaString(Version)
	author := color.MagentaString("github.com/Malwarize")
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
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
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

	bannerColor += color.HiCyanString("$ webpalm ") + version + "\n"
	bannerColor += color.HiCyanString("$ author ") + author + "\n"
	bannerColor += color.HiCyanString("$ discord ") + discord + "\n"
	bannerColor += color.YellowString("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	bannerColor += "\n"
	return bannerColor
}
