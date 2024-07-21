Feature: Delete an existing user

  Scenario: Create a valid non existing account
    Given I am on / page
    Then I click the Create tab
    * I fill the Name with delete
    * I select the Role in Administrator
    * I check the I agree to the terms and conditions
    * I click the Create button
    And I see user created notification

  Scenario: Delete a valid existing account
    Given I am on / page
    Then I click the Delete tab
    * I fill the Name with delete
    * I click the Delete button
    And I see user deleted notification
