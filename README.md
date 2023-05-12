# webpalm

![banner](https://user-images.githubusercontent.com/130087473/235356807-32b80288-7808-4f66-a6f2-fcbe7ab34b72.png)

<hr>
<br></br>
<br></br>

# Take a look 
[![asciicast](https://asciinema.org/a/Ta9V68iidfWD0DSq2J49H6Ipb.svg)](https://asciinema.org/a/Ta9V68iidfWD0DSq2J49H6Ipb)

## What is webpalm?
webpalm is a command-line tool that traverses a website and generates a tree of all the webpages and their links, additionally it can dump data from the body of the pages using regular expressions then store the result in a file.

### ⚠️ DISCLAIMER ⚠️:
this tool is intended to be used for legal purposes only,
and you are responsible for your actions.

### Features
- [x] Generate a palm tree struct of web urls
- [x] Dump data from body pages using regular expressions
- [x] live output mode 
- [x] Export the webtree to json, xml, txt
- [x] Fast and easy to use
- [x] Colorized output and error handling

### When to use webpalm?
web palm is specially used in OSINT level.
when you want to get a quick overview of a website structure
or when you want to check if there is any sensitive data using regex
it is good at spidering in websites networks and go in depth

### Installation
#### From source
```bash
git clone https://github.com/XORbit01/webpalm.git
cd webpalm
go build -o webpalm && ./webpalm
```
#### From binary
you can download the binary from
[Releases](https://github.com/XORbit01/webpalm/releases/latest)
```bash
wget https://github.com/XORbit01/webpalm/releases/download/v0.0.1/webpalm_x.x.x_os_arch.tar.gz
tar -xvf webpalm_x.x.x_os_arch.tar.gz
cd webpalm
./webpalm
```
### if you have go installed
```bash
go install github.com/XORbit01/webpalm@latest
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
  -l, --level int                level of palming / ex: -l 2
      --live                     live output mode (slow but live streaming) / ex: --live
  -o, --output string            file to export the result (f.json, f.xml, f.txt) / ex: -o result.json
      --regexes stringToString   regexes to match in each page / ex: --regexes comments="\<\!--.*?-->  (default [])
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
you can also contact me on discord:`XORbit#5945`


## Powered By Malwarize
[![image](https://user-images.githubusercontent.com/130087473/232165094-73347c46-71dc-47c0-820a-1eb36657a8c0.png)](https://discord.gg/g9y7D3xCab)



