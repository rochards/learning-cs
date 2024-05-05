package com.rochards.structural.bridge;

public class GooglePay extends Payment {

    public GooglePay(PaymentMode paymentMode) {
        super(paymentMode);
    }

    @Override
    public void pay() {
        System.out.println("GooglePay selected");
        paymentMode.executeTransaction();
    }
}
