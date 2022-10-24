<h1 style="text-align: center">Project 1 smart signage project 4 group - Functional Specifications</h1>

<details> 
<summary style="text-decoration: underline; font-size:150%">Table of contents:</summary>

- [1. Introduction](#1-introduction)
- [2. Glossary:](#2-glossary)
- [3. Stakeholders:](#3-stakeholders)
  - [3.1. The sponsor:](#31-the-sponsor)
  - [3.2. The engineering team:](#32-the-engineering-team)
- [4. Monitoring of the signage:](#4-monitoring-of-the-signage)
- [5. Remote control of the signage:](#5-remote-control-of-the-signage)
- [6. Consumption reduction and compliance with environmental laws:](#6-consumption-reduction-and-compliance-with-environmental-laws)
- [7. Personas:](#7-personas)
  - [7.1 Store manager:](#71-store-manager)
  - [7.2 Technician:](#72-technician)
  - [7.3 CEO of SignAll:](#73-ceo-of-signall)
  - [7.4 Roles on the application:](#74-roles-on-the-application)
- [8. The technology:](#8-the-technology)
- [9. Improvements for V2?](#9-improvements-for-v2)
  - [9.1 Software:](#91-software)
  - [9.2 Hardware:](#92-hardware)
- [10. Conclusion:](#10-conclusion)
</details>

---  

## 1. Introduction 

All stores are equipped with signs so that everyone can know their name. However, the sign sector has not known any innovation for several decades already. <br>
That's why SignAll wants to change that! Through the innovations they want to bring, the company wishes to develop and revolutionize this field.

They contacted us to created this intelligent sign they want to commercialize. They want their client to be able to control their signages wherever and whenever they want.

## 2. Glossary:

- LED: a light-emitting diode (a semiconductor diode which glows when a voltage is applied)
- dim: make or become less bright or distinct
- software:
- hardware: the machines, wiring, and other physical components of a computer or other electronic system
- GO: programming language

## 3. Stakeholders:

During this project for SignAll, there are 2 different kind of stakeholders.

### 3.1. The sponsor:

First of all, there is the client, at the beginning of the project it will be represented by Franck JEANNIN, but after it, it will be directly the sponsor : SignAll <br>

Specialized in signs and signage for more than 50 years for international brands, the SignAll's field of action covers all the media that bring the company image to life: signs, flags, illuminated lettering, billboards, badges, boxes, banners, totems, interior signage, adhesives...

The SignAll team who works on this project is composed of:
- Bertrand MEURIOT: CEO of SignAll 
- Cédric DUMEZ: IT manager at SignAll
- Nicolas RICHARD: Industrial manager at SignAll 

### 3.2. The engineering team:

The second stakeholders type is the engineering team who will directly work on the development and the conception of the project. This team his made of 5 people : <br>
* Robin DEBRY as the project manager 
* Quentin CLEMENT as the program manager
* Laurent BOUQUIN as the technical leader
* Florent HUREAUX as the software engineer
* Alexis LASSELIN as the quality insurance

## 4. Monitoring of the signage:

The monitoring of the signage is the main point of the development of this project. <br>

The user will use an application connected to the signage via antennas. 
He should be able to see if the signage is on, if it is off, the actual brightness of the sign (in percent depending on the maximum brightness of the led) and the time since the sign is on (in hours and minutes no need to display the seconds and minus (this would force to send information very regularly without real use). <br>

This function will really be an informative one, it has the role of informing the user of the status of the signage in real time.

## 5. Remote control of the signage:

The remote control of the signage is the second parameter to consider during this project. <br>

The user will be able to manage the signage via an application. The user must be able to choose either to switch on the signage or to switch it off by clicking on a button on the application. He will also be able to change the brightness of the LED's just be adjusting a cursor as you would do with your phone. <br>
The user will also be able to configure an automatic on/off time that he will be able to change at anytime.

This will be the interactive features between the signage and the user.

## 6. Consumption reduction and compliance with environmental laws:

As we all now, nowadays global warming is in everyone's mind and a large number of constraints are imposed on companies to avoid the various pollutions. <br>
In the case of signages it is light pollution and electricity consumption that are targeted by the different laws.

To help SignAll in this field, we will add an intelligent feature to replace the human action. If the signage wasn't turned off during the night, it will be automatically turned off at 1am and on again at 6am (french law of the 10/07/2022). <br>
If the user needs to turn it on during this period, he can do so anyway.
Depending on the 

## 7. Personas:

### 7.1 Store manager:

Phillipe CHATRIER is the manager of AXA Vierzon since 2012. He is 47 years old married with 2 children who are now adults. He loves sport and road trips with his wife during his free time. He lives in Bourges, he takes the train every morning and evening. <br>
His job is to manage the whole store so dealing with his signage is not his priority and can sometimes be a waste of time. <br>
He has to send a technician to make an annual check to be sure the signage is working. With Appsolu, he will have the role "Gérant du magasin" and he would be able to check if the signage is on or off and to know if some LEDs are broken so he can ask a technician from SignAll to come to change it. He will receive a notification if the signage is anormaly off. 
On his days off, he can turn the panel on/off from home whenever he wants or schedule a time to turn the panel on/off automatically. He will be able to dim the LED's by himself if he thinks it is necessary. <br>
With these improvements, he will save money because the frequency of interventions will be lower, as will the electricity consumption, and it will save time by focusing on things other than signage.

### 7.2 Technician:

Suzanne LENGLEN is a technician at SignAll since 2019. She is 32 years old, married and has 2 children, and lives in Méreau with them. She loves food and spending time with her friends. <br>
Her job is to install and repair the signages in all the shops of the region. She also has to make control of the signages one time a year in every shops to know if it needs a reparation. <br>
Suzanne will have the role "Technicien", she will be able to check at every moment the status of the signages (on/off, if it's anormaly off). She has to go to AXA Vierzon every december to verify if everything is ok. With Appsolu, she will be able to avoid this annual intervention and she will only intervene if she sees that the signage has a problem and it needs to be repaired. If she needs to use the remote control features to make tests she will ask via the application the authorization to Phillipe Chatrier. <br> 
It will reduce the frequency of her visits to the shop and will be able to replace the time lost with the signage to do other tasks.

### 7.3 CEO of SignAll:

Bertrand MEURIOT is the CEO of SignAll. He is 41 years old, married and has 3 children. His hobbies are running and dogs.
<br>
His job is to manage the company and especially the commercial part. He has 80 employees under his responsability. He wants to make the company grown and he is constantly trying to innovate in the signage field, for him, Appsolu is the first innovation to improve his signs. <br>
He wants a technology that will possesse these features: monitoring of the signage, a remote control and a consumption reduction of the signages. He wants something that is addable to old signages and new ones. He wants something not too big, not laggy and something that will not be too sensitive to weather constraints. Signages are a product with a long life expectancy so he wants a product that will last a long time.

### 7.4 Roles on the application:

In the future, different roles will be created on the application to help everyone (depending on what they have to do on the sign) to help manage the signages in the best way. It will be the person in charge of each signage who will choose to whom the different roles will be attributed.

## 8. The technology:

The hardware that has been entrusted to us is a *LoRa-E5 dev board* that uses the LoRaWAN protocol.
LoRaWAN protocol is a radio communication protocol, used in low energy networks. It will be use to transmit informations and orders between the user and the STM32. <br>
The STM32 will be the element that will execute the orders of the user.
There will be other elements like ST-Link, to connect the LoRa to the STM32 and different sensors to collect information on the signage and warn the user of the status of the signage and anticipate possible failures. 
We will use GO to program the STM32.

## 9. Improvements for V2?

### 9.1 Software:

The following step of this project is to create the application for the users to let them control the signages. It will be done in the following months.

### 9.2 Hardware:

There is always room for improvement. This is why everything that could not be completed during these 6 weeks compared to what was initially planned in the specifications will obviously be added and/or improved in V2.
The whole part of the fight against breakdowns will be added in the V2. This part contains:
- fight against overheating: the temperature of the LED's will be returned, the signage will automatically be shut down at 70°C and the user will be able to know the temperature of the LED's whenever he wants. <br>
- LED failure: if a led is broken, the LED's that are on the same circuit will not work anymore (because the LED's are connected in series). If it is the LED's placed in the letters of a sign and one of them is broken, the letter will be totally extinguished. It is therefore important that the whole panel goes out if a led is broken to keep the rendering as aesthetic as possible.

## 10. Conclusion:

In spite of the specifications transmitted by signall, a lot of freedom was voluntarily left to us so that we could find ideas which they had not necessarily thought of. <br>

Thanks to the communication between the stakeholders we were able to imagine and develop a rich project in the image of ALGOSUP and SignAll. <br>

We are convinced that close collaboration between the 2 teams will help SignAll and all its client to improve their way of working. We hope it will helps them saving time and limiting the energy losses to protect the planet and their wallets.
