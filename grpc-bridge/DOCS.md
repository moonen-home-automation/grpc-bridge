# gRPC Bridge Addon

gRPC Bridge is a service that accepts incoming gRPC request and passes them on to the Home Assistant WebSocket API.
This is used by other services and offers a simplified way of communicating with Home Assistant.

## Installation

The installation of this addon is pretty straightforward and not different in comparison to installing any other Home Assistant addon.

### Using My button (fastest)

1. Click the Home Assistant My button below to open the add-on on your Home
   Assistant instance.

   [![Open this add-on in your Home Assistant instance.][addon-badge]][addon]

2. Click the "Install" button to install the add-on.
3. Start the "gRPC Bridge" add-on.
4. Check the logs of "gRPC Bridge" to see if everything went well.

### Manual

1. Navigate to the Add-on Store of your Home Assistant instance. **Settings**->**Add-ons**->**Add-on Store**
2. Add the following repository https://github.com/moonen-home-automation/repository by clicking the hamburger menu at the top-right and clicking **Repositories**.
3. Search for the **gRPC Bridge** addon in the addon store and click it.
4. Click the "Install" button to install the add-on.
5. Start the "gRPC Bridge" add-on.
6. Check the logs of "gRPC Bridge" to see if everything went well.

## Usage

### Proto file

The proto file for the gRPC Bridge can be found [here](https://github.com/moonen-home-automation/proto/blob/main/hass.proto).
Use this file to integrate with the gRPC Bridge and generate client code.

This file can also be imported into tools like Postman for testing purposes.

## Changelog & Releases

This repository keeps a change log using [GitHub's releases][releases]
functionality.

Releases are based on [Semantic Versioning][semver], and use the format
of `MAJOR.MINOR.PATCH`. In a nutshell, the version will be incremented
based on the following:

- `MAJOR`: Incompatible or major changes.
- `MINOR`: Backwards-compatible new features and enhancements.
- `PATCH`: Backwards-compatible bugfixes and package updates.

## Support

For questions and bugs please [create an issue here][issues] GitHub

## Authors & contributors

The original setup of this repository is by [Rick Moonen][rickmoonex].

## License

License can be found [here][license]

[license]: https://github.com/moonen-home-automation/grpc-bridge/blob/main/LICENSE
[issues]: https://github.com/moonen-home-automation/grpc-bridge/issues
[releases]: https://github.com/moonen-home-automation/grpc-bridge/releases
[rickmoonex]: https://github.com/rickmoonex
[semver]: https://semver.org/spec/v2.0.0.html
[addon-badge]: https://my.home-assistant.io/badges/supervisor_addon.svg
[addon]: https://my.home-assistant.io/redirect/supervisor_addon/?addon=0ed99dab_grpc-bridge&repository_url=https%3A%2F%2Fgithub.com%2Fmoonen-home-automation%2Frepository
