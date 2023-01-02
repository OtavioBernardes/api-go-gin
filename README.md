<h1 align="center">
   Go-Gin - API
</h1>

<p align="center">
  <img alt="Last commit on GitHub" src="https://img.shields.io/github/last-commit/OtavioBernardes/api-go-gin?color=7D40E7">
  <img alt="Made by Rocketseat" src="https://img.shields.io/badge/made%20by-OtavioBernardes-%20?color=7D40E7">
  <img alt="Project top programing language" src="https://img.shields.io/github/languages/top/OtavioBernardes/api-go-gin?color=7D40E7">
  <img alt="Repository size" src="https://img.shields.io/github/repo-size/OtavioBernardes/api-go-gin?color=7D40E7">
</p>

## :rocket: Getting Started

1. Clone this repo: `git clone https://github.com/OtavioBernardes/api-go-gin`
2. Move to the directory: `cd api-go-gin`
3. Run docker compose: `docker-compose -f ./development/docker-compose.yml up`
4. Run `go run src/main`
5. The server runs on: http://localhost:3000
6. Run a request to create a studend in:
```
curl --request POST \
  --url http://localhost:3000/student \
  --header 'Content-Type: application/json' \
  --data '{
	"name": "Fake Name",
	"email": "jhon@fake.com"
}'
```
## :page_facing_up: License

**Free Software, Hell Yeah!**
