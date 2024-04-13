package com.rochards.creational.structural.adapter;

public class DigitalTemperatureDisplay {

    /**
     * @param temperature in Celsius
     */
    public void printMessage(CelsiusTemperature temperature) {
        String message;
        double temperatureValue = temperature.value();
        if (temperatureValue <= 0) {
            message = """
                    It's freezing today. Don't go outside!
                    Temperature: %.2f °C
                    """.formatted(temperatureValue);
        } else if (temperatureValue <= 19) {
            message = """
                    It's cold outside. Go get a coat!
                    Temperature: %.2f °C
                    """.formatted(temperatureValue);
        } else if (temperatureValue <= 30) {
            message = """
                    The summer has come. The temperature is pleasant!
                    Temperature: %.2f °C
                    """.formatted(temperatureValue);
        } else if (temperatureValue <= 45) {
            message = """
                    Maybe you're living in a desert. It's too hot outside!
                    Temperature: %.2f °C
                    """.formatted(temperatureValue);
        } else {
            message = """
                    Heat wave! Don't go outside!
                    Temperature: %.2f °C
                    """.formatted(temperatureValue);
        }

        System.out.println(message);
    }

    public record CelsiusTemperature(double value) {
    }
}
