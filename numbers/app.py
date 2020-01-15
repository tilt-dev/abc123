from random import choice

from flask import Flask
app = Flask(__name__)


@app.route('/')
def rand_number():
    nums = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
    return str(choice(nums))+'\n'


if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0', port=8002)
