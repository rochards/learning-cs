package com.rochards.creational.structural.adapter;

public class DigitalTemperatureDisplayAdapter implements TemperatureDisplay {
    private final DigitalTemperatureDisplay display;

    public DigitalTemperatureDisplayAdapter() {
        this.display = new DigitalTemperatureDisplay();
    }

    @Override
    public void print(double temperature) {
        var fahrenheitToCelsius = (temperature - 32) * 5 / 9;
        display.printMessage(new DigitalTemperatureDisplay.CelsiusTemperature(fahrenheitToCelsius));
    }
}
