package com.rochards.structural.bridge;

// The abstraction from Bridge Pattern
public abstract class Payment {

    protected final PaymentMode paymentMode;

    protected Payment(PaymentMode paymentMode) {
        this.paymentMode = paymentMode;
    }

    protected abstract void pay();
}
