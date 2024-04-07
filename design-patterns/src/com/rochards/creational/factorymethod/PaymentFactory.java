package com.rochards.creational.factorymethod;

public class PaymentFactory {
    public static Payment create(PaymentType paymentType) {
        return switch (paymentType) {
            case PAYPAL -> new PayPal();
            case GOOGLE_PAY -> new GooglePay();
        };
    }
}
