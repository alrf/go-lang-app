Simple Go application. When started (serv.go), it
will listen on port 5555 (but this may be configurable through a
command-line flag). Clients will be able to connect to this port and
send arbitrary natural language over the wire. The purpose of the
application is to process the text, and store some stats about the
different words that it sees.

The application (web.go) will also expose an HTTP interface on port 8080
(configurable): clients hitting the /stats endpoint will receive a JSON
representation of the statistics about the words that the application has
seen so far.

Specifically, the JSON response should look like:

```javascript
{
  "count": 42,
  "top_5_words": ["lorem", "ipsum", "dolor", "sit", "amet"],
  "top_5_letters": ["e", "a", "o", "i", "p"]
}
```

Where `count` represents the total number of words seen, `top_5_words`
contains the 5 words that have been seen with the highest frequency, and
`top_5_letter` contains the 5 letters that have been seen with the
highest frequency.
