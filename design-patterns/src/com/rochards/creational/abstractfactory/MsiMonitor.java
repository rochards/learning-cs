package com.rochards.creational.abstractfactory;

public class MsiMonitor implements Monitor {
    @Override
    public void assemble() {
        System.out.println("Assembling Msi Monitor");
    }
}
