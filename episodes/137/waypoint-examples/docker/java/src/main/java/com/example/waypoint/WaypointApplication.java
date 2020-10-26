package com.example.waypoint;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;

@Controller
@SpringBootApplication
public class WaypointApplication {

	public static void main(String[] args) {
		SpringApplication.run(WaypointApplication.class, args);
	}

	@RequestMapping("/")
	String index() {
	  return "index";
	}

}
