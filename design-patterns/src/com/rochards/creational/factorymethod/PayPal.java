package com.rochards.creational.factorymethod;

import java.math.BigDecimal;

public class PayPal implements Payment{
    @Override
    public void pay(BigDecimal amount) {
        System.out.printf("Paying R$ %s with Paypal%n", amount.toString());
    }
}
