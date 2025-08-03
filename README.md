# SumUp4Me

A powerful REST API service that provides intelligent audio transcription and text summarization capabilities powered by OpenAI and Google's AI models.

## 🚀 Features

- **Audio Transcription**: Convert audio files to text using OpenAI's Whisper model
- **Text Summarization**: Generate organized summaries and topic breakdowns using Google's Gemini AI
- **Combined Workflow**: Transcribe audio files and automatically summarize the content
- **Multi-format Support**: Supports MP3, WAV, and M4A audio formats
- **RESTful API**: Clean, well-documented REST endpoints
- **Configurable**: Flexible configuration through environment variables
- **Production Ready**: Comprehensive logging and error handling

## 📋 Prerequisites

- Go 1.24.1 or higher
- OpenAI API key (for audio transcription)
- Google Gemini API key (for text summarization)

## 🛠️ Installation

1. Clone the repository:
```bash
git clone https://github.com/8bury/sumup4me.git
cd sumup4me
```

2. Install dependencies:
```bash
go mod tidy
```

3. Build the application:
```bash
go build -o sumup4me ./cmd/api
```

## ⚙️ Configuration

Create a `.env` file in the root directory with the following variables:

```env
# OpenAI Configuration
OPENAI_API_KEY=your_openai_api_key_here
OPENAI_BASE_URL=https://api.openai.com/v1  # Optional, defaults to OpenAI's API

# Google Gemini Configuration
GEMINI_API_KEY=your_gemini_api_key_here
```

### Environment Variables

| Variable | Description | Required | Default |
|----------|-------------|----------|---------|
| `OPENAI_API_KEY` | Your OpenAI API key for Whisper transcription | Yes | - |
| `OPENAI_BASE_URL` | OpenAI API base URL | No | `https://api.openai.com/v1` |
| `GEMINI_API_KEY` | Your Google Gemini API key for text summarization | Yes | - |

## 🚀 Usage

1. Start the server:
```bash
./sumup4me
```

The server will start on port 8080 by default.

2. The API will be available at `http://localhost:8080`

## 📚 API Documentation

### Base URL
```
http://localhost:8080/v1
```

### Endpoints

#### 1. Transcribe Audio
Convert audio files to text.

**Endpoint:** `POST /v1/transcribe`

**Request:**
- Method: POST
- Content-Type: multipart/form-data
- Body: Audio file (key: "audio")

**Supported Formats:** MP3, WAV, M4A

**Response:**
```json
{
  "transcription": "Transcribed text content here..."
}
```

**Example using curl:**
```bash
curl -X POST http://localhost:8080/v1/transcribe \
  -F "audio=@your-audio-file.mp3"
```

#### 2. Summarize Text
Generate an organized summary of text content.

**Endpoint:** `POST /v1/sumup/text`

**Request:**
- Method: POST
- Query Parameter: `text` (the text to summarize)

**Response:**
```json
"• Topic 1: Summary point about first topic\n• Topic 2: Summary point about second topic\n..."
```

**Example using curl:**
```bash
curl -X POST "http://localhost:8080/v1/sumup/text?text=Your long text content here..."
```

#### 3. Transcribe and Summarize Audio
Transcribe audio and automatically generate a summary.

**Endpoint:** `POST /v1/sumup/audio`

**Request:**
- Method: POST
- Content-Type: multipart/form-data
- Body: Audio file (key: "audio")

**Response:**
```json
"• Topic 1: Summary of transcribed content\n• Topic 2: Another key point\n..."
```

**Example using curl:**
```bash
curl -X POST http://localhost:8080/v1/sumup/audio \
  -F "audio=@your-audio-file.mp3"
```

### Error Responses

All endpoints return appropriate HTTP status codes and error messages:

**400 Bad Request:**
```json
{
  "error": "Error description"
}
```

**500 Internal Server Error:**
```json
{
  "error": "Internal server error description"
}
```

## 🏗️ Project Structure

```
sumup4me/
├── cmd/
│   └── api/
│       └── main.go              # Application entry point
├── internal/
│   ├── audio/
│   │   └── audio.go             # Audio file validation utilities
│   ├── config/
│   │   └── config.go            # Application configuration and DI setup
│   ├── controller/
│   │   ├── sumup.controller.go      # Summary and combined workflow endpoints
│   │   └── transcribing.controller.go # Transcription endpoints
│   ├── dao/
│   │   ├── sumup.dao.go         # Gemini AI integration
│   │   └── transcribing.dao.go  # OpenAI Whisper integration
│   ├── model/
│   │   ├── transcription.model.go # Data models
│   │   └── errors.go            # Error models
│   └── service/
│       ├── sumup.service.go     # Business logic for summarization
│       └── transcribing.service.go # Business logic for transcription
├── go.mod                       # Go module definition
├── go.sum                       # Go module checksums
└── README.md                    # This file
```

## 🔧 Development

### Running in Development Mode

```bash
go run ./cmd/api
```

### Building for Production

```bash
go build -ldflags="-s -w" -o sumup4me ./cmd/api
```

### Testing the API

You can test the API using the provided curl examples or any HTTP client like Postman, Insomnia, or HTTPie.

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 Notes

- The transcription service is configured for Portuguese language by default
- Audio files are processed in memory - ensure adequate server resources for large files
- All API responses include comprehensive logging for debugging and monitoring
- The summarization uses Gemini 2.0 Flash model for optimal performance and quality

## 🔒 Security

- Store API keys securely and never commit them to version control
- Use environment variables or secure secret management systems in production
- Consider implementing rate limiting and authentication for production deployments

---

**Built with ❤️ using Go, OpenAI Whisper, and Google Gemini**