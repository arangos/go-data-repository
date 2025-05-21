import pandas as pd

excel = "Client.xlsx"
fileOutput = "agencyClients.json"

# Load the Excel file
df = pd.read_excel(excel, engine="openpyxl")

# Select specific columns and rename them
selected_columns = {
    "email": "email",
    "passport": "passport",
    "first_name": "firstName",
    "last_name": "lastName",
    "agency": "agency",
    "client_type": "clientType",
    "sponsor": "sponsor",
    "vacancy": "vacancy",
    "second_name": "secondName",
    "completed_name": "completedName",
    "round": "round",
    "customer_code": "customerCode",
    "gender": "gender",
    "nationality": "nationality"
}  # Replace with actual column names

df = df[list(selected_columns.keys())].rename(columns=selected_columns)

# Convert to JSON
json_data = df.to_json(orient="records", indent=4)

# Save to a file
with open(fileOutput, "w", encoding="utf-8") as f:
    f.write(json_data)

print("Excel converted to JSON with renamed keys!")
