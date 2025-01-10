# automatic-management
TP Service Architecture based on meta-micro-micro-services

See [Controller](controller)

See [Application](app)

# Description des services

Sensor-Service retrieves data: indoor temperature = 22°C, outdoor temperature = 20°C, CO2 = 900 ppm, presence = detected.

Decision-Service processes rules:
Outdoor temperature is lower than indoor and between 18-27°C → Open windows.
Presence detected → Keep lights on.
CO2 threshold exceeded → Open windows further.

Actuator-Service executes:
Sends command to actuators to open windows and turn on lights.


# Get started

To run the different services using Spring Boot, follow these steps:

1. **Build the project using Maven**:
    ```sh
    mvn clean install
    ```

2. **Run the Sensor-Service**:
    ```sh
    mvn spring-boot:run -pl sensor-service
    ```

3. **Run the Decision-Service**:
    ```sh
    mvn spring-boot:run -pl decision-service
    ```

4. **Run the Actuator-Service**:
    ```sh
    mvn spring-boot:run -pl actuator-service
    ```
