from flask import Flask
import json

app = Flask(__name__)

@app.route("/")
def root():
    return json.dumps('Backend of DEA')