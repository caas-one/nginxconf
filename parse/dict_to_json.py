import json
import crossplane

def dict_to_json(path):
    dict = crossplane.parse(path)
    return json.dumps(dict)
