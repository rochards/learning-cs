package com.rochards.creational.abstractfactory;

// Abstract Factory
// It can be an Interface if there's no business logic to be implemented in here
public abstract class Manufacturer {

    protected abstract Gpu createGpu();
    protected abstract Monitor createMonitor();

}
