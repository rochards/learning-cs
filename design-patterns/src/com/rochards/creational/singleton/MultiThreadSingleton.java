package com.rochards.creational.singleton;

public class MultiThreadSingleton {
    private static volatile MultiThreadSingleton singleton;

    private MultiThreadSingleton() {
        // Just to emulate a slow initialization
        try {
            Thread.sleep(1000);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
    }

    public static MultiThreadSingleton getInstance() {
        if (singleton != null) return singleton;

        synchronized (MultiThreadSingleton.class) {
            if (singleton == null) {
                System.out.println("New: instantiating object");
                singleton = new MultiThreadSingleton();
            }
        }

        System.out.println("Object already created");
        return singleton;
    }
}
