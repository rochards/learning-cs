package com.rochards.creational.factory;

public class PaymentFactory {
    public static Payment create(PaymentType paymentType) {
        return switch (paymentType) {
            case PAYPAL -> new PayPal();
            case GOOGLE_PAY -> new GooglePay();
        };
    }
}
