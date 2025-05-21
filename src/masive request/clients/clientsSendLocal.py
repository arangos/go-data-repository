import json
import requests

# Lista de archivos JSON
archivos = ['clientes-actualizada.json']

# URL de la solicitud POST
url = "http://localhost:8080/client/create"

# Credenciales de usuario
usuario = 'paginamccusa'
clave = 'postgres'

# Recorrer cada archivo
for archivo in archivos:
    # Cargar el archivo JSON
    with open(archivo, 'r', encoding='utf-8') as file:
        data = json.load(file)
    
    # Recorrer cada objeto del JSON
    for objeto in data:
        # Realizar la solicitud POST con los datos del objeto y las credenciales de usuario
        response = requests.post(url, json=objeto, auth=(usuario, clave))
        
        # Verificar el estado de la solicitud
        if response.status_code == 200:
            print(f"Solicitud exitosa para el objeto con customerCode: {objeto['customerCode']}")
        else:
            print(f"Error en la solicitud para el objeto con customerCode: {objeto['customerCode']}. Código de estado: {response.status_code}")
