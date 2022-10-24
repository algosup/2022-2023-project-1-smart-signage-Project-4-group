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


## TimeRules
---
```Go
Function TimeRules()
```
This function is use to shut down the LED during the night between 1 AM to 6 AM.

#### Example Usage:
```Go
TimeRules()
```

#### Parameters:

None

#### output:

None

## On
---
```Go
Function (l LED) On() bool
```
This function return a boolean that show if the LED is ON or OFF

#### Example Usage:
```Go
l.on()
```

#### Parameters:

None

#### output:

<ins>Type</ins> : [Boolean](https://go.dev/ref/spec#Boolean_types)

## Set
---
```Go
Function (l LED) Set(value bool)
```
This function set the LED to ON/OFF

#### Example Usage:
```Go
l.set(true) // Set the selected LED to ON
```

#### Parameters:

`value` : mandatory
Boolean parameter that indicate if you want the LED to be ON or OFF

<ins>Type</ins> : [Boolean](https://go.dev/ref/spec#Boolean_types)

#### output:

None

## Blink
---
```Go
Function (l LED) Blink()
```
This function make the selected LED blink

#### Example Usage:
```Go
l.blink()
```

#### Parameters:

None

#### output:

None



