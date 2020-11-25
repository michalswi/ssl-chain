from flask import Flask
app = Flask(__name__)

@app.route("/")
def hello():
    print("hello world!")
    return "\nworks!\n"

if __name__ == "__main__":
    app.run(host = '0.0.0.0', ssl_context = ("server.pem", "cert.key"))