package com.rochards.creational.singleton;

public class RunMultiThreadExamples {
    public static void main(String[] args) {
        runningWithoutSynchronization();
        runningWithSynchronization();
    }
    static void runningWithoutSynchronization() {
        System.out.printf("- The '%s' is not prepared to multi thread environment, so it lets two instance to be created%n", EasySingleton.class.getSimpleName());
        var firstThread = new Thread(() -> {
            for (var i = 0; i < 2; i++) {
                EasySingleton.getInstance();
            }
        });

        var secondThread = new Thread(() -> {
            for (var i = 0; i < 2; i++) {
                EasySingleton.getInstance();
            }
        });

        firstThread.start();
        secondThread.start();
    }

    static void runningWithSynchronization() {
        System.out.printf("%n- The '%s' is prepared to multi thread environment%n", MultiThreadSingleton.class.getSimpleName());
        var firstThread = new Thread(() -> {
            for (var i = 0; i < 2; i++) {
                MultiThreadSingleton.getInstance();
            }
        });

        var secondThread = new Thread(() -> {
            for (var i = 0; i < 2; i++) {
                MultiThreadSingleton.getInstance();
            }
        });

        firstThread.start();
        secondThread.start();
    }
}
