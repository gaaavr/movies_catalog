Feature: register and login
  Scenario: register
    When get "EMAIL" and "PASSWORD"
    Then there should be "1ssg@yandex.ru" and "1qwerty" remaining