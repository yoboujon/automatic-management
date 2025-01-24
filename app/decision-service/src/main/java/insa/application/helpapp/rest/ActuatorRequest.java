package insa.application.helpapp.rest;

public class ActuatorRequest {
    private int state;

    public ActuatorRequest() {}

    public ActuatorRequest(int state) {
        this.state = state;
    }

    public int getState() {
        return state;
    }
}
