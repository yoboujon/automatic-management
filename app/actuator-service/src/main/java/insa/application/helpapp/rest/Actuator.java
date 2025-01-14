package insa.application.helpapp.rest;

import jakarta.persistence.*;
import java.time.LocalDateTime;

@Entity
public class Actuator {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    private String name;
    private String type;
    private int value;
    private int room;
    private LocalDateTime updatedAt;

    public Actuator() {
    }

    public Actuator(String name, String type, int value, int room) {
        this.name = name;
        this.type = type;
        this.value = value;
        this.room = room;
        this.updatedAt = LocalDateTime.now();
    }

    // Getters and Setters

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getType() {
        return type;
    }

    public void setType(String type) {
        this.type = type;
    }

    public int getValue() {
        return value;
    }

    public void setValue(int value) {
        this.value = value;
        this.updatedAt = LocalDateTime.now();
    }

    public int getRoom() {
        return room;
    }

    public void setRoom(int room) {
        this.room = room;
    }

    public LocalDateTime getUpdatedAt() {
        return updatedAt;
    }
}