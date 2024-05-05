package com.rochards.structural.bridge;

public class CreditCard implements PaymentMode{
    @Override
    public void executeTransaction() {
        System.out.println("Paying with credit card");
    }
}
