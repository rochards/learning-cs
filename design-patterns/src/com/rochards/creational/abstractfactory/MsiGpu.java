package com.rochards.creational.abstractfactory;

// Concrete product
public class MsiGpu implements Gpu {
    @Override
    public void assemble() {
        System.out.println("Assembling Msi GPU");
    }
}
