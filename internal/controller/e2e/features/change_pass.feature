Feature: user change pass with otp
  Scenario: user change his password with confirm code
    When user with email "CHANGE_PASS_EMAIL" change password from "CHANGE_OLD_PASS" to "CHANGE_NEW_PASS"
    Then he receives response code 200, non-empty state and send confirmation code by email
    When the user enters the verification code from "CHANGE_PASS_EMAIL" with pass "CHANGE_PASS_APP_PASS" and state
    Then he receive response code 200