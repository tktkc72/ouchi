# ouchidashboard

store and display data related to `ouchi` life

## Description

This collects data related `ouchi` life from `nature remo`
and store them in `Firestore`

## Test

### Unit Test

`make test`

### Integration Test

with a firestore emulator

`docker-compose run --service-ports -d firestore`

`FIRESTORE_EMULATOR_HOST=localhost:8812 GCP_PROJECT="test" FIRESTORE_DOC_PATH="test" make integration_test`

### E2E Test

Test from a client to an api server on a container with a firestore emulator.

You need a real nature remo device.
And set the access token and the device id to environments.

`export NATURE_REMO_ACCESS_TOKEN=${NATURE_REMO_ACCESS_TOKEN}`
`export NATURE_REMO_DEVICE_ID=${NATURE_REMO_DEVICE_ID}`
`export FIRESTORE_EMULATOR_HOST=firestore:8812`
`docker-compose up -d`
`make e2e_test`
`docker-compose down`

## CI

### test

[github actions](..github/workflows/test.yml)

## CD

[Continuous deployment from Git using Cloud Build](https://cloud.google.com/run/docs/continuous-deployment-with-cloud-build?hl=ja#new-service)

## firebase deploy

need to install firebase CLI

`firebase deploy`

## setup

cloud scheduler

``` shell
gcloud beta scheduler jobs update http ouchi-dashboard-collector  --schedule=${SCHEDULE} --uri="${CLOUD_RUN_URI}" --message-body='{"roomNames":["living", "study"]}' --oidc-service-account-email=${SERVICE_ACCOUNT} --update-headers=Content-Type=application/json
```
