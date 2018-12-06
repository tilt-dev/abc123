from random import randint

from flask import Flask
app = Flask(__name__)


@app.route('/')
def rand_number():
    return str(randint(1, 10001))


if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0')
