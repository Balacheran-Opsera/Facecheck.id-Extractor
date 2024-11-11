# Facecheck.id-Extractor
FacecheckID Extractor is a Python script that automates the extraction of source URLs from FaceCheck.id search results. By sending a POST request with a specified id_search, it decodes base64-encoded image data in the response to reveal embedded URLs. This script outputs the extracted URLs to a text file, making it easy to access the sources of search results for further analysis.


# How it works
Navigate to https://facecheck.id and upload an image to search. Notice in the browser's address bar, an ID number is appended to the URL. The ID is called id_search in the underlying POST /api/search request. Copy this ID and paste it in Line 7 of the python script.

As long as the id_search value is valid, the script should create extracted_urls.txt which contains all of the source URLs from the search results. This way you can use facecheck.id and browse your results without buying credits.

**Install the requests library:** pip install requests

**Usage:** python3 facecheck-extractor.py

