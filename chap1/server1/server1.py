from flask import Flask, request

app = Flask(__name__)


@app.route('/')
def hello_world():
    return f'URL.Path = {request.path!r}'


if __name__ == "__main__":
    app.run(debug=True)
