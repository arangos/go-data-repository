import json
import requests

# Lista de archivos JSON
archivos = ['agencyInvoices.json']
# URL de la solicitud POST
# url = "https://api.mccusa.co/invoices"
url = "https://pruebas.api.mccusa.co/invoices"

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
        try:
            # Realizar la solicitud POST con los datos del objeto y las credenciales de usuario
            response = requests.post(url, json=objeto, auth=(usuario, clave))

            # Verificar el estado de la solicitud
            if response.status_code == 201:
                print(f"Solicitud exitosa para el objeto con customerCode: {objeto['customerCode']}")
            else:
                print(f"Error en la solicitud para el objeto con customerCode: {objeto['customerCode']}. "
      f"Código de estado: {response.status_code} - {response.text}")
        except requests.exceptions.RequestException as e:
            print(f"Error de conexión para el objeto con customerCode: {objeto['customerCode']}. Detalles: {e}")
