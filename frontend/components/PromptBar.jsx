
import { ArrowUp } from 'lucide-react';
import "../styles/PromptBar.css";
import { useRef } from 'react';
import { toast } from 'react-toastify';
import { useNavigate } from 'react-router-dom';

const redirection = (time, navigate)=>{
    toast.error("Your are not authenticated. Please login/register", {
            autoClose: time,
            onClose: () => {
            navigate("/login");
        },
    });
}

export default function SearchBar({prompt, setPrompt, language, setLanguage, currentBook, setConversation}){
    const navigate = useNavigate();
    // now add a section in order to get the response from the ai about the conversation 
    const promptUpload = async () =>{
    // if (books === null){
    //   alert("Upload the book first. \n Or select one.");
    //   return;
    // }

    const userPrompt  = {
      query: prompt,
      language: language,
      bookId: currentBook["bookId"],
      bookName: currentBook["bookName"],
    };

    setPrompt("");
    try{
      const token = localStorage.getItem("tokenId");
      if (!token){
        redirection(2000, navigate);
        return;
      }
      const response =  await fetch("http://localhost:8080/api/extractAnswer", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "Authorization": `Bearer ${token}`,
        },
        body: JSON.stringify(userPrompt),
      });

      if (!response.ok){
        // check for the not authenticated and redirect to login page
        const status = response.status;
        switch (status){
            case 401:
                redirection(2000, navigate);
                return;
            
            case 500:
                alert("Internal server error. \n Please try again later.")
                return;

            case 400:
                alert("Missing query to send.");
                return;

        }
        const err = await response.json();
        // Print the error here

        // This is the error being obtained.
        if (err.error.message == "No table"){
          alert("Please upload the book first.")
          return;
        }
        alert(err.error.message)
        return
      }
      // as the response now in the form of the stream so work according to it
      // set a reader and decoder
      const reader = response.body.getReader();
      const decoder = new TextDecoder("utf-8");

      while (true) {
        // get the streaming data
        const { done, value } = await reader.read();

        // decode the given data 
        const chunk = decoder.decode(value, {stream: true});

        // if finished reading the stream data
        if (done) break;

        setConversation(prev => {
            const updated = [...prev];
            updated[updated.length - 1].content += chunk;
            return updated;
            });
        }
    } catch(err){
        console.log("Error from the portion of handling raw question." )
    } finally{
        console.log("Finished till here of uploading file.");
    }
  }

    // make the search bar to be shown
    const textareaRef = useRef(null);
    const autoGrow = () => {
        const el = textareaRef.current;
        if (!el) return;
        el.style.height = 'auto';                 // reset so it can also shrink
        el.style.height = el.scrollHeight + 'px'; // grow to fit; CSS max-height caps it
    };

    const handleSearchClick = ()=>{
        if (!currentBook){
            alert("Insert the book first.");
            return;
        }

        if (prompt.trim().length <= 2 ){
            alert("Assign proper prompt.");
            return;
        }
        setConversation((prev)=>{
            console.log(prev);
            console.log("Array?", Array.isArray(prev));
            return [
                ...prev,
                {
                    role:"user",
                    content:prompt,
                },
                {
                    role: "assistant",
                    content: "",
                }
            ]});
        promptUpload();
    }

    const handleChange = (event)=>{
        setPrompt(event.target.value)
    }

    const handleKeyDown = (event)=>{
        if (event.key === 'Enter'){

            if (!currentBook) {
            alert("Upload the book first.");
            return;
}
            if (prompt.trim().length <= 2){
                alert("Assign proper prompt.");
                return;
            }
            // add to the conversation
            setConversation((prev)=>[
                ...prev,
                {
                    role:"user",
                    content:prompt,
                },
                {
                    role: "assistant",
                    content: "",
                }
            ]);

            promptUpload();
        }

        return;
    }
    const setLang = (event)=>{
        let data = event.target.value;
        if (data === "pl"){
            data = "English";
        }
        setLanguage(data);
    }
    const preferredLang = {
        code:"pl", name:"Preferred Language", symbol:"PL"
    }

    const languages = [
        { code: "en", name: "English", symbol: "A" },
        { code: "nep", name: "Nepali", symbol: "म" },
        { code: "zh", name: "Mandarin Chinese", symbol: "汉" },
        { code: "hi", name: "Hindi", symbol: "अ" },
        { code: "es", name: "Spanish", symbol: "Ñ" },
        { code: "fr", name: "French", symbol: "F" },
        { code: "ar", name: "Arabic", symbol: "ع" },
        { code: "bn", name: "Bengali", symbol: "অ" },
        { code: "pt", name: "Portuguese", symbol: "P" },
        { code: "ru", name: "Russian", symbol: "Я" },
        { code: "ur", name: "Urdu", symbol: "ق" }
    ];
    return (
        <div className="prompt-bar">
            <div className="select-wrapper">
            <select id="choices" name="choices" onChange={setLang}>
                <option value={preferredLang.code}>Language</option>
                {languages.map((language) => (
                <option key={language.code} value={language.name}>
                    {language.symbol} ({language.name})
                </option>
                ))}
            </select>
            </div>

            <div id="searchBar-box">
            <textarea
                ref={textareaRef}
                name="search"
                className="search-label"
                placeholder="Type your query........"
                onChange={handleChange}
                onInput={autoGrow}
                value={prompt}
                onKeyUp={handleKeyDown}
                rows={1}
            />
            <ArrowUp onClick={handleSearchClick} size={50} id="arrowup" />
            </div>
        </div>
    );
        
}