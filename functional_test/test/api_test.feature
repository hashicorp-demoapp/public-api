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
      --data-binary '{"query":"mutation{ pay(details:{ name: \"nic\", type: \"mastercard\", number: \"1234123-0123123\", expiry:\"10/02\",    cv2: 1231, amount: 12.23 }){id, card_plaintext, card_ciphertext, message } }"}' --compressed
    ```
  Then I expect the exit code to be 0