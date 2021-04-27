package com.example.demo;


import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

@SpringBootTest
public class CoExampleTest {
    @Autowired
    private CoExample coExample;

    @Test
    public void getName(){
        System.out.println(coExample.getName());
    }

    @Test
    public void get_age(){
        System.out.println(coExample.getAge());
    }

    @Test
    public void geAddress(){
        System.out.println(coExample.getAddress());
    }


}
