from flask import Flask, request, jsonify
from flask_cors import CORS

app = Flask(__name__)
CORS(app)

uabval = ""

@app.route('/', methods=['POST','GET', 'OPTIONS'])
def receive_data():
    if request.method == 'GET':
        # Read data from the text file
        with open('uab_values.txt', 'r') as f:
            uab_values = f.readlines()
        # Remove newline characters and create a list of uab values
        uab_values = [value.strip() for value in uab_values]
        return jsonify({'uab_value': uab_values[0]})
    
    if request.method == 'OPTIONS':
        # Handle CORS preflight request
        return '', 204

    data = request.get_json()
    uab_value = data.get('uab')

    uabval = uab_value
    print('Received uab value from client:', uab_value)
    with open('uab_values.txt', 'w') as f:
        f.write(uab_value + '\n')

    # Respond to the client
    return jsonify({'message': 'Data received and saved successfully'})

if __name__ == '__main__':
    app.run(debug=True, port=5000)
