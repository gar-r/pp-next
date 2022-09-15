# pp-next

## about

`pp-next` is a simple planning poker website. An example deployment is available [here](https://ppnext.okki.hu), but it is very easy (and recommended) to deploy it yourself.

## showcase

Join a poker room with the handle of your choice:

![join a poker room](docs/showcase01.png)

Voting while results are hidden:

![voting](docs/showcase02.png)

Revealing the results:

![results](docs/showcase03.png)

## features

   * simple and minimalist, no registration needed
   * does not track users
   * file based persistence
   * easy to host

## getting started

   1. clone the source code:

   ```
   git clone https://github.com/garricasaurus/pp-next
   ```


   2. use [docker-compose](https://docs.docker.com/compose/) to bring up the server and the database:

   ```
   cd pp-next
   docker-compose up -d
   ```

   3. verify that the application is up:

   ```
   http://localhost:38080/
   ```

The database containing the votes will be located in the `mongo/data` directory by default.


## deployment

You can configure the application with environment variables in `docker-compose.yml`:

For most standard deployments it is sufficient to add the following variables:

```yml
services:
  ppnext:
    environment:
      TLS_ENABLED: "true"
      PUBLIC_PORT: 443
      DOMAIN: "your-domain.example.com"
      SUPPORT_EMAIL: "your-email@example.com"
```

The full list of environment variables is below:

| Environment Variable       | Default Value      | Comment                                                                                              |
| -------------------------- | ------------------ | ---------------------------------------------------------------------------------------------------- |
| PORT                       | 38080              | local port - you may also want to change the port mapping in `docker-compose.yml` if you change this |
| PUBLIC_PORT                | $PORT              | public port - this is used for share-url generation                                                  |
| DOMAIN                     | "localhost"        | domain - this is used for share-url generation                                                       |
| SUPPORT_EMAIL              | "email@example.com | support links on the website will show this email                                                    |
| CLEANUP_FREQUENCY_MINUTES  | 10                 | periodic cleanup of inactive rooms occurs this often                                                 |
| CLEANUP_MAX_ROOM_AGE_HOURS | 12                 | the age at which a room is considered inactive                                                       |
| AUTH_COOKIE_NAME           | "ppnext-user"      | the auth cookie name                                                                                 |
| AUTH_COOKIE_EXPIRY_HOURS   | 6                  | the auth cookie expiry                                                                               |

### reverse proxy

It is recommended to use nginx as a reverse proxy (or a similar alternative) in a real deployment scenario. The following example nginx config redirects http and https requests to ppnext (and upgrades http requests to https):

```
server {
    server_name your-domain.example.com;

    location / {
        proxy_pass http://localhost:38080;
    }

    listen [::]:443 ssl;
    listen 443 ssl;
    ssl_certificate path/to/fullchain.pem;
    ssl_certificate_key path/to/privkey.pem;
}
server {
    if ($host = your-domain.example.com) {
        return 301 https://$host$request_uri;
    }

    server_name your-domain.example.com;

    listen 80;
    listen [::]:80;
    return 404;
}
```