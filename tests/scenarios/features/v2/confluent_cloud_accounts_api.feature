@endpoint(confluent-cloud-accounts-api) @endpoint(confluent-cloud-accounts-api-v2)
Feature: Confluent Cloud Accounts API
  Configure your Datadog Confluent Cloud integration directly through the
  Datadog API.

  Background:
    Given a valid "apiKeyAuth" key in the system
    And a valid "appKeyAuth" key in the system
    And an instance of "ConfluentCloudAccountsAPI" API

  @generated @skip @team:Datadog/web-integrations
  Scenario: Add Confluent account returns "API error response." response
    Given new "ConfluentCloudPublic3" request
    And body with value {"data": {"attributes": {"api_key": "TESTAPIKEY123", "api_secret": "test-api-secret-123", "resources": [{"id": "resource_id_abc123", "resource_type": "kafka", "tags": ["myTag", "myTag2:myValue"]}], "tags": ["myTag", "myTag2:myValue"]}, "type": "confluent-cloud-accounts"}}
    When the request is sent
    Then the response status is 404 API error response.

  @generated @skip @team:Datadog/web-integrations
  Scenario: Add Confluent account returns "OK" response
    Given new "ConfluentCloudPublic3" request
    And body with value {"data": {"attributes": {"api_key": "TESTAPIKEY123", "api_secret": "test-api-secret-123", "resources": [{"id": "resource_id_abc123", "resource_type": "kafka", "tags": ["myTag", "myTag2:myValue"]}], "tags": ["myTag", "myTag2:myValue"]}, "type": "confluent-cloud-accounts"}}
    When the request is sent
    Then the response status is 201 OK

  @generated @skip @team:Datadog/web-integrations
  Scenario: Add resource to Confluent account returns "API error response." response
    Given new "ConfluentCloudPublic8" request
    And request contains "account_id" parameter from "REPLACE.ME"
    And body with value {"data": {"id": "resource_id_abc123", "resource_type": "kafka", "tags": ["myTag", "myTag2:myValue"]}}
    When the request is sent
    Then the response status is 404 API error response.

  @generated @skip @team:Datadog/web-integrations
  Scenario: Add resource to Confluent account returns "OK" response
    Given new "ConfluentCloudPublic8" request
    And request contains "account_id" parameter from "REPLACE.ME"
    And body with value {"data": {"id": "resource_id_abc123", "resource_type": "kafka", "tags": ["myTag", "myTag2:myValue"]}}
    When the request is sent
    Then the response status is 201 OK

  @generated @skip @team:Datadog/web-integrations
  Scenario: Delete Confluent account returns "API error response." response
    Given new "ConfluentCloudPublic5" request
    And request contains "account_id" parameter from "REPLACE.ME"
    When the request is sent
    Then the response status is 404 API error response.

  @generated @skip @team:Datadog/web-integrations
  Scenario: Delete Confluent account returns "OK" response
    Given new "ConfluentCloudPublic5" request
    And request contains "account_id" parameter from "REPLACE.ME"
    When the request is sent
    Then the response status is 204 OK

  @generated @skip @team:Datadog/web-integrations
  Scenario: Delete resource from Confluent account returns "API error response." response
    Given new "ConfluentCloudPublic10" request
    And request contains "account_id" parameter from "REPLACE.ME"
    And request contains "resource_id" parameter from "REPLACE.ME"
    When the request is sent
    Then the response status is 404 API error response.

  @generated @skip @team:Datadog/web-integrations
  Scenario: Delete resource from Confluent account returns "OK" response
    Given new "ConfluentCloudPublic10" request
    And request contains "account_id" parameter from "REPLACE.ME"
    And request contains "resource_id" parameter from "REPLACE.ME"
    When the request is sent
    Then the response status is 204 OK

  @generated @skip @team:Datadog/web-integrations
  Scenario: Get Confluent account returns "API error response." response
    Given new "ConfluentCloudPublic2" request
    And request contains "account_id" parameter from "REPLACE.ME"
    When the request is sent
    Then the response status is 404 API error response.

  @generated @skip @team:Datadog/web-integrations
  Scenario: Get Confluent account returns "OK" response
    Given new "ConfluentCloudPublic2" request
    And request contains "account_id" parameter from "REPLACE.ME"
    When the request is sent
    Then the response status is 200 OK

  @generated @skip @team:Datadog/web-integrations
  Scenario: Get resource from Confluent account returns "API error response." response
    Given new "ConfluentCloudPublic7" request
    And request contains "account_id" parameter from "REPLACE.ME"
    And request contains "resource_id" parameter from "REPLACE.ME"
    When the request is sent
    Then the response status is 404 API error response.

  @generated @skip @team:Datadog/web-integrations
  Scenario: Get resource from Confluent account returns "OK" response
    Given new "ConfluentCloudPublic7" request
    And request contains "account_id" parameter from "REPLACE.ME"
    And request contains "resource_id" parameter from "REPLACE.ME"
    When the request is sent
    Then the response status is 200 OK

  @generated @skip @team:Datadog/web-integrations
  Scenario: List Confluent Account resources returns "API error response." response
    Given new "ConfluentCloudPublic6" request
    And request contains "account_id" parameter from "REPLACE.ME"
    When the request is sent
    Then the response status is 404 API error response.

  @generated @skip @team:Datadog/web-integrations
  Scenario: List Confluent Account resources returns "OK" response
    Given new "ConfluentCloudPublic6" request
    And request contains "account_id" parameter from "REPLACE.ME"
    When the request is sent
    Then the response status is 200 OK

  @generated @skip @team:Datadog/web-integrations
  Scenario: List Confluent accounts returns "API error response." response
    Given new "ConfluentCloudPublic1" request
    When the request is sent
    Then the response status is 404 API error response.

  @team:Datadog/web-integrations
  Scenario: List Confluent accounts returns "OK" response
    Given there is a valid "confluent_account" in the system
    And new "ConfluentCloudPublic1" request
    When the request is sent
    Then the response status is 200 OK

  @generated @skip @team:Datadog/web-integrations
  Scenario: Update Confluent account returns "API error response." response
    Given new "ConfluentCloudPublic4" request
    And request contains "account_id" parameter from "REPLACE.ME"
    And body with value {"data": {"attributes": {"api_key": "TESTAPIKEY123", "api_secret": "test-api-secret-123", "tags": ["myTag", "myTag2:myValue"]}, "type": "confluent-cloud-accounts"}}
    When the request is sent
    Then the response status is 404 API error response.

  @generated @skip @team:Datadog/web-integrations
  Scenario: Update Confluent account returns "OK" response
    Given new "ConfluentCloudPublic4" request
    And request contains "account_id" parameter from "REPLACE.ME"
    And body with value {"data": {"attributes": {"api_key": "TESTAPIKEY123", "api_secret": "test-api-secret-123", "tags": ["myTag", "myTag2:myValue"]}, "type": "confluent-cloud-accounts"}}
    When the request is sent
    Then the response status is 201 OK

  @generated @skip @team:Datadog/web-integrations
  Scenario: Update resource in Confluent account returns "API error response." response
    Given new "ConfluentCloudPublic9" request
    And request contains "account_id" parameter from "REPLACE.ME"
    And request contains "resource_id" parameter from "REPLACE.ME"
    And body with value {"data": {"attributes": {"resource_type": "kafka", "tags": ["myTag", "myTag2:myValue"]}, "id": "resource-id-123", "type": "confluent-cloud-resources"}}
    When the request is sent
    Then the response status is 404 API error response.

  @generated @skip @team:Datadog/web-integrations
  Scenario: Update resource in Confluent account returns "OK" response
    Given new "ConfluentCloudPublic9" request
    And request contains "account_id" parameter from "REPLACE.ME"
    And request contains "resource_id" parameter from "REPLACE.ME"
    And body with value {"data": {"attributes": {"resource_type": "kafka", "tags": ["myTag", "myTag2:myValue"]}, "id": "resource-id-123", "type": "confluent-cloud-resources"}}
    When the request is sent
    Then the response status is 201 OK
