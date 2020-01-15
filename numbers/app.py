from random import choice

import argparse

from flask import Flask
app = Flask(__name__)


@app.route('/')
def rand_number():
    nums = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
    return str(choice(nums))+'\n'


if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='start an http server that returns numbers')
    parser.add_argument('--port', type=int, default=5000, help='port on which to serve http')
    args = parser.parse_args()
    app.run(debug=True, host='0.0.0.0', port=args.port)
