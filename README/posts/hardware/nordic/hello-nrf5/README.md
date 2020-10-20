---
title: Hello nRF5!
draft: false
---

The context for this series of blogs should be of no surprise for how it came about. Once again I was introduced to an awesome thing by my good friend [Sam](https://github.com/pigeonhands). This particular topic is about the bluetooth protocol, and a chip manufacturer called Nordic. Sam works with the Nordic chips as an embedded developer, and has a couple development boards that he plays around with in his spare time. This blog is going to be about getting started with the Nordic bluetooth stack and the nRF52 DK.

First thing was to get the development environment installed, which Sam suggested [Segger embedded studio for ARM](https://www.segger.com/downloads/embedded-studio/). Fortunately, Segger offers a `.deb` file, which is easy to install on my laptop running Ubuntu via `dpkg`. If you have ever played around with an arduino, the IDE for the arduino is similar to the Segger embedded studio. Next thing to do was to download the SDK (v17.0.2) and the S132 softdevice download. With everything downloaded and ready to go, its time to play.

Sam started with the `ble_app_blinky_pca10040_s132` example project, which when ran, notifies the phone when a button is pressed on the  nRF52 DK. Lets deconstruct this example.

```C
int main(void)
{
    // Initialize.
    log_init();
    leds_init();
    timers_init();
    buttons_init();
    power_management_init();
    ble_stack_init();
    gap_params_init();
    gatt_init();
    services_init();
    advertising_init();
    conn_params_init();

    // Start execution.
    NRF_LOG_INFO("Blinky example started.");
    advertising_start();

    // Enter main loop.
    for (;;)
    {
        idle_state_handle();
    }
}
```

* `log_init` initializes logs.

* `leds_init` is for the LEDs initialization. It initializes all LEDs used by the application.

* `timers_init` is for the Timer initialization. It initializes the timer module.

* `buttons_init` is for initializing the button handler module.

* `power_management_init` is for initializing the power management. 

* `ble_stack_init` is for initializing the BLE stack. It initializes the SoftDevice and the BLE event interrupt.

*  `gap_params_init` is for the GAP initialization. It sets up all the necessary GAP (Generic Access Profile) parameters of the device including the device name, appearance, and the preferred connection parameters.

*  `gatt_init` is for initializing the GATT module.

*  `services_init` is for initializing services that will be used by the application.

* `advertising_init` is for initializing the Advertising functionality.It encodes the required advertising data and passes it to the stack. It also builds a structure to be passed to the stack when starting advertising.

* `conn_params_init` is for initializing the Connection Parameters module.

* `advertising_start` is for starting advertising. 

* `idle_state_handle` is for handling the idle state (main loop). If there is no pending log operation, then sleep until next the next event occurs.

There is a few bits of jargon so I am going to grok what I havent come across before. 

The first piece of jargon that is new to me is GAP or Generic Access Profile. GAP "controls connections and advertising in Bluetooth. GAP is what makes your device visible to the outside world, and determines how two devices can (or can't) interact with each other." [[source](https://learn.adafruit.com/introduction-to-bluetooth-low-energy/gap)] GAP has two core concepts which are called **Central** and **Peripheral** devices. A peripheral device is a device that connects to a central device, in this case the peripheral device is the nRF52 DK). A central device is something with considerable processing power and memory, such as a smart phone. There are two ways of advertising with GAP: **advertising data** and **scan response**. "Both payloads are identical and can contain up to 31 bytes of data, but only the advertising data payload is mandatory, since this is the payload that will be constantly transmitted out from the device to let central devices in range know that it exists. The scan response payload is an optional secondary payload that central devices can request, and allows device designers to fit a bit more information in the advertising payload such a strings for a device name, etc." [[source](https://learn.adafruit.com/introduction-to-bluetooth-low-energy/gap)]

The second piece of new jargon is GATT or Generic Attribute Profile. GATT "defines the way that two Bluetooth Low Energy devices transfer data back and forth using concepts called **Services** and **Characteristics**. Services are used to break data up into logic entities, and contain specific chunks of data called characteristics. A service can have one or more characteristics, and each service distinguishes itself from other services by means of a unique numeric ID called a UUID, which can be either 16-bit (for officially adopted BLE Services) or 128-bit (for custom services). The lowest level concept in GATT transactions is the Characteristic, which encapsulates a single data point (though it may contain an array of related data, such as X/Y/Z values from a 3-axis accelerometer, etc.)." [[source](https://learn.adafruit.com/introduction-to-bluetooth-low-energy/gatt)]

Interesting stuff!