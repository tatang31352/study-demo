package com.example.demo.controller;

import com.example.demo.entity.User2;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.servlet.ModelAndView;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * @author felix.xu
 * @createdata 4/29/2021 7:08 PM
 * @description
 */
@Controller
@RequestMapping("thymeleaf")
public class ThymeleafController {
    @GetMapping("/variable")
    public ModelAndView variable() {
        ModelAndView modelAndView = new ModelAndView("thymeleaf");
        String name = "felix.xu";
        Integer age=8;
        modelAndView.addObject("name", name);
        modelAndView.addObject("age", age);

        return modelAndView;
    }
    @GetMapping("/")
    public ModelAndView index() {
        ModelAndView modelAndView = new ModelAndView("index");
        return modelAndView;
    }
    @GetMapping("/list")
    public ModelAndView list() {
        List<Object> list = new ArrayList<Object>();
        list.add("北京");
        list.add("上海");
        list.add("深圳");
        ModelAndView modelAndView = new ModelAndView("list");
        modelAndView.addObject("list", list);
        return modelAndView;
    }
    @GetMapping("/list2")
    public ModelAndView list2() {
        List<User2> list = new ArrayList<>();
        list.add(new User2(1,"long"));
        list.add(new User2(2,"zhiran"));
        list.add(new User2(3,"zhiran"));
        ModelAndView modelAndView = new ModelAndView("list2");
        modelAndView.addObject("list", list);
        return modelAndView;
    }
    @GetMapping("/map")
    public ModelAndView map() {
        Map user= new HashMap();
        user.put("name", "姓名");
        user.put("sex", "male");
        ModelAndView modelAndView = new ModelAndView("map");
        modelAndView.addObject("map", user);
        return modelAndView;
    }
}