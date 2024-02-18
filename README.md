# WebPalm

![banner](https://user-images.githubusercontent.com/130087473/235356807-32b80288-7808-4f66-a6f2-fcbe7ab34b72.png)

<hr>
<br></br>
<br></br>

# Take a look 
![takealook-min](https://github.com/Malwarize/webpalm/assets/130087473/6c601672-f278-431d-854b-0a9876a2fafd)

## What is webpalm?
WebPalm is a command-line tool that enables users to traverse a website and generate a tree of all its webpages and their links. It uses a recursive approach to enter each link found on a webpage and continues to do so until all levels have been explored.
In addition to generating a site map, WebPalm can extract data from the body of each page using regular expressions and save the results in a file. This feature can be useful for web scraping or extracting specific information.

### ⚠️ DISCLAIMER ⚠️:
this tool is intended to be used for legal purposes only,
and you are responsible for your actions.

### Features
- [x] Generate a palm tree struct of web urls
- [x] Dump data from body pages using regular expressions
- [x] Multi-threading and parallelism
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
  -d, --delay int                delay (ms) between each request / ex: -d 200
  -x, --exclude-code ints        status codes to exclude / ex : -x 404,500
  -h, --help                     help for webpalm
  -i, --include strings          include only domains / ex : -i google.com,facebook.com
  -l, --level int                level of palming / ex: -l2
  -o, --output string            file to export the result (f.json, f.xml, f.txt) / ex: -o result.json
  -p, --proxy string             proxy to use / ex: -p http://proxy.com:8080
      --regexes stringToString   regexes to match in each page / ex: --regexes comments="\<\!--.*?-->" (default [])
  -t, --timeout int              timeout in seconds / ex: -t 10 (default 10)
  -u, --url string               target url / ex: -u https://google.com
  -a, --user-agent string        user agent to use / ex: -a chrome, firefox, safari, ie, edge, opera, android, ios, custom
  -v, --version                  version for webpalm
  -w, --worker int               number of workers for multi-threading  / ex: -w 10
```
### Examples

#### get the palm tree of a website: 
```bash
webpalm -u https://google.com -l1
# or
webpalm -u https://google.com -l1 -w 3 # 3 workers (multi-threading)
```

#### get palm tree of a website and exclude some status codes: 
```bash
webpalm -u https://google.com -l1 -x 404,500 

```
#### get the palm tree of a website and dump data from the body of the pages: 
```bash
webpalm -u https://google.com -l1 --regexes comments="\<\!--.*?-->" -o result.json
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
####  get the palm tree of a website using 100 workers:
```bash
webpalm -u https://google.com -l2 -w 100
```


## Regexes Examples
| Regex | Pattern                             |
|-------|-------------------------------------|
|emails | ([a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+) |
|comments | \\<\\!--.*?-->                      |
|tokens | [a-zA-Z0-9]{32}                     |
|password| \bpassword\b.{0,10}                                    |

Don't forget escaping the regexes if needed

## Tests
You can run unit tests to gain more confidence in the enhancements or changes to the code by running `go test -v ./...`

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
you can also contact me on discord:`xorbit.`


## Powered By Malwarize
[Join to Discord](https://discord.gg/ccBJZU99wT)

