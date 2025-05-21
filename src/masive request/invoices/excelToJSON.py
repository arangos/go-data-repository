import pandas as pd

excel = "Invoice.xlsx"
fileOutput = "agencyInvoices.json"

# Load the Excel file
df = pd.read_excel(excel, engine="openpyxl")

# Select specific columns and rename them
selected_columns = {
    "Invoice Payment Date": "invoicePaymentDate",
    "Payment Invoice": "paymentInvoice",
    "Invoice Payment Type": "invoicePaymentType",
    "Invoice Payment Comment": "invoicePaymentComment",
    "Invoice Amount Paid": "invoiceAmountPaid",
    "Customer Code": "customerCode",
    "Invoice Created Date": "invoiceCreatedDate",
}  # Replace with actual column names

df = df[list(selected_columns.keys())].rename(columns=selected_columns)

# Convert to JSON
json_data = df.to_json(orient="records", indent=4)

# Save to a file
with open(fileOutput, "w", encoding="utf-8") as f:
    f.write(json_data)

print("Excel converted to JSON with renamed keys!")
