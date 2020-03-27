# scopeparser
scopeparser takes a Burp Suite Project Configuration file and extractes the inforamtion about the scope into two test files, `scope.txt` and `excluded.txt`.

## Installation
1. Install [Golang](https://golang.org/).
2. Run the following command:
````
go get -u github.com/maxvaer/scopeparser
````
## Usage
```
scopeparser -f ./BurpConfigurationFile.json
```
Stdin is also supported like:
````
cat BurpConfigurationFile.json | scopeparser
````
scopeparser can be easily chained with other tools like: 
````
cat BurpConfigurationFile.json | scopeparser | cat scope.txt | httprobe | tee -a liveDomains.txt
````
Or if you want to find subdomains with [Amass](https://github.com/OWASP/Amass):
````
cat BurpConfigurationFile.json | scopeparser | amass enum --passive -df scope.txt -o subdomains.txt
````
## Flags
`-f` Path to the Burp Suite Configuration file.  
`-p` Append Ports at the end of a domain, e.g. `http://target.site:80`.

## Inspried by
[httprobe](https://github.com/tomnomnom/httprobe) by [TomNomNom](https://github.com/tomnomnom).  
[hakrawler](https://github.com/hakluke/hakrawler) by [Luke Stephens](https://github.com/hakluke).