Feature: Read an existing user

  Scenario: Create a valid non existing user
    Given I am on / page
    Then I click the Create tab
    * I fill the Name with read
    * I select the Role in Administrator
    * I check the I agree to the terms and conditions
    * I click the Create button
    And I see user created notification

  Scenario: Read a valid existing user
    Given I am on / page
    Then I click the Read tab
    * I fill the Name with read
    * I click the Read button
    And I see user found notification

