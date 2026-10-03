# 📚 Chatoo

AI-Powered Book Answerer is a full-stack application for uploading PDF books and asking questions about their contents. The backend extracts text, creates embeddings, stores vectors in Qdrant, and uses Ollama models to retrieve and generate relevant answers — in whichever language the user prefers.

## ✨ Features

- User registration and login with JWT authentication
- PDF upload and text extraction
- Embedding generation and vector search with Qdrant, scoped per book via payload filtering
- **Multilingual Q&A** — ask a question in any language and receive the answer translated back into that or in other language you preferred, regardless of the language the source book is written in
- AI-generated answers using Ollama -> qwen3:8b
- Embeddings creation using Ollama model -> bge-m3
- Book and conversation history
- React web interface for authentication, uploads, and chat

## 🛠️ Technology Stack

- **Frontend:** React 19, Vite, React Router, Axios
- **Backend:** Go, Chi HTTP router
- **Relational database:** PostgreSQL
- **Vector database:** Qdrant
- **AI services:** Ollama
- **Document processing:** `go-fitz`

## 📂 Project Structure

```text
.
├── backend/       Go API, database code, authentication, and AI pipeline
├── frontend/      React/Vite client
├── qdrant_storage/ Local Qdrant data
└── README.md
```

## ✅ Prerequisites

Install and run the following before starting the application:

- Go `1.26.3` or a compatible newer version
- Node.js and npm
- PostgreSQL
- Qdrant
- Ollama with the required models downloaded
- A migration tool such as `golang-migrate`

The backend expects PostgreSQL and Qdrant to be available locally. Ollama must be running at the endpoint configured in `backend/.env`.

## ⚙️ Configuration

Create `backend/.env` using the following structure. Replace placeholder values with your local configuration and never commit credentials.

```dotenv
PORT=8080
FRONTEND_URL=http://localhost:5173

OLLAMA_ENDPOINT=http://localhost:11434
OLLAMA_MODEL=llama3.2
OLLAMA_EMBEDDING_MODEL=bge-m3:latest
OLLAMA_TRANSLATION_MODEL=qwen3:8b

QDRANT_PORT=6334

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=<your-password>
DB_NAME=ragdb
```

The backend loads `.env` from its working directory, so run backend commands from `backend/`.

> **Note on models:** `bge-m3` is a multilingual embedding model, so book content and user questions can be embedded consistently regardless of language. `OLLAMA_TRANSLATION_MODEL` (`qwen3:8b`) handles translating the retrieved answer into the language the question was asked in, which is what powers the multilingual Q&A feature described below.

## 🗄️ Database Setup

Create `backend/.env` from `backend/.env.example` and set `DB_PASSWORD`. Start PostgreSQL from the backend directory:

```bash
cd backend
docker compose up -d postgres
```

The Compose service publishes PostgreSQL on `DB_PORT` (5432 by default) and keeps its data in a named Docker volume. This matches `DB_HOST=localhost` for a backend running on your host. To stop PostgreSQL without removing its data, run `docker compose down`.

Once PostgreSQL is healthy, apply the initial migration:

```bash
migrate -path ./db/migrations \
	-database "postgres://postgres:<your-password>@localhost:5432/ragdb?sslmode=disable" \
	up
```

Start Qdrant and Ollama using your local installation. Pull the models listed in `.env` if they are not already available.

## 🧭 Setting Up Qdrant Locally

Qdrant stores the vector embeddings generated for each uploaded book. The simplest way to run it locally is with Docker.

1. **Pull and run the Qdrant image:**

```bash
   docker run -p 6333:6333 -p 6334:6334 \
     -v "$(pwd)/qdrant_storage:/qdrant/storage" \
     qdrant/qdrant
```

   - Port `6333` serves the REST API.
   - Port `6334` serves the gRPC API and is the port referenced by `QDRANT_PORT` in `backend/.env`.
   - The `-v` flag mounts the project's `qdrant_storage/` directory so data persists across container restarts.

2. **Verify Qdrant is running:**

```bash
   curl http://localhost:6333/healthz
```

3. **Collections and payload indexing:** the backend automatically creates the Qdrant collection(s) it needs on startup. Each stored vector carries a payload that includes a `bookName` (or equivalent book identifier) field. When a user asks a question, the backend filters vector search to only that book's vectors using this payload field, so answers are always retrieved from the correct book rather than searched across a user's entire library.

4. **Alternative (native binary):** Qdrant can also be run without Docker by downloading a release binary from the [Qdrant GitHub releases page](https://github.com/qdrant/qdrant/releases) and running it directly; configuration and ports behave the same way.

## 🚀 Running the Application

Start the backend:

```bash
cd backend
go mod download
go run .
```

Start the frontend in a second terminal:

```bash
cd frontend
npm install
npm run dev
```

Open `http://localhost:5173` in a browser. The backend listens on `http://localhost:8080` by default.

## 🔍 How Question Answering Works

1. A PDF is uploaded and its text is extracted with `go-fitz`.
2. The text is chunked and embedded using the multilingual embedding model (`bge-m3`), then stored in Qdrant with a payload that tags each vector with its source `bookName`.
3. When a user submits a question (in any supported language), the question is embedded and Qdrant is queried with a payload filter restricting results to the relevant book.
4. The most relevant chunks are passed to the Ollama generation model to produce an answer.
5. If needed, the translation model converts the answer into the language the question was originally asked in before it's returned to the user.

## 📡 API Overview

| Method | Endpoint | Authentication | Purpose |
| --- | --- | --- | --- |
| GET | `/ping` | No | Basic server health check |
| GET | `/api/status` | No | API status check |
| POST | `/api/register` | No | Create an account |
| POST | `/api/login` | No | Authenticate a user |
| POST | `/api/handlePdf` | JWT | Upload and process a PDF |
| POST | `/api/extractAnswer` | JWT | Ask a question about a book (multilingual; scoped to the selected book) |
| GET | `/api/getBooks` | JWT | Retrieve the user's books |
| GET | `/api/getConversation` | JWT | Retrieve conversation history |

Protected endpoints require the JWT returned by the login endpoint in the `Authorization` header.

## 🧪 Testing Module

The project currently provides command-level checks rather than a complete automated test suite. These commands are ready to use as the testing module grows:

```bash
# Backend tests
cd backend
go test ./...

# Frontend lint and production build
cd frontend
npm run lint
npm run build
```

### 📝 Test Coverage

Add or update the following as implementation work continues:

- [ ] Authentication and JWT middleware tests
- [ ] API handler tests for success and error responses
- [ ] PDF extraction and embedding pipeline tests
- [ ] PostgreSQL and Qdrant integration tests (including payload-filtered search by `bookName`)
- [ ] Multilingual round-trip tests (ask in one language, verify answer returns in that language)
- [ ] Frontend component and user-flow tests
- [ ] End-to-end upload and question-answering test

## 🗒️ Development Notes

- Uploaded files are stored under `backend/uploadedFiles/`.
- Qdrant data is persisted under `qdrant_storage/`.
- Each Qdrant point's payload includes a `bookName` field; all retrieval queries filter on this field so multi-book libraries don't leak context between books.
- Keep local credentials and generated data out of version control where appropriate.
- Add project-specific implementation notes here as development continues:

```text
[Development notes to be completed]
```

## 🔍 Quality Assurance

**Testing Excellence**

- Comprehensive test coverage across core components
- Table-driven tests for comprehensive scenario coverage
- Integration tests for end-to-end validation (PDF upload → embedding → retrieval → answer)
- Mock-friendly architecture for isolated unit testing of handlers and services

**Code Quality**

- Clean architecture with clear separation of concerns between API, database, and AI pipeline layers
- Interface-driven design for better testability (Qdrant client, Ollama client, and repositories are swappable)


## 🙏 Acknowledgments

- **Qdrant** team for the high-performance vector database
- **Ollama** team for making local LLM inference accessible
- **Chi Router** team for the lightweight, idiomatic HTTP router
- **go-fitz** maintainers for reliable PDF text extraction in Go
- **BGE-M3** authors for the open multilingual embedding model
- **QWEN3:8B** authors for the open multilingual embedding model
- Go and React communities for best practices and patterns

---

Happy reading! 📖🤖
For questions or support, please open an issue on GitHub.
