package com.rochards.creational.structural.adapter;

public class RunExamples {
    public static void main(String[] args) {

        /*
         * Scenario: you bought a temperature sensor from the US that gives you temperature only in Fahrenheit, but you
         * have a library to print the temperature in a display that only works with Celsius
         * */

        TemperatureDisplay temperatureDisplay = new DigitalTemperatureDisplayAdapter();

        for (var i = 0; i < 5; i++) {
            temperatureDisplay.print(TemperatureSensor.acquireTemperature());
        }
    }
}
