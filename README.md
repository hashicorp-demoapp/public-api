# public-api

A GraphQL public API for the demo app

[![CircleCI](https://circleci.com/gh/hashicorp-demoapp/public-api.svg?style=svg)](https://circleci.com/gh/hashicorp-demoapp/public-api)  

Docker Image: [https://hub.docker.com/repository/docker/hashicorpdemoapp/public-api](https://hub.docker.com/repository/docker/hashicorpdemoapp/public-api)

# Creating a new release
The build pipeline is setup with Circle CI to build and create a new Docker image whenever a new tag is pushed to this repo. To create a new release execute the following commands:

```shell
# Use sem var for tags, i.e. v0.0.1
git tag [tag]
git push origin [tag]
```

# Payments API interaction

You can interact with the graphql service using the ui which is accessible in your browser at the root path of the service.

## Request

```graphql
mutation {
  pay(
    details: {
      name: "nic",
      type: "mastercard",
      number: "1234-1231-1231-2322"
      expiry: "10/02",
      cv2: 123,
      amount: 10.23,
    }
  ){
    id,
    card_plaintext,
    card_ciphertext,
    message
  } 
}
```

## Response

```
{
  "data": {
    "pay": {
      "id": "40487460-ae6b-4b02-92ae-01ab24f738a4",
      "card_plaintext": "1234123-0123123",
      "card_ciphertext": "Encryption Disabled",
      "message": "Payment processed successfully, card details returned for demo purposes, not for production"
    }
  }
}
```

Or you can use cURL:

```shell
curl 'http://localhost:8080/api' \
      -H 'Accept-Encoding: gzip, deflate, br' \
      -H 'Content-Type: application/json' \
      -H 'Accept: application/json' \
      -H 'Connection: keep-alive' \
      -H 'DNT: 1' \
      -H 'Origin: http://localhost:8080' \
      --data-binary '{"query":"mutation{ pay(details:{ name: \"nic\", type: \"mastercard\", number: \"1234123-0123123\", expiry:\"10/02\",    cv2: 1231, amount: 12.23 }){id, card_plaintext, card_ciphertext, message } }"}' --compressed
```