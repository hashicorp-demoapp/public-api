# public-api
A GraphQL public API for the demo app

# Creating a new release
The build pipeline is setup with Circle CI to build and create a new Docker image whenever a new tag is pushed to this repo. To create a new release execute the following commands:

```shell
# Use sem var for tags, i.e. v0.0.1
git tag [tag]
git push origin [tag]
```

# Payments API interaction

```json
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
  )  
}
```