# automatic-management
TP Service Architecture based on meta-micro-micro-services


# Description des services

Sensor-Service retrieves data: indoor temperature = 22°C, outdoor temperature = 20°C, CO2 = 900 ppm, presence = detected.

Decision-Service processes rules:
Outdoor temperature is lower than indoor and between 18-27°C → Open windows.
Presence detected → Keep lights on.
CO2 threshold exceeded → Open windows further.

Actuator-Service executes:
Sends command to actuators to open windows and turn on lights.