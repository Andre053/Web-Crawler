# Web Crawler
- Simple web crawler written in Go
- Given a root URL, it parses the page for web links and recursively checks each site, up to a max depth
- Returns a list of sites found and their frequency

# How to Use
1. With Go installed, run:  go build main.go
2. Start the program with:  ./main
3. Enter a URL to search, omitting the protocol, HTTP is assumed
    - Ex. google.com, twitter.com, etc.
4. Enter a depth to search to

# Notes
- Final output sorts URLs based on frequency

# Extensions TODO
- Show progress of URLs found 
- Utilize concurrency
- Get cookies
- Enumerate all network calls made by a website