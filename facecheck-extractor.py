import requests
import base64
import json
import re

# Replace 'YOUR_ID_SEARCH' with the actual ID_Search number
ID_SEARCH = ""
output_file = "extracted_urls.txt"

# URL and headers for the POST request
url = "https://facecheck.id/api/search"
headers = {
    "User-Agent": "Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/115.0",
    "Accept": "text/plain, */*; q=0.01",
    "Content-Type": "application/json; charset=utf-8",
    "X-Requested-With": "XMLHttpRequest",
    "Origin": "https://facecheck.id",
    "Referer": "https://facecheck.id/"
}

# Payload for the POST request
payload = {
    "id_search": ID_SEARCH,
    "with_progress": True,
    "status_only": True,
    "demo": False
}

# Send the POST request to the API
response = requests.post(url, headers=headers, json=payload)

# Check if the response was successful
if response.status_code == 200:
    # Parse JSON response
    data = response.json()

    # Extract items containing base64-encoded data and scores
    items = data.get("output", {}).get("items", [])
    results = []  # List to store extracted URLs with scores

    for item in items:
        # Extract the base64 data and decode it
        base64_data = item.get("base64")
        score = item.get("score")  # Get the score
        if base64_data and score is not None:
            try:
                # Remove "data:image/webp;base64, " prefix before decoding
                base64_data = base64_data.split(",")[1]
                decoded_data = base64.b64decode(base64_data).decode("utf-8", errors="ignore")

                # Search for URL pattern in decoded data
                url_match = re.search(r'"url":"(https?://[^\"]+)"', decoded_data)
                if url_match:
                    extracted_url = url_match.group(1)
                    results.append(f"[{score}] - {extracted_url}")
            except Exception as e:
                print(f"Error decoding base64 data: {e}")

    # Write extracted URLs with scores to a file
    with open(output_file, "w") as file:
        for result in results:
            file.write(f"{result}\n")

    print(f"Extracted {len(results)} URLs and saved to '{output_file}'")
else:
    print(f"Failed to retrieve data. Status code: {response.status_code}")
