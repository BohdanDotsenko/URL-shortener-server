# URL-shortener-server


Created URL shortener server
Implemented HTTP server that can generate shortened URLs.
- The requests to shortened URLs should be redirected to their original
URL (status 302) or
return 404 for unknown URLs.
- Simple HTML form served on the index page where users can
input URL and
retrieve the shortened version from server.
- All of the implemented HTTP handlers have unit tests.
- All shortened URLs persisted locally to a file using
SQLite

<code>go run main.go</code>
start use http://localhost:8181/
