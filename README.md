# example-plugin

This is an example plugin that demonstrates one of the ways to [extend a Service Manager](https://github.com/Peripli/service-manager/blob/master/docs/plugins.md).

What `myplugin` does is intercept OSB catalog requests and modify the response.

## Create the plugin

We create `MyPlugin` in its own package `myplugin`.

A plugin is a struct that implements the `Name` method and at least one OSB operation - `FetchCatalog` in this example.

By implementing OSB operation handlers, we can intercept the corresponding requests and:
* add custom logic before and after the call
* modify the request and response objects
* handle errors
* or even not forward the call at all

In this example we implement the `FetchCatalog` method which passes catalog requests to the service broker and then modifies the response by adding a suffix to the first service's name.

We then create a `main.go` file that imports `myplugin` and the service manager.
It creates a new service manager, registers the plugin and runs the server.

## Test it works

You can manually test that the plugin works correctly:

1. Deploy the extended service manager we created (on pcfdev for example)

2. Register the service broker in the SM

3. Register a new platform in the SM. You will need the basic credentials from the response

4. Use the broker's credentials to request the original catalog directly through the broker's url

5. Use the credentials of the platform you created to request the modified catalog through the OSB API of the SM

   You should see the modified catalog with a suffix on the first service name
