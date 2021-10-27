# Performance Test
Resource usage while Locust tests show that our application is mostly bottlenecked by Network IO,
however nodejs backend did have considerably higher CPU and memory usage while maintaining the same latency as go backend.

During [test 2](https://htmlpreview.github.io/?https://github.com/WP-Team-Bach/HW1/blob/main/locust/report_2.html) in which the load was distributed evenly,
nodejs backend had almost twice the CPU usage compared to go backend (28% vs 15%) next to a higher memory usage of 38MB vs 8MB.

[Test1](https://htmlpreview.github.io/?https://github.com/WP-Team-Bach/HW1/blob/main/locust/report_1.html) was done on /sha API using nginx as a weighted load balancer (go 3 : 2 nginx).
In this scenario, average CPU usage of both backends were identical (20% node, 18% go) while memory usage was unchanged.

# How to run Locust
In order to run the Locust script, Python 3 is needed.
```
pip3 install locust
python3 -m locust
```
