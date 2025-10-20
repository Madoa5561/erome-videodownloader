require 'net/http'
require 'uri'
require 'set'

class EromeDownloader
    BASE_URL = 'https://hu.erome.com'
    VIDEO_PATTERN = /<source src="(https:\/\/v\d+\.erome\.com[^"]+_720p\.mp4)"/
    USER_AGENT = 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36'
    def initialize(album_url)
        @album_url = album_url
    end

    def download
        html_content = fetch_html
        video_urls = extract_video_urls(html_content)
        return puts "No video URLs found" if video_urls.empty?
        puts "Found #{video_urls.size} video(s)"
        download_videos(video_urls)
    rescue => e
        puts "Error: #{e.message}"
    end

    private

    def fetch_html
        uri = URI(@album_url)
        Net::HTTP.start(uri.host, uri.port, use_ssl: true) do |http|
            request = build_html_request(uri)
            response = http.request(request)
            raise "HTTP Error: #{response.code}" unless response.code == '200'
            response.body
        end
    end

    def build_html_request(uri)
        Net::HTTP::Get.new(uri).tap do |request|
            request['Accept'] = 'text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8'
            request['Accept-Language'] = 'ja,en;q=0.9,en-GB;q=0.8,en-US;q=0.7'
            request['Referer'] = BASE_URL
            request['User-Agent'] = USER_AGENT
        end
    end

    def extract_video_urls(html)
        html.scan(VIDEO_PATTERN).flatten.uniq
    end

    def download_videos(urls)
        urls.each_with_index do |url, index|
            video_number = index + 1
            puts "Downloading video #{video_number}/#{urls.size}: #{url}"
            download_single_video(url, video_number)
            puts "Video #{video_number} downloaded successfully"
        rescue => e
            puts "Error downloading video #{video_number}: #{e.message}"
        end
    end

    def download_single_video(url, index)
        uri = URI(url)
        filename = generate_filename(uri.path, index)
        Net::HTTP.start(uri.host, uri.port, use_ssl: true) do |http|
            request = build_video_request(uri)
            response = http.request(request)
            raise "HTTP Error: #{response.code}" unless %w[200 206].include?(response.code)
            save_video(response, filename)
            log_download_info(response, filename)
        end
    end

    def build_video_request(uri)
        Net::HTTP::Get.new(uri).tap do |request|
            request['Accept'] = '*/*'
            request['Accept-Encoding'] = 'identity;q=1, *;q=0'
            request['Accept-Language'] = 'ja,en;q=0.9,en-GB;q=0.8,en-US;q=0.7'
            request['Range'] = 'bytes=0-'
            request['Referer'] = BASE_URL
            request['Sec-Fetch-Dest'] = 'video'
            request['Sec-Fetch-Mode'] = 'no-cors'
            request['Sec-Fetch-Site'] = 'same-site'
            request['User-Agent'] = USER_AGENT
        end
    end

    def generate_filename(path, index)
        original_filename = File.basename(path)
        name_without_ext = original_filename.gsub('_720p.mp4', '')
        "video_#{index}_#{name_without_ext}.mp4"
    end

    def save_video(response, filename)
        File.binwrite(filename, response.body)
    end

    def log_download_info(response, filename)
        puts "Status: #{response.code}"
        puts "Content-Length: #{response['Content-Length'] || 'Unknown'}"
        puts "Saved as: #{filename}"
    end
end

if __FILE__ == $0
    # Example: album_url = "https://www.erome.com/a/Ijac718O"
    album_url = "https://www.erome.com/a/*****"
    downloader = EromeDownloader.new(album_url)
    downloader.download
end
