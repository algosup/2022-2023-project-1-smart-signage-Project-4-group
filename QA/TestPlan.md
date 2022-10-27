<h1 style="text-align: center">Project 1 smart signage project 4 group - Test Plan</h1>

<h2 style="text-decoration: underline">Table of contents:</h2>

- [**I. Introduction**](#i-introduction)
- [**II. Analysis of the product**](#ii-analysis-of-the-product)
- [**III. Test information**](#iii-test-information)
  - [**III.1. Test strategy**](#iii1-test-strategy)
    - [**III.1.a. Scope of testing**](#iii1a-scope-of-testing)
    - [**III.1.b. Testing type**](#iii1b-testing-type)
    - [**III.1.c. Document risks and issues**](#iii1c-document-risks-and-issues)
    - [**III.1.d. Test logistics**](#iii1d-test-logistics)
  - [**III.2. Test objectives**](#iii2-test-objectives)
  - [**III.3. Test criteria**](#iii3-test-criteria)
- [**IV. Ressource planning**](#iv-ressource-planning)
- [**V. Plan test environment**](#v-plan-test-environment)
- [**VI. Schedules and estimation**](#vi-schedules-and-estimation)
- [**VII. Determine test deliverables**](#vii-determine-test-deliverables)
  
---  

## **I. Introduction**

This project has been ordered by SignAll, a signage company in Vierzon. <br>
They want to make their signage evolve to be able to check at the status of the panels, and control it remotely to make it easier, cheaper and more respectfull of the environment.

## **II. Analysis of the product**
This project has been ordered by Signall but he's gonna be used by the SignAll's customers. 
The goals of this project are: 
- Have a real time follow-up of the signage; 
- To be able to remotely control the signage;
- Apply the respect of environmental laws and reduce energy consumption. 

## **III. Test information**
### **III.1. Test strategy**
#### **III.1.a. Scope of testing**
The scope of testing determine what's gonna be tested and what isn't gonna be tested

#### **III.1.b. Testing type**
A Testing Type is a standard test procedure that gives an expected test outcome.
Each testing type is specific to one type of product bugs. But they have one common goal: <br>
Early detection of the defects before releasing the product to the customer.

![image](https://www.guru99.com/images/TestManagement/testmanagement_article_2_4_7.png)

In our case, for this project, we will need the:
- API testing;
- System testing;
- Integration test.

#### **III.1.c. Document risks and issues**
When starting a project, one must imagine all possible eventualities in order to reduce the loss of time caused by them. The risks we thought about are:
| **Risk**                                                                      | **Mitigation**                                        |
|-------------------------------------------------------------------------------|-------------------------------------------------------|
| The project schedule is too tight; it’s hard to complete this project on time | Set Priority for each of the activity.      |
| A team member is missing                                                      | Distribute his tasks to the rest of the team members. |


#### **III.1.d. Test logistics**
The tests started the 3 october, realised in a first part by Robin and Clément. <br>
In a second time, Laurent and Alexis joined Robin and Quentin to accelerate the speed of the tests. <br>

### **III.2. Test objectives**
Test Objective is the overall goal and achievement of the test execution. The objective of the testing is finding as many software defects as possible; ensure that the software under test is bug free before release.

![image](https://cdn.discordapp.com/attachments/450761035016699906/1031832888460709948/unknown.png)

### **III.3. Test criteria**
**Suspension Criteria** <br>
If the team members report that there are 10% of test cases failed, suspend testing until the
development team fixes all the failed cases.

**Exit Criteria** <br>
Specifies the criteria that denote a successful completion of a test phase
- Run rate is mandatory to be 100% unless a clear reason is given.
- Pass rate is 90%, achieving the pass rate is mandatory.

## **IV. Ressource planning**
This project is carried out by a team of 5 people: <br>
* Robin DEBRY as the project manager
* Quentin CLEMENT as the program manager
* Laurent BOUQUIN as the technical leader
* Florent HUREAUX as the software engineer
* Alexis LASSELIN as the quality insurance

There is the hardware part than we used during the test phase.<br>
* A computer with TinyGo.
* We are using is a [*power supply*](https://glpower.eu/en/product/gpv-18/) that will take a 200 to 240V as input and will output a 12V current. That is used to reduce the voltage that input into the components that would break if we input them with hundreds of Volts.

* There is two types of current sensors that we were given. The first one is a [*Magnetic current sensor*](https://electropeak.com/learn/interfacing-zmct103c-5a-ac-current-transformer-module-with-arduino/) that will read the intensity of the input current using magnetism.<br>
The second one is a [*current sensor*](https://www.elecrow.com/acs712-current-sensor-30a-p-710.html) that will measure the intensity of the current when it will pass throught the device.

* The [*power mosfet module*](https://www.robotics.org.za/XY-MOS) is a device that will be controlled by the board and will depending on the code that is input in it, will allow current to pass from the power supply to the leds.
* The leds ( Light Emitting Diode ) are used to brighten the sign and will be controled using power mosfet module.

* The [*LoRa-E5 dev board*](https://www.mouser.fr/new/seeed-studio/seeed-lora-e5-development-kit/) is the microcontroller that we will use as the core of the device, in which we will inject the code and that will be connected to all the modules. It's connected by the [*Blue Pill STM32F103C8T6*](https://rees52.com/other-devlopment-boards/2581-stm32f103c8t6-arm-stm32-minimum-system-development-board-module-for-arduino-na266), this one is put on a Bread Board.

## **V. Plan test environment**
A testing environment is a setup of software and hardware on which the testing team is going to execute test cases. The test environment consists of real business and user environment, as well as physical environments, such as server, front end running environment.

Our test environment will set up as below: <br>
![image](https://cdn.discordapp.com/attachments/450761035016699906/1035236828543995924/Diagramme_sans_nom.drawio_1.png)


## **VI. Schedules and estimation**
| **Task**                      | **Member(s)**  | **Estimate effort** |
|-------------------------------|----------------|---------------------|
| Create the test specification |Florent/ Laurent|       3 days        |
| Perform Test Execution        |Florent/ Laurent|       1 day         |
| Test Report                   |     Alexis     |        1 day        |
| Test Delivery                 |                |       1 day         |
| Total                         |                |       6 days        |

## **VII. Determine test deliverables**
Test deliverables are provided as below

**7.1 Before testing phase**
- Test plans document.
  
**7.2 During the testing**
- Test Data
- Error logs and execution logs.

**7.3 After the testing cycles is over**
- Test Results/reports
- Defect Report