from locust import HttpUser, task
import string
import random
import json


class Test(HttpUser):

    @task(0)
    def task1(self):
        ln = random.randint(8, 50)
        tx = ''.join(random.choices(string.ascii_letters + string.digits, k=ln))
        self.client.post("/node/sha256", {"str": tx})

    @task(0)
    def task2(self):
        ln = random.randint(8, 50)
        tx = ''.join(random.choices(string.ascii_letters + string.digits, k=ln))
        self.client.post("/go/sha256", {"str": tx})

    @task(0)
    def task3(self):
        ln = 44
        tx = ''.join(random.choices(string.ascii_letters + string.digits, k=ln)) + "="
        self.client.get("/node/sha256?sha=%s" % tx, name="/node/sha256")

    @task(0)
    def task4(self):
        ln = 44
        tx = ''.join(random.choices(string.ascii_letters + string.digits, k=ln)) + "="
        self.client.get("/go/sha256?sha=%s" % tx, name="/go/sha256")

    @task
    def task5(self):
        ln = random.randint(8, 50)
        tx = ''.join(random.choices(string.ascii_letters + string.digits, k=ln))
        res =  self.client.post("/node/sha256", {"str": tx}).text
        res = json.loads(res)
        self.client.get("/node/sha256?sha=%s" % res["result"], name="/node/sha256")

    @task
    def task6(self):
        ln = random.randint(8, 50)
        tx = ''.join(random.choices(string.ascii_letters + string.digits, k=ln))
        res =  self.client.post("/go/sha256", {"str": tx}).text
        res = json.loads(res)
        self.client.get("/go/sha256?sha=%s" % res["result"], name="/go/sha256")

