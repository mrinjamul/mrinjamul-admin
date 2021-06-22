# Golang Adminer

Admin Panel for mrinjamul.github.io

## Used Tools and Technologies

- [Golang](https://golang.org/)
- [NodeJs](https://nodejs.org/) (Required: 12+ ,Recommanded: 14 LTS)
- [Docker](https://www.docker.com/) (Optional)

- [Auth0 Account](https://auth0.com/)
- [Firebase](https://firebase.google.com/)
- [Gin Gonic](https://gin-gonic.com/)
- [JWT](https://github.com/square/go-jose)
- [xid](https://github.com/rs/xid)
- [Reactstrap](https://reactstrap.github.io/)
- [React Font Awesome](https://github.com/FortAwesome/react-fontawesome)
- [highlight.js](https://highlightjs.org/) (Imported buthNot Used)

## Required secrets

For ENV,

    # Golang ENV (.env)
    GIN_MODE=release
    PORT=<your port>
    AUTH0_API_IDENTIFIER=<auth0 api identifier>
    AUTH0_DOMAIN=<auth0 api domain>

    # ReactJs ENV (ui/.env)
    REACT_APP_API_URL=<your API server url>
    REACT_APP_AUTH0_DOMAIN=<auth0 api domain>
    REACT_APP_AUTH0_CLIENTID=<auth0 Application Client ID>
    REACT_APP_AUTH0_AUDIENCE=<auth0 API Audience>

And Firebase admin sdk private key,

- `serviceAccountKey.json`

Put `serviceAccountKey.json` in project's root directory.

Note: auth_config.json does not required. It's removed.

## Build

### Docker

```shell
docker build -t mrinjamul-admin:latest .
```

### Normal Build

For server,

```shell
go mod download
go build -o main .
```

For UI,

```shell
cd ui
npm install
touch .env # Write environment variables
npm run build
cp -rf build ../static
cd ..
```

## Running

### Run normally

```shell
touch .env # Write environment variables
source .env
./main
```

### Docker

```shell
docker run --rm -dp 3000:3000 --name myadmin mrinjamul-admin:latest
```

## Endpoints

| Methods | Endpoints         | Description                                            |
| ------- | ----------------- | ------------------------------------------------------ |
| GET     | /api/ping         | Use for ping                                           |
| GET     | /api/projects     | fetch project informations                             |
| POST    | /api/messages     | send messeges to firestore                             |
| GET     | /api/messages     | fetch messeges from firestore (protected)              |
| DELETE  | /api/messages/:id | NEED TO IMPLEMENT (WIP) (delete a messege) (protected) |
| PUT     | /api/messages     | NEED TO IMPLEMENT (WIP) (mark as read) (protected)     |

## Author

- Injamul Mohammad Mollah <mrinjamul@gmail.com>

## License

- under [MIT license](https://github.com/mrinjamul/mrinjamul-admin/blob/master/LICENSE)
