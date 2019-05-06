from flask import Flask

app = Flask(__name__)


@app.route("/")
def hello_whale():
    return "welcome from driver service, using python and flask microframework!"

if __name__ == "__main__":
    app.run(debug=True, host="0.0.0.0", port=4000)
