import json
import requests

# Cargar el archivo JSON
with open('Sponsor.json', 'r') as file:
    data = json.load(file)

# URL de la solicitud POST
url = "https://api.mccusa.co/sponsors"

# Credenciales de usuario
usuario = 'paginamccusa'
clave = 'postgres'

# Recorrer cada objeto del JSON
for objeto in data:
    # Realizar la solicitud POST con los datos del objeto y las credenciales de usuario
    response = requests.post(url, json=objeto, auth=(usuario, clave))
    
    # Verificar el estado de la solicitud
    if response.status_code == 201:
        print(f"Solicitud exitosa para el objeto: {objeto}")
    else:
        print(f"Error en la solicitud para el objeto: {objeto}. Código de estado: {response.status_code}")
