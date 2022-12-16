# go-whosonfirst-spatial-www-pmtiles

Go package implementing the `whosonfirst/go-whosonfirst-spatial-www` server application with support for `whosonfirst/go-whosonfirst-spatial-pmtiles` databases.

## Documentation

Documentation is incomplete at this time.

Please start with the [whosonfirst/go-whosonfirst-spatial-pmtiles](https://github.com/whosonfirst/go-whosonfirst-spatial-pmtiles) documentation.

## Example

```
package main

import (
	"context"
	_ "github.com/whosonfirst/go-whosonfirst-spatial-pmtiles"
	"github.com/whosonfirst/go-whosonfirst-spatial-www/application/server"
	"log"
)

func main() {

	ctx := context.Background()
	logger := log.Default()

	err := server.Run(ctx, logger)

	if err != nil {
		logger.Fatal(err)
	}
}
```

Importantly this is a very thin wrapper around the application code defined in the [whosonfirst/go-whosonfirst-spatial-www](https://github.com/whosonfirst/go-whosonfirst-spatial-www) package to enable support for the [whosonfirst/go-whosonfirst-spatial-pmtiles](https://github.com/whosonfirst/go-whosonfirst-spatial-pmtiles) package.

## Tools

```
$> make cli
go build -mod vendor -o bin/server cmd/server/main.go
```

### server

```
$> > ./bin/server -h
  -authenticator-uri string
    	A valid sfomuseum/go-http-auth URI. (default "null://")
  -cors-allow-credentials
    	Allow HTTP credentials to be included in CORS requests.
  -cors-origin value
    	One or more hosts to allow CORS requests from; may be a comma-separated list.
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
  -enable-www
    	Enable the interactive /debug endpoint to query points and display results.
  -is-wof
    	Input data is WOF-flavoured GeoJSON. (Pass a value of '0' or 'false' if you need to index non-WOF documents. (default true)
  -iterator-uri string
    	A valid whosonfirst/go-whosonfirst-iterate/v2 URI. Supported schemes are: directory://, featurecollection://, file://, filelist://, geojsonl://, null://, repo://. (default "repo://")
  -leaflet-enable-draw
    	Enable the Leaflet.Draw plugin.
  -leaflet-enable-fullscreen
    	Enable the Leaflet.Fullscreen plugin.
  -leaflet-enable-hash
    	Enable the Leaflet.Hash plugin. (default true)
  -leaflet-initial-latitude float
    	The initial latitude for map views to use. (default 37.616906)
  -leaflet-initial-longitude float
    	The initial longitude for map views to use. (default -122.386665)
  -leaflet-initial-zoom int
    	The initial zoom level for map views to use. (default 14)
  -leaflet-max-bounds string
    	An optional comma-separated bounding box ({MINX},{MINY},{MAXX},{MAXY}) to set the boundary for map views.
  -leaflet-tile-url string
    	A valid Leaflet 'tileLayer' layer URL. Only necessary if -map-provider is "leaflet".
  -log-timings
    	Emit timing metrics to the application's logger
  -map-provider string
    	The name of the map provider to use. Valid options are: leaflet, protomaps, tangram
  -nextzen-apikey string
    	A valid Nextzen API key. Only necessary if -map-provider is "tangram".
  -nextzen-style-url string
    	A valid URL for loading a Tangram.js style bundle. Only necessary if -map-provider is "tangram". (default "/tangram/refill-style.zip")
  -nextzen-tile-url string
    	A valid Nextzen tile URL template for loading map tiles. Only necessary if -map-provider is "tangram". (default "https://tile.nextzen.org/tilezen/vector/v1/512/all/{z}/{x}/{y}.mvt")
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
    	A valid whosonfirst/go-reader.Reader URI. Available options are: [fs:// null:// pmtiles:// repo:// sqlite:// stdin://]. If the value is {spatial-database-uri} then the value of the '-spatial-database-uri' implements the reader.Reader interface and will be used.
  -protomaps-bucket-uri string
    	The gocloud.dev/blob.Bucket URI where Protomaps tiles are stored. Only necessary if -map-provider is "protomaps" and -protomaps-serve-tiles is true.
  -protomaps-caches-size int
    	The size of the internal Protomaps cache if serving tiles locally. Only necessary if -map-provider is "protomaps" and -protomaps-serve-tiles is true. (default 64)
  -protomaps-database string
    	The name of the Protomaps database to serve tiles from. Only necessary if -map-provider is "protomaps" and -protomaps-serve-tiles is true.
  -protomaps-serve-tiles
    	A boolean flag signaling whether to serve Protomaps tiles locally. Only necessary if -map-provider is "protomaps".
  -protomaps-tile-url string
    	A valid Protomaps .pmtiles URL for loading map tiles. Only necessary if -map-provider is "protomaps". (default "/tiles/")
  -server-uri string
    	A valid aaronland/go-http-server URI. (default "http://localhost:8080")
  -spatial-database-uri string
    	A valid whosonfirst/go-whosonfirst-spatial/data.SpatialDatabase URI. options are: [pmtiles:// sqlite://]
  -tilezen-enable-tilepack
    	Enable to use of Tilezen MBTiles tilepack for tile-serving. Only necessary if -map-provider is "tangram".
  -tilezen-tilepack-path string
    	The path to the Tilezen MBTiles tilepack to use for serving tiles. Only necessary if -map-provider is "tangram" and -tilezen-enable-tilezen is true.
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

## AWS

_AWS documentation is incomplete._

### Lambda

| Key | Value |
| --- | --- |
| WHOSONFIRST_ENABLE_CORS | true |
| WHOSONFIRST_CORS_ORIGIN | * |
| WHOSONFIRST_LOG_TIMINGS | true |
| WHOSONFIRST_PROPERTIES_READER_URI | {spatial-database-uri} |
| WHOSONFIRST_SERVER_URI | lambda:// |
| WHOSONFIRST_SPATIAL_DATABASE_URI | pmtiles://?tiles=s3blob%3A%2F%2Fexample%3Fregion%3Dus-east-1%26prefix%3Dtiles%2F%26credentials%3Dexample&database=sfomuseum |

## See also

* https://github.com/whosonfirst/go-whosonfirst-spatial-pmtiles
* https://github.com/whosonfirst/go-whosonfirst-spatial-www
* https://github.com/aaronland/go-http-maps