package com.rochards.creational.abstractfactory;

// Concrete product
public class AsusGpu implements Gpu {
    @Override
    public void assemble() {
        System.out.println("Assembling Asus GPU");
    }
}
