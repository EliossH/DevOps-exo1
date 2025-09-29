from flask import Flask

app = Flask(__name__)

@app.route("/status")
def status():
    print("status called")
    return "Status"

@app.route("/log")
def log():
    print("log called")
    return "log"