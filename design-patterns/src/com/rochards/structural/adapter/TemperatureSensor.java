package com.rochards.structural.adapter;

public class TemperatureSensor {
    /**
     * @return temperature in Fahrenheit
     */
    public static double acquireTemperature() {
        // in a real case scenario the temperature would be coming from some hardware
        return Math.random() * 110 - 10; // to simulate a range from -10 °F to 100 °F
    }
}
