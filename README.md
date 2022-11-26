# go-whosonfirst-spatial-www-pmtiles

Go package implementing the `whosonfirst/go-whosonfirst-spatial-www` server application with support for `whosonfirst/go-whosonfirst-spatial-pmtiles` databases.

## Documentation

Documentation is incomplete at this time.

Please start with the [whosonfirst/go-whosonfirst-spatial-pmtiles](https://github.com/whosonfirst/go-whosonfirst-spatial-pmtiles) documentation.

## Tools

```
$> make cli
go build -mod vendor -o bin/server cmd/server/main.go
```

### server

```
$> ./bin/server  -h
  -authenticator-uri string
    	A valid sfomuseum/go-http-auth URI. (default "null://")
  -cors-allow-credentials
    	...
  -cors-origin value
    	...
  -custom-placetypes string
    	A JSON-encoded string containing custom placetypes defined using the syntax described in the whosonfirst/go-whosonfirst-placetypes repository.
  -enable-cors
    	Enable CORS headers for data-related and API handlers.
  -enable-custom-placetypes
    	Enable wof:placetype values that are not explicitly defined in the whosonfirst/go-whosonfirst-placetypes repository.
  -enable-geojson
    	Enable GeoJSON output for point-in-polygon API calls.
  -enable-gzip
    	Enable gzip-encoding for data-related and API handlers.
  -enable-tangram
    	Use Tangram.js for rendering map tiles
  -enable-www
    	Enable the interactive /debug endpoint to query points and display results.
  -is-wof
    	Input data is WOF-flavoured GeoJSON. (Pass a value of '0' or 'false' if you need to index non-WOF documents. (default true)
  -iterator-uri string
    	A valid whosonfirst/go-whosonfirst-iterate/v2 URI. Supported schemes are: directory://, featurecollection://, file://, filelist://, geojsonl://, null://, repo://. (default "repo://")
  -leaflet-initial-latitude float
    	The initial latitude for map views to use. (default 37.616906)
  -leaflet-initial-longitude float
    	The initial longitude for map views to use. (default -122.386665)
  -leaflet-initial-zoom int
    	The initial zoom level for map views to use. (default 14)
  -leaflet-max-bounds string
    	An optional comma-separated bounding box ({MINX},{MINY},{MAXX},{MAXY}) to set the boundary for map views.
  -leaflet-tile-url string
    	A valid Leaflet (slippy map) tile template URL to use for rendering maps (if -enable-tangram is false)
  -log-timings
    	Emit timing metrics to the application's logger
  -nextzen-apikey string
    	A valid Nextzen API key
  -nextzen-style-url string
    	The URL for the style bundle file to use for maps rendered with Tangram.js (default "/tangram/refill-style.zip")
  -nextzen-tile-url string
    	The URL for Nextzen tiles to use for maps rendered with Tangram.js (default "https://tile.nextzen.org/tilezen/vector/v1/512/all/{z}/{x}/{y}.mvt")
  -path-api string
    	The root URL for all API handlers (default "/api")
  -path-data string
    	The URL for data (GeoJSON) handler (default "/data")
  -path-ping string
    	The URL for the ping (health check) handler (default "/health/ping")
  -path-pip string
    	The URL for the point in polygon web handler (default "/point-in-polygon")
  -path-prefix string
    	Prepend this prefix to all assets (but not HTTP handlers). This is mostly for API Gateway integrations.
  -properties-reader-uri string
    	A valid whosonfirst/go-reader.Reader URI. Available options are: [fs:// null:// repo:// sqlite:// stdin://]
  -server-uri string
    	A valid aaronland/go-http-server URI. (default "http://localhost:8080")
  -spatial-database-uri string
    	A valid whosonfirst/go-whosonfirst-spatial/data.SpatialDatabase URI. options are: [pmtiles:// sqlite://]
  -verbose
    	Be chatty.
```

#### Example

```
$> ./bin/server \
	-spatial-database-uri 'pmtiles://?tiles=file:///usr/local/data&database=wof'

2022/11/26 09:57:03 Register /data/ handler
2022/11/26 09:57:03 Register /api/point-in-polygon handler
2022/11/26 09:57:03 Listening on http://localhost:8080
2022/11/26 09:57:03 time to index paths (0) 83.545µs
2022/11/26 09:57:03 finished indexing in 618.003µs
```

And then:

```
$> curl -s http://localhost:8080/api/point-in-polygon \
	-d '{"latitude": 37.621131, "longitude": -122.384292}' \

| jq '.places[]["wof:name"]'

"Earth"
"United States"
"San Francisco International Airport"
"California"
"North America"
"San Mateo"
"94128"
```

## See also

* https://github.com/whosonfirst/go-whosonfirst-spatial-pmtiles
* https://github.com/whosonfirst/go-whosonfirst-spatial-www