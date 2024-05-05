package com.rochards.structural.bridge;

public class RunExamples {
    public static void main(String[] args) {
        Payment payment = new GooglePay(new CreditCard());
        payment.pay();
    }
}
