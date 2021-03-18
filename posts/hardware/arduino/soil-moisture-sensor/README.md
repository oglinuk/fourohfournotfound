# YL-38/69 Soil Moisture Sensor

One of the various sensors I have to play with is the YL-38/69 soil moisture sensor. There are two components to this sensor: the YL-38 controller, and the YL-69 sensor. In this blog we are going to light up LEDs (yellow, blue, and green) to indicate the soils moisture level. First thing to do is connect the sensor and the LEDs to the arduino. The YL-38 controller is plugged directly into a breadboard, VCC connected to 5V pin, the analog output connected to the analog A0 pin, and GND connected to the another section of the breadboard. Each LED is fitted with a 10k ohm resistor to the negative, then plugged into the breadboard in the same column as the GND wire from the sensor, and the positive is plugged into a GPIO pin on the arduino. The yellow LED is plugged into the GPIO pin #2, blue into #3, and green into #4. Now for the code.

```C
int yellow = 2;
int blue = 3;
int green = 4;

void setup()
{
    Serial.begin(9600);
    pinMode(yellow, OUTPUT);
    pinMode(blue, OUTPUT);
    pinMode(green, OUTPUT);
}

void loop()
{
    int sensor_level = analogRead(A0);

    if (sensor_level >= 750) {
        digitalWrite(yellow, HIGH);
    } else if (sensor_level >= 500 && sensor_level <= 750) {
        digitialWrite(blue, HIGH);
    } else {
        digitalWrite(green, HIGH);
    }

    delay(300);
}
```

Lets grok whats going on. The first thing thats happening is the initialization of three int variables: `yellow`, `blue`, and `green`. Then the `setup` function calls `Serial.begin(9600)`, which according to the documentation for [begin](https://www.arduino.cc/reference/en/language/functions/communication/serial/begin/), "sets the data rate in bits per second (baud) for serial data transmission". In this case we are sending `9600` bits per second. The next thing happening is configuring each of the three pins to be in `OUTPUT` mode.

The `loop` function first sets the value of a variable called `sensor_level` to the returned value of the `analogRead(A0)` function. The [analogRead](https://www.arduino.cc/reference/en/language/functions/analog-io/analogread/) function "reads the value from the specified analog pin. Arduino boards contain a multichannel, 10-bit analog to digital converter. This means that it will map input voltages between 0 and the operating voltage(5V or 3.3V) into integer values between 0 and 1023. On an Arduino UNO, for example, this yields a resolution between readings of: 5 volts / 1024 units or, 0.0049 volts (4.9 mV) per unit." We are then checking the value of `sensor_level`, setting the voltage of an LED that corresponds (yellow for anything higher than 750, blue for anything between 500-750, and green for anything less than 500). Finally we `delay` for 300 milliseconds.

Time to test it!