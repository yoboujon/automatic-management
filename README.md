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


# Get started

To run the different services using Spring Boot, follow these steps:

1. **Navigate to the project directory**:
    ```sh
    cd /home/robin/Documents/automatic-management
    ```

2. **Build the project using Maven**:
    ```sh
    mvn clean install
    ```

3. **Run the Sensor-Service**:
    ```sh
    mvn spring-boot:run -pl sensor-service
    ```

4. **Run the Decision-Service**:
    ```sh
    mvn spring-boot:run -pl decision-service
    ```

5. **Run the Actuator-Service**:
    ```sh
    mvn spring-boot:run -pl actuator-service
    ```
