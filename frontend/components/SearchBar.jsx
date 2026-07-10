// show the search bar after the vector database is saved successfully
import '../styles/searchBar.css'
import '../styles/selectBar.css'

import { ArrowUp } from 'lucide-react';

export default function SearchBar({handlePrompt, setPromptInput, promptInput, setLanguage}){
    // make the search bar to be shown
    const handleSearchClick = ()=>{
        // main task here is to send this data to the prompt upload function in 
        // app.jsx and also assign data to 

        // check if the prompt input is null or what
        
        if (promptInput.trim().length <= 2 || promptInput == null){
            alert("Assign proper prompt.");
        }

        // now assign data to the handle Prompt function in App.jsx
        handlePrompt();
    }

    const handleChange = (event)=>{
        // if there is change in the data set the prompt input here
        setPromptInput(event.target.value)
    }

    const handleKeyDown = (event)=>{
        if (event.key === 'Enter'){
            console.log("Enter is pressed")
            if (promptInput.trim().length <= 2 || promptInput == null){
                alert("Assign proper prompt.");
                return;
            }
            handlePrompt();
        }
    }

    const setLang = (event)=>{
        const data = event.target.value;
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
    return(
        <div id="searchBar-box">
            {/* Take the language and assign it key and values */}
        <div className='select-wrapper'>
            <select id='choices' name='choices' onChange={setLang}>
                <option value={preferredLang.code}>Preferred answer</option>
                {languages.map((language) => (
                    <option key={language.code} value={language.name}>
                        {language.symbol} ({language.name})
                    </option>
                ))}
            </select>


        </div>
            <textarea  name="search" className="search-label" placeholder="Type your query........" onChange={handleChange} value={promptInput} onKeyUp={handleKeyDown}></textarea>
            <ArrowUp onClick={handleSearchClick} size={50} id='arrowup'/>
            {/* <button type="submit" onClick={handleSearchClick}><h2>↑</h2></button> */}
        </div>
    )
}