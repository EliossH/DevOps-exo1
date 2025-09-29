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

if __name__ == "__main__":
    app.run(host='0.0.0.0', port=8070, debug=True)