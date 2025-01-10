package insa.application.helpapp.rest;

import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import java.time.LocalDateTime;

@Entity // This annotation marks the class as a JPA entity for the sensor table
public class Sensor {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    private String type;
    private Double value;
    private LocalDateTime timestamp;

    public Sensor() {}

    public Sensor(String type, Double value, LocalDateTime timestamp) {
        this.type = type;
        this.value = value;
        this.timestamp = timestamp;
    }

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public String getType() {
        return type;
    }

    public void setType(String type) {
        this.type = type;
    }

    public Double getValue() {
        return value;
    }

    public void setValue(Double value) {
        this.value = value;
    }

    public LocalDateTime getTimestamp() {
        return timestamp;
    }

    public void setTimestamp(LocalDateTime timestamp) {
        this.timestamp = timestamp;
    }

    @Override
    public String toString() {
        return "Sensor{" +
                "id=" + id +
                ", type='" + type + '\'' +
                ", value=" + value +
                ", timestamp=" + timestamp +
                '}';
    }
}
