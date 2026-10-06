# 📚✨ Book Buddy

### **Your book. Your language. Your conversation.**

> **What if your book could talk back?**

Book Buddy turns your PDF books into **interactive conversation partners**.

Upload a book, choose it from your library, and start asking questions. Book Buddy finds the most relevant parts of the book, understands the context, and uses a **local LLM** to generate an answer grounded in the book.

And because reading shouldn't have a language barrier, you can **ask questions in different languages and receive answers in your preferred language.**

No sending your books to a third-party AI service.

Everything runs locally using **Go + Qdrant + PostgreSQL + Ollama**.

---

## 🌟 What is Book Buddy?

Reading a 500-page book and trying to find one specific piece of information can be painful.

Book Buddy turns this:

```text
📖 Open book
   ↓
🔍 Search through hundreds of pages
   ↓
😵 Find the relevant section
   ↓
🧠 Understand the context
```

into this:

```text
📖 Upload your book
       ↓
💬 Ask a question
       ↓
🔎 Book Buddy finds the relevant passages
       ↓
🧠 Local LLM understands the context
       ↓
✨ Get an answer from your book
```

The goal isn't to replace reading.

**It's to make reading more interactive.**

---

# 🚀 Features

### 📚 Talk to your books

Upload a PDF and ask questions about its content.

Instead of manually searching through pages, simply ask:

---

### 🌍 Your language, your conversation

Book Buddy supports multilingual question answering.

You can ask questions in one language while the system retrieves information from the book and generates the response in your preferred language.

**Your book doesn't need to speak only one language.**

---

### 🧠 Retrieval-Augmented Generation

Book Buddy uses a RAG pipeline to ground answers in the uploaded book.

Instead of simply asking an LLM:

```text
Question → LLM → Answer
```

the application does:

```text
Question
   ↓
Embedding
   ↓
Qdrant similarity search
   ↓
Relevant book passages
   ↓
LLM + retrieved context
   ↓
Grounded answer
```

This helps keep the conversation focused on the selected book.

---

### 🔐 User accounts

Each user has their own books and conversations.

Authentication is handled using **JWTs**, with protected API routes for book uploads, questions, and conversation history.

---

### 💬 Conversation history

Book Buddy remembers previous conversations for each book.

You can return to a book and continue exploring it instead of starting from zero every time.

---

### 🏠 Fully local AI

The AI pipeline runs locally using **Ollama**.

That means the application can use local models for:

* Embeddings
* Translation
* Answer generation

Your uploaded books don't need to be sent to a hosted AI API.

---

# 🛠️ Tech Stack

| Layer              | Technology                          |
| ------------------ | ----------------------------------- |
| 🎨 Frontend        | React 19, Vite, React Router, Axios |
| ⚙️ Backend         | Go, Chi                             |
| 🗄️ Database       | PostgreSQL                          |
| 🔎 Vector Database | Qdrant                              |
| 🤖 AI              | Ollama                              |
| 📄 PDF Processing  | go-fitz                             |
| 🔐 Authentication  | JWT                                 |
| 🐳 Infrastructure  | Docker                              |

---

# 🏗️ How Book Buddy Works

Book Buddy has three main runtime pieces.

```text
                    ┌─────────────────────┐
                    │      React UI       │
                    │                     │
                    │ Upload • Chat       │
                    │ History • Auth      │
                    └──────────┬──────────┘
                               │
                               ▼
                    ┌─────────────────────┐
                    │      Go API         │
                    │                     │
                    │ Auth • Books • Chat │
                    │ PDF • RAG Pipeline  │
                    └──────┬───────┬──────┘
                           │       │
                ┌──────────┘       └──────────┐
                ▼                             ▼
        ┌──────────────┐              ┌──────────────┐
        │ PostgreSQL   │              │    Qdrant    │
        │              │              │              │
        │ Users        │              │ Embeddings   │
        │ Books        │              │ Book chunks  │
        │ Conversations│              │ Similarity   │
        └──────────────┘              └──────┬───────┘
                                             │
                                             ▼
                                      ┌──────────────┐
                                      │    Ollama    │
                                      │              │
                                      │ Embeddings   │
                                      │ Translation  │
                                      │ LLM          │
                                      └──────────────┘
```

---

# 🔄 The Book Journey

When you upload a book, Book Buddy takes it through a small journey.

### 1. 📤 Upload

The user uploads a PDF through the React frontend.

### 2. 📄 Extract

The Go backend extracts the text from the PDF using `go-fitz`.

### 3. ✂️ Chunk

The extracted text is split into smaller chunks that can be searched efficiently.

### 4. 🧠 Embed

Each chunk is converted into a vector embedding using the configured Ollama embedding model.

### 5. 🔎 Store

The embeddings are stored in Qdrant together with metadata such as:

```text
bookName
chunkId
```

This allows Book Buddy to search inside the correct book.

### 6. 💬 Ask

The user asks a question.

### 7. 🔍 Retrieve

The question is embedded and sent to Qdrant.

Qdrant finds the most relevant chunks from the selected book.

### 8. 🤖 Answer

The retrieved context is passed to the local LLM.

The model generates the final answer using the retrieved book content.

### 9. 🌍 Translate

If necessary, the answer can be translated back into the user's preferred language.

---

# 🧩 Repository Structure

```text
.
├── backend/
│   ├── api/                   # HTTP handlers for auth, upload, books, chat
│   ├── config/                # Environment configuration
│   ├── db/                    # PostgreSQL connection and migration support
│   ├── db/migrations/         # Database schema migrations
│   ├── internal/
│   │   ├── middlewares/       # Authentication and logging middleware
│   │   ├── models/            # Shared request/response models
│   │   └── pipeline/          # PDF, embeddings, Qdrant and LLM pipeline
│   ├── uploadedFiles/         # Uploaded PDF files
│   ├── .env                   # Local environment settings
│   ├── .env.example           # Example environment configuration
│   ├── docker-compose.yml     # PostgreSQL Docker configuration
│   ├── go.mod
│   ├── go.sum
│   ├── main.go                # Backend entrypoint
│   └── ...
│
├── frontend/
│   ├── api/                   # API client configuration
│   ├── components/            # Upload, chat, history, navbar, etc.
│   ├── pages/                 # Login, signup, home
│   ├── src/                   # Application entrypoint and styles
│   ├── styles/                # CSS
│   ├── package.json
│   └── ...
│
├── qdrant_storage/            # Local Qdrant storage
├── README.md
└── .gitignore
```

---

# ⚡ Getting Started

Want to talk to your first book?

You'll need:

* Go 1.22+
* Node.js + npm
* Docker
* PostgreSQL
* Qdrant
* Ollama

---

# 1. Clone the repository

```bash
git clone https://github.com/MilanBist/Book-Buddy.git
cd Book-Buddy
```

# 🔧 1. Configure the Environment

Create the environment file:

```bash
cp backend/.env.example backend/.env
```

Then configure your local environment.

```dotenv
PORT=8080
FRONTEND_URL=http://localhost:5173

OLLAMA_ENDPOINT=http://localhost:11434
OLLAMA_MODEL=llama3.2
OLLAMA_EMBEDDING_MODEL=bge-m3:latest
OLLAMA_TRANSLATION_MODEL=qwen3:8b

QDRANT_PORT=6334

DB_HOST=localhost
DB_PORT=5433
DB_USER=postgres
DB_PASSWORD=ragdb123
DB_NAME=ragdb

SECRET_KEY=replace_with_a_strong_secret
```

### A couple of important notes

`DB_PORT` is `5433` because Docker exposes PostgreSQL's internal port `5432` on host port `5433`.

Qdrant uses:

```text
6333 → REST
6334 → gRPC
```

The Go backend uses the Qdrant gRPC port.

---

# 🐘 2. Start PostgreSQL

From the backend directory:

```bash
cd backend
docker compose up -d postgres
```

Check that it is running:

```bash
docker compose ps
```

PostgreSQL should be available at:

```text
localhost:5433
```

---

# 🗃️ 3. Run Database Migrations

Book Buddy uses `golang-migrate`.

If you don't have it installed:

```bash
go install -tags 'postgres' \
github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

Then run:

```bash
migrate -path ./db/migrations \
  -database "postgres://postgres:ragdb123@localhost:5433/ragdb?sslmode=disable" \
  up
```

---

# 🔎 4. Start Qdrant

Run Qdrant locally with Docker:

```bash
docker run -p 6333:6333 -p 6334:6334 \
  -v "$(pwd)/qdrant_storage:/qdrant/storage" \
  qdrant/qdrant
```

Check that Qdrant is healthy:

```bash
curl http://localhost:6333/healthz
```

---

# 🤖 5. Start Ollama

Start Ollama:

```bash
ollama serve
```

Pull the required models:

```bash
ollama pull bge-m3:latest
ollama pull qwen3:8b
```

You can change the models through `.env`.

---

# ⚙️ 6. Start the Backend

Open a terminal:

```bash
cd backend
go mod download
go run .
```

The Go API should be available at:

```text
http://localhost:8080
```

---

# 🎨 7. Start the Frontend

Open another terminal:

```bash
cd frontend
npm install
npm run dev
```

The React application should be available at:

```text
http://localhost:5173
```

---

# 🔌 Frontend ↔ Backend

The frontend communicates with the Go backend through Axios.

```js
const apiClient = axios.create({
  baseURL: "http://localhost:8080/api",
  headers: { "Content-Type": "application/json" },
});
```

Example routes:

```text
POST /api/login
POST /api/register
POST /api/handlePdf
POST /api/extractAnswer
GET  /api/getBooks
GET  /api/getConversation
```

Protected routes use:

```http
Authorization: Bearer <token>
```

The frontend stores the authentication token in `localStorage` and sends it with authenticated requests.

---

# 🌐 API Reference

All API responses follow a common structure:

```json
{
  "success": true,
  "message": "Description",
  "data": {}
}
```

## 🔓 Public Endpoints

| Method | Endpoint        | Auth | Description         |
| ------ | --------------- | ---- | ------------------- |
| GET    | `/ping`         | ❌    | Simple health check |
| GET    | `/api/status`   | ❌    | API status          |
| POST   | `/api/register` | ❌    | Create a user       |
| POST   | `/api/login`    | ❌    | Authenticate a user |



# 🔐 Protected Endpoints

| Method | Endpoint               | Auth | Description                 |
| ------ | ---------------------- | ---- | --------------------------- |
| POST   | `/api/handlePdf`       | JWT  | Upload and process a PDF    |
| POST   | `/api/extractAnswer`   | JWT  | Ask a question about a book |
| GET    | `/api/getBooks`        | JWT  | Get the user's books        |
| GET    | `/api/getConversation` | JWT  | Get conversation history    |

---


# 🧠 The RAG Pipeline

The core of Book Buddy looks like this:

```text
                 USER QUESTION
                       │
                       ▼
              ┌─────────────────┐
              │ Question        │
              │ Embedding       │
              └────────┬────────┘
                       │
                       ▼
              ┌─────────────────┐
              │     Qdrant      │
              │                 │
              │ Similarity      │
              │ Search          │
              │ + Book Filter   │
              └────────┬────────┘
                       │
                       ▼
              Relevant Book Chunks
                       │
                       ▼
              ┌─────────────────┐
              │     Ollama      │
              │      LLM        │
              └────────┬────────┘
                       │
                       ▼
                  Final Answer
                       │
                       ▼
                 User's Language
```

The important part is that the LLM doesn't have to rely only on its pretrained knowledge.

**It gets relevant information from the book first.**

That's what makes the conversation book-aware.

---

# 🧪 Running Tests

Backend Go tests:

```bash
cd backend
go test ./...
```

---

# 📝 Development Notes

Uploaded PDFs are stored under:

```text
backend/uploadedFiles/
```

Qdrant's local data is stored under:

```text
qdrant_storage/
```

Local services:

```text
Frontend     → localhost:5173
Go API       → localhost:8080
PostgreSQL   → localhost:5433
Qdrant REST  → localhost:6333
Qdrant gRPC  → localhost:6334
Ollama       → localhost:11434
```

Keep environment secrets out of version control.

For production, use a proper secret-management solution instead of committing credentials to the repository.

---

# ❤️ Built With Open Source

Book Buddy stands on the shoulders of some excellent open-source projects:

* **Qdrant** — vector search and storage
* **Ollama** — local LLM and embedding inference
* **PostgreSQL** — relational data storage
* **Chi** — lightweight Go HTTP router
* **go-fitz** — PDF text extraction
* **React** — frontend UI
* **Vite** — frontend tooling

And, of course, the open-source model community that makes local AI possible.

---

# 📚 Happy Reading!

**Book Buddy**

### *Your book. Your language. Your conversation.*

> **Ask more. Understand more. Read better.** ✨