# v2.0.0

## Features
- Support for `OAuth2` client config
- Added support for Keyfactor client config file with `OAuth2` client config(s)

## Breaking Changes
- fix: Return errors if client setup fails rather than `log.fatalf`
- fix: `NewAPIClient` returns `(*APIClient, error)`

# v1.0.0
- Initial release