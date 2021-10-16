# file: $GOPATH/godogs/features/godogs.feature
Feature: test httpref

  Scenario: Fails with no parameters
    Given the args are ""
    Then return status code is not 0

  Scenario: Check that filtering from root works
    Given the args are "1xx"
    Then return status code is 0

  Scenario: Check that filtering from alias status works
    Given the args are "statuses 1xx"
    Then return status code is 0

  Scenario: Check that filtering from status works
    Given the args are "status 1xx"
    Then return status code is 0

  Scenario: Check that filtering from methods works
    Given the args are "methods GET"
    Then return status code is 0

  Scenario: Check that filtering from methods alias works
    Given the args are "method GET"
    Then return status code is 0

  Scenario: Check that filtering from headers works
    Given the args are "headers Accept-Ranges"
    Then return status code is 0

  Scenario: Check that filtering from headers alias works
    Given the args are "header Accept-Ranges"
    Then return status code is 0

  Scenario: Finds unique entry on exact match on root
    Given the args are "Accept"
    Then return status code is 0
    And the output will contain the Accept headers

  Scenario: Finds unique entry on exact match on headers
    Given the args are "Accept*"
    Then return status code is 0
    And the output will not contain the Accept headers

  Scenario: Finds unique entry on exact match on headers
    Given the args are "headers  Accept"
    Then return status code is 0
    And the output will contain the Accept headers

  Scenario: Full-text search option works on root
    Given the args are "--search clear"
    Then return status code is 0

  Scenario: Full-text search option works on headers
    Given the args are "headers --search cache"
    Then return status code is 0

  Scenario: Full-text search option works on methods
    Given the args are "methods --search cache"
    Then return status code is 0

  Scenario: Full-text search option works on statuses
    Given the args are "statuses --search cache"
    Then return status code is 0

  Scenario: Full-text search option does NOT work on ports
    Given the args are "ports --search cache"
    Then return status code is 1

  Scenario: Can change the width of the output
    Given the args are "-w 70 100"
    Then return status code is 0
    And each line in output is shorter than 70 characters

  Scenario: Ports do not appear in the normal searches
    Given the args are "port 80"
    Then return status code is 0
    And the output will contain the http string

  Scenario: You can look up a port that is in a range
     Given the args are "port 47809"
    Then return status code is 0
    And the output will not contain the http string 
