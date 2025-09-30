from flask import Flask
from time import time
from os import statvfs
import requests

start = time()
app = Flask(__name__)

@app.route("/status", methods=["GET"])
def status():
    own = own_status()
    write_status(own)
    other = other_service_status()
    both = f"{own}\n{other}"
    response = requests.post("http://storage:8080/log", data=both)
    return both + "\n"

@app.route("/log", methods=["GET"])
def log():
    response = requests.get("http://storage:8080/log")
    return response.text

def own_status():
    now = time()
    uptime = (now - start) / 3600
    storageStat = statvfs("/")
    space = storageStat.f_bfree * storageStat.f_frsize / 1000000
    return f"Timestamp1: uptime {uptime:.2f} hours, free disk in root: {round(space)} MBytes"

def other_service_status():
    response = requests.get("http://service2:8090/status")
    return response.text

def write_status(status):
    with open("/vStorage", "a") as f:
        f.write(status + "\n")

if __name__ == "__main__":
    app.run(host='0.0.0.0', port=8070, debug=True)