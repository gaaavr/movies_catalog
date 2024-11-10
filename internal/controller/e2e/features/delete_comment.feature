Feature: error delete comment for movie
  Scenario: unauthorized user with wrong token delete a comment
    When user with a "token" tries to delete a comment with ID 1
    Then he will get response code 401