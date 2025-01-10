package insa.application.helpapp.rest;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.client.RestTemplate;

import java.time.LocalDateTime;
import java.util.List;

@SpringBootApplication
@RestController
@RequestMapping("/sensors")
public class SensorApplication {

    @Autowired
    private SensorRepository sensorRepository;

    public static void main(String[] args) {
        SpringApplication.run(SensorApplication.class, args);
    }

    private final String externalSensorsUrl = "http://localhost:8085/sensors/";

    @GetMapping("/refresh")
    public List<Sensor> refreshSensors() {
        // Récupération des données externes
        RestTemplate restTemplate = new RestTemplate();
        SensorExternal[] externalSensors = restTemplate.getForObject(externalSensorsUrl, SensorExternal[].class);

        if (externalSensors != null) {
            // Vider la base existante
            sensorRepository.deleteAll();

            // Enregistrer les nouvelles données
            for (SensorExternal externalSensor : externalSensors) {
                Sensor sensor = new Sensor(
                        externalSensor.getName(),
                        externalSensor.getType(),
                        externalSensor.getValue(),
                        externalSensor.getUnit(),
                        LocalDateTime.now()
                );
                sensorRepository.save(sensor);
            }
        }

        // Retourner toutes les données enregistrées
        return sensorRepository.findAll();
    }

    @GetMapping
    public List<Sensor> getAllSensors() {
        return sensorRepository.findAll();
    }
}
