package com.rochards.creational.abstractfactory;

// Concrete factory
public class AsusManufacturer extends Manufacturer {
    @Override
    protected Gpu createGpu() {
        return new AsusGpu();
    }

    @Override
    protected Monitor createMonitor() {
        return new AsusMonitor();
    }
}
