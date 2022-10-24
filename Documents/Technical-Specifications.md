<h1 style="text-align: center">Project 1 smart signage - group 4 - Technical Specifications</h1>

<details> 
<summary style="text-decoration: underline; font-size:150%">Table of contents:</summary>

- [1. How does it have to work:](#1-how-does-it-have-to-work)
	- [1.1. Radio part:](#11-radio-part)
	- [1.2. Coding part:](#12-coding-part)
	- [1.3. Electronic part:](#13-electronic-part)
- [2. The possibilities:](#2-the-possibilities)
	- [2.1. What can the device do:](#21-what-can-the-device-do)
	- [2.2. What can't the device do:](#22-what-can't-the-device-do)
	- [2.3. What could the device do:](#23-what-could-the-device-do)
- [3. How will we develop the device:](#3-how-will-we-develop-the-device)
	- [3.1. What will we use:](#31-what-will-we-use)
		- [3.1.1. Electronic components:](#311-electronic-components)
			- [3.1.1.1. Power supply:](#3111-power-supply)
			- [3.1.1.2. Current sensors:](#3112-current-sensors)
			- [3.1.1.3. Power mosfet module:](#3113-power-mosfet-module)
			- [3.1.1.4. Leds:](#3114-leds)
			- [3.1.1.5. Stm32 blue pill:](#3115-stm32-blue-pill)
			- [3.1.1.6. LoRa-E5 dev board:](#3116-lora-e5-dev-board)
			- [3.1.1.7. USB to UART module:](#3117-usb-to-uart-module)
			- [3.1.1.8. Light sensor:](#3118-light-sensor)
		- [3.1.2. Arduino:](#312-arduino)
		- [3.1.3. STM32CubeProgrammer:](#313-stm32cubeprogrammer)
	- [3.2. What will be modify:](#32-what-will-we-modify)
	- [3.3. What will be created (functions):](#33-what-will-be-created-functions)
		- [3.3.1. GetCurrentSensorData](#331-getcurrentsensordata)
		- [3.3.2. Blink](#332-blink)
		- [3.3.3. Set](#333-set)
		- [3.3.4. On](#334-on)
- [4. What for the rollout or if we need rollback:](#4-what-for-the-rollout-or-if-we-need-rollback)
	- [4.1. Rollout:](#41-rollout)
	- [4.2. Rollback:](#42-rollback)
- [5. How will we adapt it to our client:](#5-how-will-we-adapt-it-to-our-client)
- [6. How will we provide support or maintenance:](#6-how-will-we-provide-support-or-maintenance)
	- [6.1. Support:](#61-support)
	- [6.2. Maintenance:](#62-maintenance)
- [7. Project timeline:](#7-project-timeline)
	- [7.1. Week 1](#71-week-1)
	- [7.2. Week 2](#72-week-2)
	- [7.3. Week 3](#73-week-3)
	- [7.4. Week 4](#74-week-4)
	- [7.5. Week 5](#75-week-5)

</details>

--- 
# 1. How does it have to work:

This part is meant to explain how each part have to work in an ideal world where we have no problems. It's divided into three parts to allow the searches to be quicker. This device is being created because the client wants to control the device and be able to change his state ( On or Off ) without sending someone in order to also be a good point for the ecology. The client wants to have a device that will turn off automatically as long as the leds are too hot.

## 1.1. Radio part:

The project have to use a *LoRa-E5 dev board* to be connected ( with radio communication ) to a software ([The Things Network](https://thethingsnetwork.org)) Using LoRAWAN protocol. We use the LoRaWAN protocol because we need to send and receive information even if the device is not placed in a place where there is an internet connection, so using radio transmition is to make sure that we will be able to send and receive the informations we want.

## 1.2. Coding part:

We need to use [Golang](https://go.dev/) to make a code that will sense when there is or there is not current passing throught the cables and send these informations using the *LoRa-E5 dev board* so that the client can see if it is functionning well or not. To inject the code on the board, we were requested to use [Tiny GO](https://tinygo.org/). We have to write a code that will receive the datas from all the module, treat the datas ( detect abnomalities ) and send them to the app, and that will also receive from the app requests from the user to turn on or turn off the device.

## 1.3. Electronic part:

The leds will be linked with the power supply with a power mosfet module that allow us to control the energy passing by it from the power supply to the leds. We also have two current sensor, the first one will be put between the power supply and the leds to make sure the leds receive energy, we will be using this current sensor. As for the second one, we will put it on the cable between the outlet and the power supply, in order to determine if there is a problem if it comes from the leds, the power supply or just a power outage. We were also given an USB to UART module to connect the board to the computer to send code throught it.
<br>Furthermore, all the devices appart from the power supply and the leds will be connected to the board and there datas will be read to detect problems.

# 2. The possibilities:

In this part, we will explain what the device will be able to do in the field we want and what it won't be able to do. This part is used to make sure everyone understand what we have to do and what is out of the scope of the project. We will also detailled some ideas that could be good if we have enough time for them or if others want to improve it later on.

## 2.1. What can the device do:

The device is based on the need of visibility of the problems, so by connecting to the app, the client will be able to see if the leds are still on or if there is a problem and they are off.<br> 
Moreover, if there is a problem, the client will be able to see exactly where the problem is depending on where the problem is located.<br>
The client will also be able to turn on or turn off the leds from the app so that he doesn't need to send somemone on site just for that.<br> 
By making the leds blink, we will be able to control the intensity of the light and so to reduce the luminosity and reduce the energy consuption.<br> The device will automatically turn off at one in the morning and turn on at six in the morning, because of the european laws. However, the client will be able to turn it on or off during that time because according to the laws, the lights can be turned off one hour after you finished working (even if it is late in the night) and turned on one hour before the start of the activity.<br>
The device is able to change the intensity of the leds using a photosensor. <br>
The problems the device can detect are: 
 - A misfunctionning in the sensors
 - A power outage
 - A problem with the power supply
 - An overheat of the leds

## 2.2. What can't the device do:

Acording to what we are doing, the client won't be able to resolve the problem without sending a technician to see in person. Using the sensors, the device will be able to determine roughly where is the problem but won't be able to detect precisly where it is located. The client won't be able to see exactly what is happening and will only have the informations about the state of the device.

## 2.3. What could the device do:

We could add a way for the user to change the place the device is located and adapt the time break to the country laws.

# 3. How will we develop the device:

From this part, you will know how the project is build, from the parts that are existing and that are used, to the things that didn't exist and that was created passing throught the applications and IDEs used.

## 3.1. What will we use:

Here is listed all the software, hardware, and IDE that are use or were used during the creation of the device. We will also explain what each of them are used for.

### 3.1.1. Electronic components:

Here is decribed all the electronic components we used during the conception.<br>
All the components will have a link to see the datasheet of the component.

#### 3.1.1.1. Power supply:

The power supply we are using is a [*power supply*](https://glpower.eu/en/product/gpv-18/) that will take a 200 to 240V as input and will output a 12V current. That is used to reduce the voltage that input into the components that would break if we input them with hundreds of Volts. We are using a 12V current because it is the one that is used by the leds.

#### 3.1.1.2. Current sensors

There is two types of current sensors that we were given. The first one is a [Magnetic current sensor](https://electropeak.com/learn/interfacing-zmct103c-5a-ac-current-transformer-module-with-arduino/) that will read the intensity of the input current using magnetism.<br>
The second one is a [*current sensor*](https://www.elecrow.com/acs712-current-sensor-30a-p-710.html) that will measure the intensity of the current when it will pass throught the device.<br>
This sensor could be use multiple times to measure the intensity of the current in different parts of the device.

#### 3.1.1.3. Power mosfet module

The [*power mosfet module*](https://www.robotics.org.za/XY-MOS) is a device that will be controlled by the board and will depending on the code that is input in it, will allow current to pass from the power supply to the leds. It will be used to control the intensity of the light and to turn on or turn off the leds.

#### 3.1.1.4. Leds

The leds ( Light Emitting Diode ) are used to brighten the sign and will be controled using power mosfet module.

#### 3.1.1.5. Stm32 blue pill

The STM32F103C8T6 ARM Development Board Microcomputer STM32 Core Board is a low cost STM32 microcontroller development board for learners. This board is a low-cost Minimum System Development Board for ARM Microcontroller STM32F103C8T6.

#### 3.1.1.6. LoRa-E5 dev board

The [LoRa-E5 dev board](https://www.mouser.fr/new/seeed-studio/seeed-lora-e5-development-kit/) is the microcontroller that we will use as the core of the device, in which we will inject the code and that will be connected to all the modules. It is a LoRaWAN module that will allow us to send and receive data using radio communication.

#### 3.1.1.7. USB to UART module:

The [*USB to UART module*](https://www.makerfabs.com/usb-uart-convertor-cp2102.html) is a device that will allow us to connect the board to the computer and to send code to it. It will be used to inject the code on the board.

#### 3.1.1.8. Light sensor:

The [light sensor](https://eepower.com/resistor-guide/resistor-types/photo-resistor/#) is a device that will allow us to measure the intensity of the light. It will be used to determine if the light is enough or not and to send the information to the card to adjust the intensity of the leds. The sensor is a photo resistor, so the resistor will change its resistance depending on the intensity of the light. With that, we will be able to change the luminosity depending on the intensity of the sunlight.

### 3.1.2. Arduino: 

Arduino is an open source hardware ( boards, modules ) and a software ( IDE: Integrated Development Environment ) company. They are well-known to build boards and microcontrollers. Arduino is one of the most used harware in the world ( with raspberry ) and the Arduino IDE is used a lot for injecting code in the Arduino's bords by people. We used Arduino to try to manage the serial ports.

### 3.1.3. STM32CubeProgrammer:

STM32CubeProgrammer is a tool created for programming in the STM products and that replaces STM32 ST-Link Utility. It is used to inject code in the STM32F103C8T6 ARM Development Board Microcomputer STM32 Core Board. We tried to use it but were able to flash the stm32 blue pill using Visual Studio Code instead.

### 3.1.4. Visual Studio Code:

Visual Studio Code is a code editor that is used to write code. It is used to write the code that will be injected in the STM32F103C8T6 ARM Development Board Microcomputer STM32 Core Board. It is also used to write the code that will be injected in the LoRa-E5 dev board.

## 3.2. What will we modify:

Here is listed all the software, hardware, and IDE that we will modify or that we modified during the conception.

## 3.3. What will be created (functions):

Here is listed all the functions that will be created during the conception.

### 3.3.1. GetCurrentSensorData:

the GetCurrentSensorData function return the current voltage of the current sensor and can be used by any of the two current sensors that we have, depending on the input that is used when calling the function.

### 3.3.2. Blink:

The Blink function will be used to change the intensity of the light emitted by the leds. It will be used to reduce the energy consuption and to make the leds blink. It will need to use the TurnOnModuleControl and the TurnOffModuleControl functions.

### 3.3.3. Set:

The Set function will be used to set the leds on or off. It will need to use the On function.

### 3.3.4. On:

The On function will be used to turn on the leds. It will need a boolean to turn on or off.


## 3.4. How will we secure it:

There will be no reason to secure the software part of the device because the only information that are passing by are the state of the leds and the requests sent by the customer.

# 4. What for the rollout or if we need rollback:

In this section, You will be explained how the rollout is planned and what would happen if we had to do a rollback during the conception of the product.

## 4.1 Rollout:

For the rollout, we think the best way to do it is to start by testing the product in all the signs of one city and wait for some months and if there is any problem, solve it ( using continous integration ) and if there is no more problems, it can be released for other cities. The rollout will be done in a way that the client will be able to use the product without any problem.

## 4.2 Rollback:

In case we need to do a rollback during the developement of the product, we would have to first select the time we need to set the rollback ( if one week ago is selected we will take the files in there states & week ago ), when we selected the good moment, verify that we didn't created some files we don't want to be deleted ( and save them if that is the case ) and finaly we would use Github to take the commits that were in place at the specified date.

# 5. How will we adapt it to our client:

From what our client want, the device has to scan the modules each thirty minutes to send to the customer the state of the device. Furthermore, if the customer wants to have a update of the device's state, he could click on a button on the app and the device's clock would be reset to zero and sending an update.

# 6. How will we provide support or maintenance:

This point is to explain how we would be able to help the client if he was in need of it.

## 6.1. Support:

To be sure that the client have a support that is understandable and that can help him, we will do an explanatory note that will contain all the problems we are aware of and the way to solve them if the client can solve them, if not, we will provide our contact for the client to be able to contact us at any time. 

## 6.2. Maintenance:

If there is any problem regarding the device during a defined period, we would have to first, try to solve it by sending messages or calling someone to explain him how to solve it, and in last resort, go see for ourselves and try to solve it, all the cost would be count in the maintenance cost rent.

# 7. Project timeline:

In the timeline, we will put every things that can be considered a millestone in it and add the date of it to have a complete agenda of what have been done all along the process.

## 7.1. Week 1:

<ins>27 september 2022:</ins> Setting the rules and explaining the project and the ideas of all the team.<br>

<ins>30 september 2022:</ins> Starting of the Functional Specifications and first attempts to connect the hardware parts.<br>

## 7.2. Week 2:

<ins>3 october 2022:</ins> Adding details to the Functional Specifications.<br>

<ins>4 october 2022:</ins> End of the first version of the Functional Specifications, start of the Technical Specifications, trying to inject code into the board.<br>

<ins>5 october 2022:</ins> First release of the Software Architecture Design.<br>

<ins>7 october 2022:</ins> As we finally got the good module to connect the LoRa-E5 to connect to the computer, we started working on the code to inject.<br>

## 7.3. Week 3:

<ins>10 october 2022:</ins> For the first day of that week, we were still trying to connect the St-Link and inject some code in it, but after some searches, we were able to determine that just needed soldering.

<ins>11 october 2022:</ins> After soldering the St-Link, we were able to inject code in it and were searching for a hardware able to do what we wanted.

<ins>14 october 2022:</ins> After being able to inject code in the board, we were able to start working on the code to inject.

## 7.4. Week 4:

<ins>17 october 2022:</ins> We managed to make the code work correctly.

<ins>18 october 2022:</ins> We were able to implement code to make use of the sensors.

<ins>19 october 2022:</ins> The LoRa got flashed.

## 7.5. Week 5:

<ins>24 october 2022:</ins> We started creating the presentation for the 28 of October.

<ins>25 october 2022:</ins> We finished preparing what we needed.

<ins>28 october 2022:</ins> Presentation of the product.