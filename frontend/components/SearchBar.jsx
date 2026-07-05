// show the search bar after the vector database is saved successfully
<<<<<<< HEAD

=======
import '../styles/searchBar.css'

import { ArrowUp } from 'lucide-react';
>>>>>>> fc0f258 (All commit saved)

export default function SearchBar({handlePrompt, setPromptInput, promptInput}){
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
    return(
        <div id="searchBar-box">
<<<<<<< HEAD
            <textarea  name="search" id="search" placeholder="Type your query........" onChange={handleChange} value={promptInput} onKeyDown={handleKeyDown}></textarea>
            <button type="submit" onClick={handleSearchClick}><h2>↑</h2></button>
=======
            <textarea  name="search" className="search-label" placeholder="Type your query........" onChange={handleChange} value={promptInput} onKeyUp={handleKeyDown}></textarea>
            <ArrowUp onClick={handleSearchClick} size={30} id='arrowup'/>
            {/* <button type="submit" onClick={handleSearchClick}><h2>↑</h2></button> */}
>>>>>>> fc0f258 (All commit saved)
        </div>
    )
}