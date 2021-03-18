# Hello Arduino!

Hardware was a side of technology that I never learned about, other than what I learned from my CSA501 (computer system architecture) class, which was the components of a computer. My brother is the one who was always into hardware and gave me his old arduino uno, along with a box of various bits (LEDs, various sensors, ect). Up until recently I never really took the time to play around with it, partially because I didnt know any C, and also because I never really knew how to approach it. With my recent exposure to [nrf5](/posts/programming/c/nrf5) via [Sam](https://github.com/pigeonhands), I have decided to dig into the arduino as a means to learn more.

Much like the nrf5 stuff, the first thing needed is the IDE, which arduino provides two options. The first is a web editor, and the second is an open source IDE available on the arduino websites [download page](https://www.arduino.cc/en/Main/Software). I will be using the arduino IDE and am using ubuntu 20.04, so I downloaded the Linux x64 file. Alternatively if you are using ubuntu, you can download the IDE via `apt install arduino`.

Like anything in programming its time to start with a classic `hello world!`. Since this is (embedded) hardware, the hello world equivelant (in my opinion) would be getting an LED to turn on and off. The arduino IDE has a bunch of examples, one of which is for making an LED turn on and off. Lets grok the code.

```C
/*
  Blink
  Turns on an LED on for one second, then off for one second, repeatedly.
 
  This example code is in the public domain.
 */
 
// Pin 13 has an LED connected on most Arduino boards.
// give it a name:
int led = 13;

// the setup routine runs once when you press reset:
void setup() {                
  // initialize the digital pin as an output.
  pinMode(led, OUTPUT);     
}

// the loop routine runs over and over again forever:
void loop() {
  digitalWrite(led, HIGH);   // turn the LED on (HIGH is the voltage level)
  delay(1000);               // wait for a second
  digitalWrite(led, LOW);    // turn the LED off by making the voltage LOW
  delay(1000);               // wait for a second
}
```

The first thing thats happening is an int variable called `led` is created with the value `13`. This is because the LED will be plugged into GPIO pin #13. GPIO stands for [*general-purpose input/output*](https://www.egr.msu.edu/classes/ece480/capstone/fall09/group03/AN_balachandran.pdf). The reason behind the use of the specific GPIO pin (#13) is because its right next to a GND (ground) pin. This means you can plug the LED directly into the board: positive (the longer end) plugged into the #13 pin, and negative (the shorter end) plugged into the GND pin.

The `setup` function is calling the `pinMode` function, which as the documentation says is initializing the digital pin as an output. Let look at the [documentation for pinMode](https://www.arduino.cc/reference/en/language/functions/digital-io/pinmode/) to understand a bit more about it. According to the documentation, pinMode "configures the specified pin to behave as an input or an output." In this case `led` (aka pin #13) is being set as `OUTPUT`.

Finally in the `loop` function we are calling the `digitalWrite` function and the `delay` function. The [digitalWrite function](https://www.arduino.cc/reference/en/language/functions/digital-io/digitalwrite/), according to the documentation, "writes a `HIGH` or `LOW` value to a digital pin." Since in this case the pin is configured as an `OUTPUT`, "its voltage will be set to the corresponding value: 5V (or 3.3V on 3.3V boards) for `HIGH`, 0V (ground) for `LOW`." The [delay function](https://www.arduino.cc/reference/en/language/functions/time/delay/) is pretty self-explanatory, it "pauses the program for the amount of time (in milliseconds) specified as parameter."