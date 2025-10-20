import requests
import re
import os
from urllib.parse import urlparse

def main():
  	# Example: album_url = "https://www.erome.com/a/Ijac718O"
    album_url = "https://www.erome.com/a/*****"
    try:
        html_content = fetch_html(album_url)
        video_urls = extract_video_urls(html_content)
        if not video_urls:
            print("No video URLs found")
            return
        print(f"Found {len(video_urls)} video(s)")
        for i, url in enumerate(video_urls, 1):
            print(f"Downloading video {i}/{len(video_urls)}: {url}")
            try:
                download_video(url, i)
                print(f"Video {i} downloaded successfully")
            except Exception as e:
                print(f"Error downloading video {i}: {e}")
    except Exception as e:
        print(f"Error fetching HTML: {e}")

def fetch_html(url):
    headers = {
        "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
        "Accept-Language": "ja,en;q=0.9,en-GB;q=0.8,en-US;q=0.7",
        "Referer": "https://hu.erome.com/",
        "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36"
    }
    response = requests.get(url, headers=headers)
    response.raise_for_status()
    return response.text

def extract_video_urls(html):
    urls = []
    seen = set()
    pattern = r'<source src="(https://v\d+\.erome\.com[^"]+_720p\.mp4)"'
    matches = re.findall(pattern, html)
    for url in matches:
        if url not in seen:
            seen.add(url)
            urls.append(url)
    return urls

def download_video(url, index):
    headers = {
        "Accept": "*/*",
        "Accept-Encoding": "identity;q=1, *;q=0",
        "Accept-Language": "ja,en;q=0.9,en-GB;q=0.8,en-US;q=0.7",
        "Range": "bytes=0-",
        "Referer": "https://hu.erome.com/",
        "Sec-Fetch-Dest": "video",
        "Sec-Fetch-Mode": "no-cors",
        "Sec-Fetch-Site": "same-site",
        "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36"
    }
    response = requests.get(url, headers=headers, stream=True)
    response.raise_for_status()
    parsed_url = urlparse(url)
    filename = os.path.basename(parsed_url.path)
    name_without_ext = filename.replace("_720p.mp4", "")
    final_filename = f"video_{index}_{name_without_ext}.mp4"
    with open(final_filename, 'wb') as f:
        for chunk in response.iter_content(chunk_size=8192):
            if chunk:
                f.write(chunk)
    print(f"Status: {response.status_code}")
    print(f"Content-Length: {response.headers.get('Content-Length', 'Unknown')}")
    print(f"Saved as: {final_filename}")

if __name__ == "__main__":
    main()
