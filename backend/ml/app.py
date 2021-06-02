from flask import request, Flask, jsonify
from flask_cors import CORS
from routes import routes
import config

app = Flask(__name__)
CORS(app)

# Start Routing
routes(app)

if __name__ == "__main__":
    app.run(host=config.HOST, port=config.PORT, debug=config.DEBUG)
