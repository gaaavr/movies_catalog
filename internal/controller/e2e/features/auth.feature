Feature: register and login
  Scenario: register
    When get "EMAIL" and "PASSWORD"
    Then there should be "ssg@yandex.ru" and "qwerty" remaining