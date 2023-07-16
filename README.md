# webpalm

![banner](https://user-images.githubusercontent.com/130087473/235356807-32b80288-7808-4f66-a6f2-fcbe7ab34b72.png)

<hr>
<br></br>
<br></br>

# Take a look 
[![asciicast](https://asciinema.org/a/Ta9V68iidfWD0DSq2J49H6Ipb.svg)](https://asciinema.org/a/Ta9V68iidfWD0DSq2J49H6Ipb)

## What is webpalm?
WebPalm is a command-line tool that enables users to traverse a website and generate a tree of all its webpages and their links. It uses a recursive approach to enter each link found on a webpage and continues to do so until all levels have been explored.
In addition to generating a site map, WebPalm can extract data from the body of each page using regular expressions and save the results in a file. This feature can be useful for web scraping or extracting specific information.

### ⚠️ DISCLAIMER ⚠️:
this tool is intended to be used for legal purposes only,
and you are responsible for your actions.

### Features
- [x] Generate a palm tree struct of web urls
- [x] Dump data from body pages using regular expressions
- [x] live output mode 
- [x] Export the web-tree to json, xml, txt
- [x] Fast and easy to use
- [x] Colorized output and error handling

### Installation
#### From source
```bash
git clone https://github.com/Malwarize/webpalm.git
cd webpalm
go build -o webpalm && ./webpalm
```
#### From binary
you can download the binary from
[Releases](https://github.com/Malwarize/webpalm/releases/latest)
```bash
wget https://github.com/Malwarize/webpalm/releases/download/v0.0.1/webpalm_x.x.x_os_arch.tar.gz
tar -xvf webpalm_x.x.x_os_arch.tar.gz
cd webpalm
./webpalm
```
### if you have go installed
```bash
go install github.com/Malwarize/webpalm/v2@latest
```
### Usage
```bash
webpalm -h
```
```
Flags:
  -x, --exclude-code ints        status codes to exclude / ex : -x 404,500
  -h, --help                     help for webpalm
  -i, --include strings          include only domains / ex : -i google.com,facebook.com
  -l, --level int                level of palming / ex: -l2
      --live                     live output mode (slow but live streaming) use only 1 thread / ex: --live
  -m, --max-concurrency int      max concurrent tasks / ex: -m 10 (default 10)
  -o, --output string            file to export the result (f.json, f.xml, f.txt) / ex: -o result.json
      --regexes stringToString   regexes to match in each page / ex: --regexes comments="\<\!--.*?-->" (default [])
  -u, --url string               target url / ex: -u https://google.com
```
### Examples

#### get the palm tree of a website: 
```bash
webpalm -u https://google.com -l1 --live
```

#### get palm tree of a website and exclude some status codes: 
```bash
webpalm -u https://google.com -l1 -x 404,500 

```
#### get the palm tree of a website and dump data from the body of the pages: 
```bash
webpalm -u https://google.com -l1 --regexes comments="\<\!--.*?-->" -o result.json"
```

this  will dump the comments of each page in the body of the page
```bash
webpalm -u https://google.com -l1 --regexes comments="\<\!--.*?-->",emails="([a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+.[a-zA-Z0-9-.]+)"
```
this will dump the comments and emails of each page in the body of the page

#### get the palm tree of a website and export it to xml,txt: 
```bash
webpalm -u https://google.com -l3 -o result.xml
```
```bash
webpalm -u https://google.com -l2 -o result.txt
```

#### get the palm tree of a website and include only some domains: 
```bash
webpalm -u https://google.com -l2 -i google.com,facebook.com
```
this will crawl only the urls that contains google.com or facebook.com

### threading and concurrency
####  get the palm tree of a website and use only 5 concurrent tasks:
```bash
webpalm -u https://google.com -l2 -m 5
```
📝 **Note**  that the live mode is working with only 1 thread so you can't use it with the live mode


## Regexes Examples
| Regex | Pattern                             |
|-------|-------------------------------------|
|emails | ([a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+) |
|comments | \\<\\!--.*?-->                      |
|tokens | [a-zA-Z0-9]{32}                     |
|password| \bpassword\b.{0,10}                                    |

Don't forget escaping the regexes if needed

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
you can also contact me on discord:`xorbit.`


## Powered By Malwarize
[![image](https://user-images.githubusercontent.com/130087473/232165094-73347c46-71dc-47c0-820a-1eb36657a8c0.png)](https://discord.gg/g9y7D3xCab)



