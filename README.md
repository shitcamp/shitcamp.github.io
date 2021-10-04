# Shitcamp

Shitcamp was a 4-day collaborative live streaming event that was hosted by a bunch of popular Twitch streamers like xQcOW, HasanAbi, Ludwig, Sodapoppin, QTCinderella any many more. They came together to participate in challenges and games, and ultimately to create the best content ever.

This repo contains the server and client code that is used for making the Shitcamp website: https://shitcamp.github.io

## Getting started

### Running the API server

The API server is writted using Go.
Run the server using the below commands.
You should be able to access the API server at `localhost:9090`.

```
cd cmd/shitcamp-api
go run . --config=../config.json --log=../log/shitcamp-api.log

# OR run the compiled binary
cd cmd/shitcamp-api
go build .
./shitcamp-api --config=../config.json --log=../log/shitcamp-api.log
```

### Running the client app

The web app is written using JavaScript and React.
Run the web app using the below commands.
You should be able to access the app at `localhost:3000`.

```
cd web
npm run start
```

## Deployment

The server is hosted using the DigitalOcean App Platform for the sake of simplicity, ease of use and low cost.
This does provide less flexibility though, since a lot of stuff (like metrics, logs) are auto-managed.

The website is hosted at `shitcamp.github.io` and managed by GitHub pages since it is free-of-cost, and fairly easy to remember.

### Deploying the API server

Use the below script to deploy the server to the production environment.
This script fetches the latest server logs, and pushes to the `main` branch of the repo used for deployment.
This triggers a deployment, since the DigitalOcean app has been configured to start a new deployment upon pushing new commits.

```
./deploy-server.sh
```

### Deploying the client app

Use the below script to deploy a new version of the web app.
The script updates the app version, installs dependencies and builds the app and then pushes it to the repo that has GitHub Pages enabled.
After the deployment, you may need to wait a couple of minutes or force refresh the page in order for the changes to become effective.

```
./deploy-client.sh
```

## Notes

- The API server unfortunately does not use any persistent storage such as a DB.
  This was done for the sake of easy setup, usage and deployment.

  However, this means that it cannot be scaled horizontally to increase the server capacity.
  Instead, all the data that should have been stored in a persistent store is initialised in `pkg/data/const.go`.
  So every time you make updates using the POST APIs, you need to update it in this "store" too.
  Otherwise, you will lose all your changes when you restart the API server.

- The React web app code is a bit messy since a lot of changes were made over a short period of time, including during the Shitcamp events, to improve existing features and add new ones that would improve the end-user experience.
- Displaying times in users' time zones, and displaying the amount of time left for an event to start, were simple yet much loved features.
- The APIs return responses in a JSON object format. This provides flexibility and makes it extremely easy to make backwards compatible changes to the API responses.

## Future Improvements

- Use a persistent store such as a DB for the API server
  - This DB can also be used to store data retrieved using the Twitch APIs, which can be done by a background sync job
- Clips page
  - Add pagination or infinite scrolling
  - Allow users to sort clips by new or popularity
- Dark mode for the web app
