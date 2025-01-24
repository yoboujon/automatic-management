package insa.application.helpapp.rest;

public class ActuatorStateRequest {

    private Integer state;

    public ActuatorStateRequest() {
    }

    public ActuatorStateRequest(Integer state) {
        this.state = state;
    }

    public Integer getState() {
        return state;
    }

    public void setState(Integer state) {
        this.state = state;
    }
}
