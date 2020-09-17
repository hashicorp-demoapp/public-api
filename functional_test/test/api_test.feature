Feature: Test the GraphQL API
  In order to ensure the latest code works in an integration setting
  run some automated tests against the public GraphQL API

Scenario: Test Query
  Given I have a running blueprint
  When I run the script
    ```
    #!/bin/bash
    curl 'http://localhost:8080/api' \
      -H 'Accept-Encoding: gzip, deflate, br' \
      -H 'Content-Type: application/json' \
      -H 'Accept: application/json' \
      -H 'Connection: keep-alive' \
      -H 'DNT: 1' \
      -H 'Origin: http://localhost:8080' \
      --data-binary '{"query":"{\n  coffees{id}\n}"}' \
      --compressed
    ```
  Then I expect the exit code to be 0

Scenario: Test Mutation
  Given I have a running blueprint
  When I run the script
    ```
    #!/bin/bash
    curl 'http://localhost:8080/api' \
    -H 'Accept-Encoding: gzip, deflate, br' \
    -H 'Content-Type: application/json' \
    -H 'Accept: application/json' \
    -H 'Connection: keep-alive' \
    -H 'DNT: 1' \
    -H 'Origin: http://localhost:8080' \
    --data-binary '{"query":"mutation {\n  pay(\n    details: {\n      name: \"nic\",\n      type: \"mastercard\",\n      number: \"1234-1231-1231-2322\"\n      expiry: \"10/02\",\n      cv2: 123,\n      amount: 10.23,\n    }\n  )  \n}"}'\
     --compressed
    ```
  Then I expect the exit code to be 0