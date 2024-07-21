Feature: Create a new user

  Scenario: Create a valid non existing user
    Given I am on / page
    Then I click the Create tab
    * I fill the Name with create
    * I select the Role in Administrator
    * I check the I agree to the terms and conditions
    * I click the Create button
    And I see user created notification

  Scenario: Create already existing user
    Given I am on / page
    Then I click the Create tab
    * I fill the Name with create
    * I select the Role in Administrator
    * I check the I agree to the terms and conditions
    * I click the Create button
    But I see already registered notification
