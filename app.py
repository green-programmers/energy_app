from flask import Flask
from web_part import route,other

app = Flask(__name__)

@app.route(route.main)
def main_page():
    return 'Hello World!'

if __name__ == '__main__':
    app.run(debug=True ,port=other.server_port,ssl_context=other.server_ssl)