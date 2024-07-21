Feature: Update an existing user

  Scenario: Create a valid non existing account
    Given I am on / page
    Then I click the Create tab
    * I fill the Name with update
    * I select the Role in Administrator
    * I check the I agree to the terms and conditions
    * I click the Create button
    And I see user created notification

  Scenario: Create already existing account
    Given I am on / page
    Then I click the Update tab
    * I fill the Name with update
    * I select the Role in Contributor
    * I click the Update button
    But I see user updated notification
