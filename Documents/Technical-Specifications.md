<h1 style="text-align: center">Project 1 smart signage project 4 group - Technical Specifications</h1>

<h2 style="text-decoration: underline">Table of contents:</h2>

- [1. How does it have to work:](#1-how-does-it-have-to-work)
  - [1.1. Radio part:](#11-radio-part)
  - [1.2. Coding part:](#12-coding-part)
  - [1.3. Electronic part:](#13-electronic-part)
- [2. The possibilities:](#2-the-possibilities)
  - [2.1. What can it do:](#21-what-can-it-do)
  - [2.2. What can't it do:](#22-what-cant-it-do)
- [3. How will we develop it:](#3-how-will-we-develop-it)
  - [3.1. What will we use:](#31-what-will-we-use)
    - [3.1.1. Arduino:](#311-arduino)
    - [3.1.2. STM32CubeProgrammer:](#312-stm32cubeprogrammer)
  - [3.2. What will be modify:](#32-what-will-we-modify)
  - [3.3. What will be created:](#33-what-will-be-created)
- [4. What for the rollout or if we need rollback:](#4-what-for-the-rollout-or-if-we-need-rollback)
  - [4.1. Rollout:](#41-rollout)
  - [4.2. Rollback:](#42-rollback)
- [5. How will we adapt it to our client:](#5-how-will-we-adapt-it-to-our-client)
- [6. How will we provide support or maintenance:](#6-how-will-we-provide-support-or-maintenance)
  - [6.1. Support:](#61-support)
  - [6.2. Maintenance:](#62-maintenance)
- [7. Project timeline:](#7-project-timeline)
  
--- 
## 1. How does it have to work:

This part is meant to explian how each part have to work in an ideal world where we have no problems. It's divided into three parts to allow the searches to be quicker.

### 1.1. Radio part:

The project have to use a *LoRa-E5 dev board* to be connected ( with radio communication ) to a software ( ***Put a name here*** ).

### 1.2. Coding part:

We need to use *Golang* to make a code that will sense when there is or there is not current passing throught the cables and send these informations using the *LoRa-E5 dev board* so that the client can see if it is functionning well or not.

### 1.3. Electronic part:

The leds will be linked with the [*power supply*](https://glpower.eu/en/product/gpv-18/) with a [*power mosfet module*](https://www.robotics.org.za/XY-MOS) that allow us to control the energy passing by it from the power supply to the leds. We also have two current sensor, the first one will be put between the power supply and the leds to make sure the leds recieve energy, we will be using this [*current sensor*](https://www.elecrow.com/acs712-current-sensor-30a-p-710.html). As for the second one, we will put it on the cable between the outlet and the power supply, in order to determine if there is a problem if it comes from the leds, the power supply or just a power outage.
<br>Furthermore, all the devices appart from the power supply and the leds will be connected to the board and there datas will be read to detect problems. 

## 2. The possibilities:

In this part, we will explain what the device will be able to do in the field we want and what it won't be able to do.

### 2.1 What can it do:

The device is based on the need of visibility of the problems, so by connecting to the app, the client will be able to see if the leds are still on or if there is a problem and they are off. Moreover, if there is a problem, the client will be able to se exactly where the problem is depending on where the problem is located.

### 2.2 What can't it do:

Acording to what we are doing, the client won't be able to resolve the problem without sending a technician to see in person.

## 3. How will we develop it:

From this part, you will know what the project is build with from the parts that are existing and that are used to the things that didn't exist and that was crzated passing throught the applications and IDEs used.

### 3.1. What will we use:

Here is listed all the software, hardware, and IDE that are use or were used during the creation of the device.

#### 3.1.1. Arduino: 

Arduino is an open source hardware ( boards, modules ) and a software ( IDE: Integrated Development Environment ) company. They are well-known to build boards and microcontrollers. Arduino is one of the most used harware in the world ( with raspberry ) and the Arduino IDE is used a lot for injecting code in the Arduino's bords by people.

#### 3.1.2. STM32CubeProgrammer:

STM32CubeProgrammer is a tool created for programming in the STM products and that replaces STM32 ST-Link Utility.

### 3.2. What will we modify:

(***To complete***)

### 3.3. What will be created:

(***Put functions here***)

### 3.4. How will we secure it:

(***To Complete***)

## 4. What for the rollout or if we need rollback:

In this section, You will be explained how the rollout is planned and what would happen if we have to do a rollback during the conception of the product.

### 4.1 Rollout:

(***To complete***)

### 4.2 Rollback:

In case we need to do a rollback during the developement of the product, we would have to first select the time we need to set the rollback ( if one week ago is selected we will take the files in there states & week ago ), when we selcted the good moment, verify that we didn't created some files we don't want to be deleted ( and save them if that is the case ) and finaly we would use github to take the commits that were in place at the specified date.

## 5. How will we adapt it to our client:

## 6. How will we provide support or maintenance:

### 6.1. Support:

### 6.2. Maintenance:

## 7. Project timeline: