# Function

## Light
---
```Go
Function Light(isReduce bool, isOn bool)
```

#### Example Usage:
```Go
Light(true true)
```
This function is used to light the LED and attenuate the lumiosity if necessary.

#### Parameters:

`isReduce` : mandatory
Boolean parameter that indicate if you want to reduce the intensity of the LED

<ins>Type</ins> : [Boolean](https://go.dev/ref/spec#Boolean_types)

`IsOn` : mandatory
Boolean parameter that indicate if you want to light the LED or not.

<ins>Type</ins> : [Boolean](https://go.dev/ref/spec#Boolean_types)

#### Output:

None

## GetCurrentSensorData
---
```Go
Function GetCurrentSensorData() (int, int)
```
This function is used to get the state of the two electronical sensor to see if there is an anomaly.

#### Example Usage:
```Go
GetCurrentSensorData() (int, int)
```

#### Parameters:

None

#### Output:

(int, int) : mandatory
Return two integers that show the state of the two electronical sensor 

<ins>Type</ins> : [Int](https://go.dev/ref/spec#Numeric_types)

