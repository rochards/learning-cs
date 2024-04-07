package com.rochards.creational.abstractfactory;

public class RunExamples {
    public static void main(String[] args) {
        configureMsiManufacturer();
        configureAsusManufacturer();
    }

    static void configureMsiManufacturer() {
        System.out.println("-- Exemplifying with Msi Manufacturer");

        var app = new Application(new MsiManufacturer());
        app.assemble();
    }

    static void configureAsusManufacturer() {
        System.out.println("-- Exemplifying with Asus Manufacturer");

        var app = new Application(new AsusManufacturer());
        app.assemble();
    }
}

class Application {
    private final Gpu gpu;
    private final Monitor monitor;

    public Application(Manufacturer manufacturer) {
        this.gpu = manufacturer.createGpu();
        this.monitor = manufacturer.createMonitor();
    }

    public void assemble() {
        gpu.assemble();
        monitor.assemble();
    }
}
