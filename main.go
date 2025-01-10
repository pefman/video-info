package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
)

// CodecInfo holds detailed codec information for both video and audio streams
type CodecInfo struct {
	CodecType      string `json:"codec_type"`
	CodecName      string `json:"codec_name"`
	Bitrate        string `json:"bitrate"`
	ColorSpace     string `json:"color_space,omitempty"`
	ColorRange     string `json:"color_range,omitempty"`
	ColorPrimaries string `json:"color_primaries,omitempty"`
	Profile        string `json:"profile,omitempty"`
	Channels       int    `json:"channels,omitempty"`
	SampleRate     string `json:"sample_rate,omitempty"`
}

// VideoMetadata holds metadata for the entire video file
type VideoMetadata struct {
	Format      string      `json:"format"`
	Duration    string      `json:"duration"`
	Bitrate     string      `json:"bitrate"`
	StreamInfos []CodecInfo `json:"streams"`
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <video-file>\n", os.Args[0])
	}
	videoFile := os.Args[1]

	// Run the ffprobe command
	cmd := exec.Command("ffprobe", "-v", "quiet", "-print_format", "json", "-show_format", "-show_streams", videoFile)
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error running ffprobe: %v", err)
	}

	// Parse the ffprobe output
	var metadata map[string]interface{}
	err = json.Unmarshal(output, &metadata)
	if err != nil {
		log.Fatalf("Failed to parse ffprobe output: %v", err)
	}

	// Extract video-level metadata
	formatData, ok := metadata["format"].(map[string]interface{})
	if !ok {
		log.Fatalf("Failed to extract format metadata")
	}

	videoMetadata := VideoMetadata{
		Format:   getStringValue(formatData, "format_name"),
		Duration: getStringValue(formatData, "duration"),
		Bitrate:  getStringValue(formatData, "bit_rate"),
	}

	// Extract detailed stream metadata
	streams, ok := metadata["streams"].([]interface{})
	if !ok {
		log.Fatalf("Failed to extract streams metadata")
	}

	for _, stream := range streams {
		streamMap, ok := stream.(map[string]interface{})
		if !ok {
			log.Println("Skipping invalid stream entry")
			continue
		}

		codecInfo := CodecInfo{
			CodecType:      getStringValue(streamMap, "codec_type"),
			CodecName:      getStringValue(streamMap, "codec_name"),
			Bitrate:        getStringValue(streamMap, "bit_rate"),
			Profile:        getStringValue(streamMap, "profile"),
			ColorSpace:     getStringValue(streamMap, "color_space"),
			ColorRange:     getStringValue(streamMap, "color_range"),
			ColorPrimaries: getStringValue(streamMap, "color_primaries"),
		}

		// Additional audio-specific fields
		if codecInfo.CodecType == "audio" {
			codecInfo.Channels = getIntValue(streamMap, "channels")
			codecInfo.SampleRate = getStringValue(streamMap, "sample_rate")
		}

		videoMetadata.StreamInfos = append(videoMetadata.StreamInfos, codecInfo)
	}

	// Convert to JSON and print
	jsonOutput, err := json.MarshalIndent(videoMetadata, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling metadata to JSON: %v", err)
	}

	fmt.Println(string(jsonOutput))
}

// Helper function to safely get a string value from a map
func getStringValue(data map[string]interface{}, key string) string {
	if value, ok := data[key].(string); ok {
		return value
	}
	return ""
}

// Helper function to safely get an integer value from a map
func getIntValue(data map[string]interface{}, key string) int {
	if value, ok := data[key].(float64); ok {
		return int(value)
	}
	return 0
}

// Helper function to safely get a float value from a map
func getFloatValue(data map[string]interface{}, key string) float64 {
	if value, ok := data[key].(float64); ok {
		return value
	}
	return 0
}
