package com.rochards.creational.singleton;

public class RunSingleThreadExamples {
    public static void main(String[] args) {
        for (var i = 0; i < 5; i++) {
            EasySingleton.getInstance();
        }
    }
}
