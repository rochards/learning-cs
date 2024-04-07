package com.rochards.creational.abstractfactory;

// Concrete manufacturer
public class MsiManufacturer extends Manufacturer {
    @Override
    protected Gpu createGpu() {
        return new MsiGpu();
    }

    @Override
    protected Monitor createMonitor() {
        return new MsiMonitor();
    }
}
