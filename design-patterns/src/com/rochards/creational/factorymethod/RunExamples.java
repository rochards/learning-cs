package com.rochards.creational.factorymethod;

import java.math.BigDecimal;

public class RunExamples {
    public static void main(String[] args) {
        Payment payment1 = PaymentFactory.create(PaymentType.PAYPAL);
        payment1.pay(new BigDecimal("2000"));

        Payment payment2 = PaymentFactory.create(PaymentType.GOOGLE_PAY);
        payment2.pay(new BigDecimal("2000"));
    }
}
