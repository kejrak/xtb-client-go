# XTB - GO CLIENT LIBRARY

For detailed API documentation and official support, refer to the [XTB API Documentation](http://developers.xstore.pro/documentation/). It provides comprehensive information on how to use the various endpoints, authentication methods, and other advanced features offered by XTB's trading platform.

## Example Usage and Initialization

```
	serverURL := "wss://ws.xtb.com/real"    
	c, err := xtb.NewClient(serverURL, nil) 
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer c.Close()

	userID := <USERID>                      
	password := <PASSWORD>

	_, err = c.Login(userID, password)
	if err != nil {
		log.Fatalf("Login failed: %v", err)
	}
```

## Disclaimer

This library provides a convenient interface for trading with XTB but use it at your own risk. Automated trading can lead to losses, and markets are inherently volatile. It is recommended that you fully understand how the API works before executing live trades. Always test your strategies thoroughly in demo environments and consult with a financial advisor if needed.

The developers of this library are not responsible for any financial losses or damages that may result from using this software.