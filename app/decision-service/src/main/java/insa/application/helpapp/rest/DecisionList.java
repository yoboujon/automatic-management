package insa.application.helpapp.rest;

public class DecisionList {
    enum ActionEnum {
        HEATING_STOP,
        HEATING_START,
        WINDOWS_OPEN,
        WINDOWS_CLOSE,
        DOORS_OPEN,
        DOORS_CLOSE
    }

    enum SensorEnum {
        CARBON_DIOXIDE,
        TEMPERATURE,
    }

    // private ActionEnum actionEnum;
    private String action;
    private String condition;
    private static String UNKNOWN_CONDITION="Unknown condition";

    public DecisionList(ActionEnum action, SensorEnum sensor) {
        this.action = convertAction(action);
        this.condition = convertCondition(action,sensor);
    }

    public String getAction() {
        return action;
    }

    public String getCondition() {
        return condition;
    }

    private String convertAction(ActionEnum action) {
        switch (action) {
            case HEATING_STOP:
                return "Heating Stopped";
            case HEATING_START:
                return "Starting Heating";
            case WINDOWS_OPEN:
                return "Opening Windows";
            case WINDOWS_CLOSE:
                return "Closing Windows";
            case DOORS_OPEN:
                return "Opening Doors";
            case DOORS_CLOSE:
                return "Closing Doors";
            default:
                return "";
        }
    }

    private String convertCondition(ActionEnum action, SensorEnum sensor) {
        switch (action) {
            case HEATING_STOP:
                if (sensor == SensorEnum.TEMPERATURE) {
                    return "Temperature too high (> 25°C)";
                } else {
                    return UNKNOWN_CONDITION;
                }
            case HEATING_START:
                if (sensor == SensorEnum.TEMPERATURE) {
                    return "Temperature too low (< 20°C)";
                } else {
                    return UNKNOWN_CONDITION;
                }
            case WINDOWS_OPEN:
                if(sensor == SensorEnum.TEMPERATURE) {
                    return "Temperature too high (> 15°C)";
                } else if (sensor == SensorEnum.CARBON_DIOXIDE) {
                    return "Carbon Dioxide Level too high (> 800 ppm)";
                }
            case WINDOWS_CLOSE:
                if(sensor == SensorEnum.TEMPERATURE) {
                    return "Heating Device started.";
                } else {
                    return UNKNOWN_CONDITION;
                }
            case DOORS_OPEN:
                return "User Action";
            case DOORS_CLOSE:
                return "User Action";
            default:
                return UNKNOWN_CONDITION;
        }
    }
}
