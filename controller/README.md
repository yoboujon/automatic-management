# Controller

## Motivations
Controller is a simulator of a specific room entierly coded in [GoLang](https://go.dev/). We choose this language because it can scale much easier than Java and does not need any Object-Oriented manipulation. It is closer to a C-synthax and has way less overhead than SpringBoot. The other argument in favor of GoLang would be that we simply wanted to learn about this language.

## Simulation
To write

## Rest API

### '/sensors' 
Expected Input

`GET` `localhost:8085/sensors` 

Expected Output
```json
[
  {
    "name": "CCS811",
    "type": "CO2",
    "id": 0,
    "value": 400,
    "unit": "ppm"
  },
  ...
]
```

### '/sensors/{id}' 
Expected Input

`GET` `localhost:8085/sensors/1` 

Expected Output
```json
{
  "name": "LM35",
  "type": "Température Intérieur",
  "id": 0,
  "value": 16.791693476679523,
  "unit": "°C"
}
```

Wrong Input

`GET` `localhost:8085/sensors/x` 

> [!WARNING]  
> If the ID entered is too high or not a number you could receive an error response

```
HTTP/1.1 400 Bad Request
id too high
You must provide an id number for the actuator.
```

### '/actuators' 
Expected Input

`GET` `localhost:8085/actuators` 

Expected Output
```json
[
  {
    "name": "Heating",
    "type": "HEATING4000",
    "id": 0,
    "value": 6
  },
  ...
]
```

### '/actuators/{id}' 
Expected Input

`GET` `localhost:8085/actuators/1` 

Expected Output
```json
{
  "name": "Windows",
  "type": "AUTOW1",
  "id": 1,
  "value": 0
}
```

`PUT` `localhost:8085/actuators/2` `{"state":6}`

Expected Output
```json
{
  "name": "Doors",
  "type": "DOOR2032X",
  "id": 2,
  "value": 6
}
```

Wrong Input

`GET`/`PUT` `localhost:8085/sensors/x` 

> [!WARNING]  
> If the ID entered is too high or not a number you could receive an error response

```
HTTP/1.1 400 Bad Request
id too high
You must provide an id number for the actuator.
```