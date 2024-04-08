package com.rochards.creational.singleton;

public class EasySingleton {
    private static EasySingleton easySingleton;

    private EasySingleton() {
        // Just to emulate a slow initialization
        try {
            Thread.sleep(1000);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
    }

    public static EasySingleton getInstance() {
        if (easySingleton == null) {
            System.out.println("New: instantiating object");
            easySingleton = new EasySingleton();
        }

        System.out.println("Object already created");
        return easySingleton;
    }
}
