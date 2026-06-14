// show the search bar after the vector database is saved successfully


export default function SearchBar({}){
    // make the search bar to be shown
    return(
        <div id="searchBar-box">
            <textarea name="search" id="search" placeholder="Type your query........"></textarea>
            <button type="submit"><h2>↑</h2></button>
        </div>
    )
}