<h1 style="text-align: center">Project 1 smart signage project 4 group - Functional Specifications</h1>

<h2 style="text-decoration: underline">Table of contents:</h2>

- [1. Introduction](#1-introduction)
- [2. Stakeholders:](#2-stakeholders)
  - [2.1. The sponsor:](#21-the-sponsor)
  - [2.2. The engineering team:](#22-the-engineering-team)
- [3. Personaes](#3-personaes)
  - [3.1 Manager:](#31-manager)
  - [3.2 Technician:](#32-technician)
- [4. Monitoring of the signage:](#4-monitoring-of-the-signage)
- [5. Remote control of the signage:](#5-remote-control-of-the-signage)
- [6. Consumption reduction and compliance with environmental laws:](#6-consumption-reduction-and-compliance-with-environmental-laws)
- [7. The hardware:](#7-the-hardware)
  
---  

## 1. Introduction 

This project has been ordered by SignAll, a signage company in Vierzon. <br>
They want to make their signage evolve to be able to check at the status of the panels, and control it remotely to make it easier, cheaper and more respectfull of the environment.

## 2. Stakeholders:

During this project for SignAll, there are 2 different kind of stakeholders.

### 2.1. The sponsor:

First of all, there is the client, at the beginning of the project it will be represented by Franck JEANNIN, but after it, it will be directly the sponsor : SignAll <br>

Specialized in signs and signage for more than 50 years for international brands, the SignAll's field of action covers all the media that bring the company image to life: signs, flags, illuminated lettering, billboards, badges, boxes, banners, totems, interior signage, adhesives...

### 2.2. The engineering team:

The second stakeholders type is the engineering team who will directly work on the development and the conception of the project. This team his made of 5 people : <br>
* Robin DEBRY as the project manager
* Quentin CLEMENT as the program manager
* Laurent BOUQUIN as the technical leader
* Florent HUREAUX as the software engineer
* Alexis LASSELIN as the quality insurance

## 3. Personaes

### 3.1 Manager:

Phillipe CHATRIER is the manager of the technical branch of AXA. He has to send a technician to make a weekly check on every store of the brand to be sure signages are working. <br>
With the application he could be able to check every signages and turn them on/off by himself and he will receive a notification if the signage is anormaly off. 
He will need less technicians and he will not pay for the trip each week for every shops so it will save a lot of money for the company.

### 3.2 Technician:

Suzanne LENGLEN is a technician at SignAll and she is the one who works with AXA. Every monday she has to go to the Vierzon's AXA shop to check if the signage is well working. <br>
Suzanne should be able to check at every moment the state of the signage (on/off, if it's anormaly off), she will reduce the frequency of the visits to the shop and will be able to replace the time lost with the signage to do other tasks.

## 4. Monitoring of the signage:

The monitoring of the signage is the main point of the development of this project. <br>

The user will use an application connected to the signage via antennas. He should be able to see if the signage is on, if it is off, since when it is on/off, the dim of the panel, the heat of the signage to prevent from over-heating, and if it is off abnormally warn the user there could be a breakdown.

This function will really be an informative one, it has the role of informing the user of the status of the signage in real time.

## 5. Remote control of the signage:

The remote control of the signage is the second parameter to consider during this project. <br>

The user will be able to manage the signage via an application. He should be able to switch the signage on or off and change its dim manually as he would do for a smartphone.

This will be the interactive feature between the signage and the user.

## 6. Consumption reduction and compliance with environmental laws:

As we all now, nowadays global warming is in everyone's mind and a large number of constraints are imposed on companies to avoid the various pollutions. <br>
In the case of signages it is light pollution and electricity consumption that are targeted by the different laws.

To help SignAll in this field, we will add an intelligent feature to replace the human action. The signage will automatically turn on and off depending on the hour choose by the user (probably in function of environmental laws). <br>
There will also be a feature to automatically adapt the dim of the signage depending on the ambient brightness

## 7. The hardware:

The hardware that has been entrusted to us is a *LoRa-E5 dev board* that uses the LoRaWAN protocol.
LoRaWAN protocol is a radio communication protocol, used in low energy networks.
We were also given a *XY-MOS Power Mosfet Module* used to command the current given to the leds and a *current sensor ACS712* to percieve the current passing by the leds.