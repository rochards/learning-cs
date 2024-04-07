package com.rochards.creational.factorymethod;

import java.math.BigDecimal;

public class GooglePay implements Payment{
    @Override
    public void pay(BigDecimal amount) {
        System.out.printf("Paying R$ %s with GooglePay%n", amount.toString());
    }
}
