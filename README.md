# XTB - GO CLIENT LIBRARY

## Important Reminder (as of 14 March 2025) !!!

Due to XTB API access being disabled as of 14.03.2025, any code within this project relying on the XTB API will currently not function as intended. Please refrain from using or integrating XTB API calls until further notice. We are actively exploring alternative solutions and will update this documentation with any changes or new instructions as they become available.

---

For detailed API documentation and official support, refer to the [XTB API Documentation](http://developers.xstore.pro/documentation/). It provides comprehensive information on how to use the various endpoints, authentication methods, and other advanced features offered by XTB's trading platform.

## Example Usage and Initialization of demo environment

```
	c, err := xclient.NewClient("demo", userID, password)
	if err != nil {
		log.Fatalf("Failed to initialize client: %v", err)
	}

	defer c.Close()

```

## Disclaimer

This library provides a convenient interface for trading with XTB but use it at your own risk. Automated trading can lead to losses, and markets are inherently volatile. It is recommended that you fully understand how the API works before executing live trades. Always test your strategies thoroughly in demo environments and consult with a financial advisor if needed.

The developers of this library are not responsible for any financial losses or damages that may result from using this software.
