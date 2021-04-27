package com.example.demo;

import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.test.context.SpringBootTest;

@SpringBootTest
public class MultiYmlDemoApplicationTests {
    @Value("${myenvironment.name}")
    private String name;

    @Test
    public void getMyEnvironment(){
        System.out.println(name);
    }
}
