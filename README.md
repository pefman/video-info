# Video Metadata Extractor

This is a command-line tool written in Go that extracts detailed metadata from video files. The tool uses `ffprobe` to gather information such as video codec, audio codec, bitrate, color mapping, color space, and other essential details, ensuring that the video is compatible with platforms like YouTube.

## Features
- Extract video codec and audio codec information.
- Retrieve bitrate, duration, and profile information.
- Gather color mapping and color space details for video.
- Validate audio properties like channels and sample rate.

## Prerequisites

1. **Go Programming Language**
   - Install Go from the [official website](https://golang.org/dl/).

2. **FFmpeg (including `ffprobe`)**
   - Ensure `ffprobe` is installed on your system.
   - Verify the installation by running:
     ```bash
     ffprobe -version
     ```
   - If not installed:
     - On Ubuntu/Debian:
       ```bash
       sudo apt install ffmpeg
       ```
     - On macOS:
       ```bash
       brew install ffmpeg
       ```
     - On Windows:
       Download the static build from [FFmpeg.org](https://ffmpeg.org/download.html).

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/video-metadata-extractor.git
   cd video-metadata-extractor
   ```

2. Install Go dependencies (if any):
   ```bash
   go mod tidy
   ```

## Usage

1. Build the application:
   ```bash
   go build -o video-metadata-extractor
   ```

2. Run the tool with a video file as input:
   ```bash
   ./video-metadata-extractor path/to/video.mp4
   ```

   Replace `path/to/video.mp4` with the path to your video file.

3. Example output:
   ```json
   {
     "format": "mp4",
     "duration": "120.34",
     "bitrate": "2500000",
     "streams": [
       {
         "codec_type": "video",
         "codec_name": "h264",
         "bitrate": "2000000",
         "color_space": "bt709",
         "color_range": "tv",
         "color_primaries": "bt709",
         "profile": "Main"
       },
       {
         "codec_type": "audio",
         "codec_name": "aac",
         "bitrate": "128000",
         "channels": 2,
         "sample_rate": "48000"
       }
     ]
   }
   ```

## Code Overview

The tool uses the `os/exec` package to call `ffprobe` and parse its JSON output to extract metadata. It is designed to handle missing or null values gracefully.

### Main Features
- **Format Information**: Extracts format name, duration, and overall bitrate.
- **Stream Information**: Processes video and audio streams to extract codec details and other relevant properties.

### Key Functions
- `getStringValue`: Safely extracts string values from a map.
- `getIntValue`: Safely extracts integer values from a map.
- `getFloatValue`: Safely extracts float values from a map.

## Contributing

Contributions are welcome! If you have suggestions for improvements or additional features, please create an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [FFmpeg](https://ffmpeg.org/) for providing a robust tool to handle video and audio files.
- The Go programming community for their excellent resources and libraries.
