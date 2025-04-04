# v24.0.0

## Features
- Support for Keyfactor Command REST API endpoints up to 24.4
- Support for v1 and v2 REST API endpoints

## Fixes
- fix: Fixes an issue where API calls are hardcoded to `/KeyfactorAPI`. Now, the API clients can point to API subpaths defined in your authentication configuration.

## Breaking Changes
- chore: Support for v1 and v2 REST endpoints results in new structure for calling APIs. The base API client exposes a `V1` and `V2` API client. Package names are different as a result.
- chore: Changes in the OpenAPI specification results in different method and class names. Please refer to the [v1](./v24/api/keyfactor/v1/README.md#documentation-for-api-endpoints) and [v2](./v24/api/keyfactor/v2/README.md#documentation-for-api-endpoints) documentation for which class / method to use.

# v2.0.0

## Features
- Support for `OAuth2` client config
- Added support for Keyfactor client config file with `OAuth2` client config(s)

## Breaking Changes
- fix: Return errors if client setup fails rather than `log.fatalf`
- fix: `NewAPIClient` returns `(*APIClient, error)`

# v1.0.0
- Initial release
