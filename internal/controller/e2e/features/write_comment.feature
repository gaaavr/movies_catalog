Feature: write comment for movie
  Scenario: register, login and write comment
    When user registers with "REGISTER_EMAIL" and "REGISTER_PASSWORD"
    Then should be register successfully and receive code 200
    When the user tries to log in with "REGISTER_EMAIL" and "REGISTER_PASSWORD"
    Then he receives response code 200, non-empty state and send confirmation code by email
    When the user enters the verification code from "REGISTER_EMAIL" with pass "EMAIL_APP_PASS" and state
    Then he receive response code 200 and a non-empty token
    When user with a token leaves a comment "Good movie" on a movie with ID 4
    Then he receives a response code of 200 in response