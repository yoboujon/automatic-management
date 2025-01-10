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

    @GetMapping
    public List<Sensor> getAllSensors() {
        // Récupération des données externes
        RestTemplate restTemplate = new RestTemplate();
        SensorExternal[] externalSensors = restTemplate.getForObject(externalSensorsUrl, SensorExternal[].class);
    
        if (externalSensors != null) {
            for (SensorExternal externalSensor : externalSensors) {
                // Rechercher un capteur existant par nom et type
                Sensor existingSensor = sensorRepository.findByNameAndType(
                        externalSensor.getName(),
                        externalSensor.getType()
                );
    
                if (existingSensor != null) {
                    // Mettre à jour les valeurs du capteur existant
                    existingSensor.setValue(externalSensor.getValue());
                    existingSensor.setUnit(externalSensor.getUnit());
                    existingSensor.setTimestamp(LocalDateTime.now());
                    sensorRepository.save(existingSensor);
                } else {
                    // Ajouter un nouveau capteur si non existant
                    Sensor newSensor = new Sensor(
                            externalSensor.getName(),
                            externalSensor.getType(),
                            externalSensor.getValue(),
                            externalSensor.getUnit(),
                            LocalDateTime.now()
                    );
                    sensorRepository.save(newSensor);
                }
            }
        }
    
        // Retourner toutes les données enregistrées
        return sensorRepository.findAll();
    }
    
}
