import { useState } from 'react';
import NavBar from "../components/NavBar.jsx";
import SearchBar from '../components/PromptBar.jsx';
import ConversationHistory from '../components/Conversation.jsx';
import { useNavigate } from 'react-router-dom';
import ChatHistory from '../components/ChatHistory.jsx';
import '../styles/Home.css'

export default function Home() {
  const navigate = useNavigate();

  const [uploadBar, setUploadBar] = useState(false);
  const [books, setBooks] = useState([]);
  const [conversation, setConversation] = useState([]);
  const [currentBook, setCurrentBook] = useState();
  const [prompt, setPrompt] = useState();
  const [language, setLanguage] = useState();

  return (
    <div className="home">

      {/* Top Navbar */}
      <header className="navbar">
        <NavBar
          uploadBar={uploadBar}
          setUploadBar={setUploadBar}
        />
      </header>

      {/* Main area below navbar */}
      <div className="main-layout">

        {/* Left Sidebar */}
        <aside className="sidebar">
          <ChatHistory
            books={books}
            setBooks={setBooks}
            currentBook={currentBook}
            setCurrentBook={setCurrentBook}
          />
        </aside>

        {/* Right Chat Area */}
        <main className="chat-area">

          {/* Conversation */}
          <div className="conversation-container">
            <ConversationHistory
              conversation={conversation}
              setConversation={setConversation}
              currentBook={currentBook}
              uploadBar={uploadBar}
              setUploadBar={setUploadBar}
              setBooksData={setBooks}
            />
          </div>

          {/* Search / Prompt */}
          <div className="prompt-container">
            <SearchBar
              prompt={prompt}
              setPrompt={setPrompt}
              language={language}
              setLanguage={setLanguage}
              currentBook={currentBook}
              setConversation={setConversation}
            />
          </div>

        </main>
      </div>

    </div>
  );
}