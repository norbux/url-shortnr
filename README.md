## POC for a simple URL shortener

### What is this project?

It is the back-end for a simple URL shortener. It doesn't include any type of front-end application. It works with sequential consecutive numbers as the input to generate strings that represent the short URLs. It offers two options to generate such strings:

* Hashing the consecutive sequence number with the XXH3 algorithm (using the [xxh3](https://github.com/zeebo/xxh3) package)
* Simple base62 encoding of the consecutive sequence number

This back-end works through a REST API (using [Fiber](https://gofiber.io/)) that exposes an endpoint to request new short URLs that map to the original long URL and an endpoint that redirects to the original URLs.

The maping of short and long URLs is stored in a MongoDB collection. The database setup and initialization is provided in a [docker-compose file](https://github.com/norbux/url-shortnr/blob/main/data-store/docker-compose.yaml).

### What this project is NOT

This project is a proof of concept and does not contemplate running at scale. It doesn't include user/access management and only focuses in URL shortening mechanics. The data store administration is out of the scope of this project.