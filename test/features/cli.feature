# file: $GOPATH/godogs/features/godogs.feature
Feature: test httpref

  Scenario: Fails with no parameters
    Given the app runs without parameters
    Then the app exits with an error

  Scenario: Check that filtering from root works
    Given the app runs with parameters "1xx"
    Then the app exits without error

  Scenario: Check that filtering from alias status works
    Given the app runs with parameters "statuses 1xx"
    Then the app exits without error

  Scenario: Check that filtering from status works
    Given the app runs with parameters "status 1xx"
    Then the app exits without error

  Scenario: Check that filtering from methods works
    Given the app runs with parameters "methods GET"
    Then the app exits without error

  Scenario: Check that filtering from methods alias works
    Given the app runs with parameters "method GET"
    Then the app exits without error

  Scenario: Check that filtering from headers works
    Given the app runs with parameters "headers Accept-Ranges"
    Then the app exits without error

  Scenario: Check that filtering from headers alias works
    Given the app runs with parameters "header Accept-Ranges"
    Then the app exits without error

  Scenario: Check that html filtering from root works
    Given the app runs with parameters "<abbr>"
    Then the app exits without error

  Scenario: Check that filtering from html works
    Given the app runs with parameters "html <abbr>"
    Then the app exits without error

  Scenario: Finds unique entry on exact match on root
    Given the app runs with parameters "Accept"
    Then the app exits without error
    And the app output contains "/docs/Web/HTTP/Headers/Accept"

  Scenario: Finds unique entry on exact match on headers
    Given the app runs with parameters "Accept*"
    Then the app exits without error
    And the app output does not contain "/docs/Web/HTTP/Headers/Accept"

  Scenario: Finds unique entry on exact match on headers
    Given the app runs with parameters "headers Accept"
    Then the app exits without error
    And the app output contains "/docs/Web/HTTP/Headers/Accept"

  Scenario: Finds unique html entry on exact match on root
    Given the app runs with parameters "<abbr>"
    Then the app exits without error
    And the app output contains "docs/Web/HTML/Element/abbr"

  Scenario: Finds unique entry on exact match on html
    Given the app runs with parameters "html <abbr>"
    Then the app exits without error
    And the app output contains "docs/Web/HTML/Element/abbr"

  Scenario: Full-text search option works on root
    Given the app runs with parameters "--search clear"
    Then the app exits without error

  Scenario: Full-text search option works on headers
    Given the app runs with parameters "headers --search cache"
    Then the app exits without error

  Scenario: Full-text search option works on methods
    Given the app runs with parameters "methods --search cache"
    Then the app exits without error

  Scenario: Full-text search option works on statuses
    Given the app runs with parameters "statuses --search cache"
    Then the app exits without error

  Scenario: Full-text search option does NOT work on ports
    Given the app runs with parameters "ports --search cache"
    Then the app exits with an error

  Scenario: Full-text search option works on html
    Given the app runs with parameters "html --search anchor"
    Then the app exits without error

  Scenario: Can change the width of the output
    Given the app runs with parameters "-w 70 100"
    Then the app exits without error

  Scenario: Ports do not appear in the normal searches
    Given the app runs with parameters "port 80"
    Then the app exits without error
    And the app output contains "Hypertext Transfer Protocol"

  Scenario: You can look up a port that is in a range
     Given the app runs with parameters "port 47809"
    Then the app exits without error
    And the app output does not contain "Hypertext Transfer Protocol" 
