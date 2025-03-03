<h1 align="center">Domainpeek?</h1>
<p align="center">
<i>Free & Open Source WHOIS Lookup Service</i>
<br />
<b>🌐 <a href="https:/domainseek.net/">domainseek.as93.net</a></b><br />
</p>

---

## API Usage

> **TL;DR** Get the WHOIS records for a site: `curl https:/domainseek.net/example.com`

For detailed request and response schemas, you can reference the [domainseek API Spec]()

### Base URL

The base URL for the public API is [`domainseek.as93.net`](https:/domainseek.net)

If you're self-hosting (reccomended) then replace this with your own base URL.

### Endpoints

<details>
  <summary><h4>Single Domain Lookup <code>/[domain]</code></h4></summary>

- **URL**: `/[domain]`
- **Method**: `GET`
- **URL Params**: None
- **Success Response**:
  - **Code**: 200
  - **Content**: WHOIS data for the specified domain in JSON format.
- **Error Response**:
  - **Code**: 400 BAD REQUEST
  - **Content**: `{ "error": "Domain not specified" }`
  - **Code**: 404 NOT FOUND
  - **Content**: `{ "error": "Domain not found" }`
- **Sample Call**:

```
curl https://your-api-url.com/example.com`
```

</details>

<details>
  <summary><h4>Multiple Domain Lookup <code>/multi</code></h4></summary>

- **URL**: `/multi`
- **Method**: `GET`
- **Query Params**: 
  - **domains**: A comma-separated list of domains.
- **Success Response**:
  - **Code**: 200
  - **Content**: Array of WHOIS data for the specified domains in JSON format.
- **Error Response**:
  - **Code**: 400 BAD REQUEST
  - **Content**: `{ "error": "No domains specified" }`
  - **Code**: 500 INTERNAL SERVER ERROR
  - **Content**: `{ "error": "[error message]" }`
- **Sample Call**:

```
curl "https://your-api-url.com/multi?domains=example.com,example.net"
```

</details>

---

## Deployment

#### Option 1: Vercel

Click the button below to deploy to Vercel 👇

[![Deploy with Vercel](https://vercel.com/button)](https://vercel.com/new/clone?repository-url=https%3A%2F%2Fgithub.com%2Fobaraelijah%2Fdomainseek&demo-title=domainseek%20Demo&demo-url=https%3A%2F%2Fdomainseek.as93.net&demo-image=https%3A%2F%2Fi.ibb.co%2FJ5r1zCP%2Fdomainseek-square.png)

#### Option 2: Docker

```shell
docker run -p 8080:8080 obaraelijah/domainseek
```

#### Option 3: Binary

Head to the [Releases Tab](/releases), download the pre-built executable for your system, then run it.

#### Option 4: Build from Source

Follow the setup instructions in the [Development](#development) section.<br>
Then run `go build -a -installsuffix cgo -o domainseek .` to generate the binary for your system.<br>
You'll then be able to execute the newly built `./domainseek` file directly to start the application.

---

## Development

Prerequisites: You'll need [Go](https://go.dev/) and [Node](https://nodejs.org/) installed. You will likley also want [Git](https://git-scm.com/) and/or [Docker](https://www.docker.com/).

```
git clone git@github.com:obaraelijah/domainseek.git
cd domainseek
go get
npm install
npm run build
```

Then run either `npx vercel dev`, or `go run main.go`

Alternativley, build the Docker container with `docker build -t domainseek .`

---

## Credits

##### Contributors

##### Sponsors

---

## License

> _**[obaraelijah/domainseek](https://github.com/obaraelijah/domainseek)** is licensed under [MIT](https://github.com/obaraelijah/domainseek/blob/HEAD/LICENSE) © [Elijah Samson](https://elijahobara.com) 2025._<br>
> <sup align="right">For information, see <a href="https://tldrlegal.com/license/mit-license">TLDR Legal > MIT</a></sup>

